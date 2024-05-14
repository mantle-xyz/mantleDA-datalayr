// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@eigenlayer/contracts/libraries/Merkle.sol";
import "@eigenlayer/contracts/libraries/BN254.sol";

import "../interfaces/IDataLayrServiceManager.sol";


/**
 * @title Stand-alone utility contract that implements reuseable 'challenge'-related functionality for DataLayr.
 * @author Layr Labs, Inc.
 */
contract DataLayrChallengeUtils is Initializable {

    struct MultiRevealProof {
        BN254.G1Point interpolationPoly;
        BN254.G1Point revealProof;
        BN254.G2Point zeroPoly;
        bytes zeroPolyProof;
    }

    struct DataStoreKZGMetadata {
        BN254.G1Point c;
        uint48 degree;
        uint32 numSys;
        uint32 numPar;
    }

    // TODO: set this value
    uint256 internal ZERO_POLY_TREE_HEIGHT;

    constructor() {
        _disableInitializers();
    }

    /**
    * @notice Check that the DataLayr operator who is getting slashed was
    * actually part of the quorum for the dataStoreId.
    *
    * The burden of responsibility lies with the challenger to show that the DataLayr operator
    * is not part of the non-signers for the DataStore. Towards that end, challenger provides
    * @param nonSignerIndex such that if the relationship among nonSignerPubkeyHashes (nspkh) is:
    * uint256(nspkh[0]) <uint256(nspkh[1]) < ...< uint256(nspkh[index])< uint256(nspkh[index+1]),...
    * then,
    * uint256(nspkh[index]) <  uint256(operatorPubkeyHash) < uint256(nspkh[index+1])

    * @dev checkSignatures in DataLayrBLSSignatureChecker.sol enforces the invariant that hash of
    * non-signers pubkey is recorded in the compressed signatory record in an  ascending
    * manner.
    
     * @notice Makes sure that operatorPubkeyHash was *excluded* from set of non-signers
     * @dev Reverts if the operator *is* in the non-signer set.
     */
    function checkExclusionFromNonSignerSet(
        bytes32 operatorPubkeyHash,
        uint256 nonSignerIndex,
        IDataLayrServiceManager.SignatoryRecordMinusDataStoreId calldata signatoryRecord
    )
        external
        pure
    {
        
        if (signatoryRecord.nonSignerPubkeyHashes.length != 0) {
            // check that uint256(nspkh[index]) <  uint256(operatorPubkeyHash)
            require(
                //they're either greater than everyone in the nspkh array
                (
                    nonSignerIndex == signatoryRecord.nonSignerPubkeyHashes.length
                        && uint256(signatoryRecord.nonSignerPubkeyHashes[nonSignerIndex - 1]) < uint256(operatorPubkeyHash)
                )
                //or nonSigner index is greater than them
                || (uint256(signatoryRecord.nonSignerPubkeyHashes[nonSignerIndex]) > uint256(operatorPubkeyHash)),
                "DataLayrChallengeUtils.checkExclusionFromNonSignerSet: Provided nonsigner index is incorrect"
            );

            //  check that uint256(nspkh[index - 1]) < uint256(operatorPubkeyHash)
            if (nonSignerIndex != 0) {
                //require that the index+1 is before where operatorpubkey hash would be
                require(
                    uint256(signatoryRecord.nonSignerPubkeyHashes[nonSignerIndex - 1]) < uint256(operatorPubkeyHash),
                    "DataLayrChallengeUtils.checkExclusionFromNonSignerSet: Provided nonsigner index is incorrect"
                );
            }
        }
    }

    /**
     * @notice Makes sure that operatorPubkeyHash was *included* in set of non-signers.
     * Reverts if the operator is *not* in the non-signer set.
     */
    function checkInclusionInNonSignerSet(
        bytes32 operatorPubkeyHash,
        uint256 nonSignerIndex,
        IDataLayrServiceManager.SignatoryRecordMinusDataStoreId calldata signatoryRecord
    )
        external
        pure
    {
        require(
            operatorPubkeyHash == signatoryRecord.nonSignerPubkeyHashes[nonSignerIndex],
            "operator not included in non-signer set"
        );
    }

    /// @notice Parses the KZGMetadata from a DataStore header.
    function getDataCommitmentAndMultirevealDegreeAndSymbolBreakdownFromHeader(
        // bytes calldata header
        bytes calldata header
    )
        public
        pure
        returns (DataStoreKZGMetadata memory)
    {
        // return x, y coordinate of overall data poly commitment
        // then return degree of multireveal polynomial
        BN254.G1Point memory point;
        uint48 degree = 0;
        uint32 numSys = 0;
        uint32 numPar = 0;
        uint256 pointer;

        assembly {
            pointer := header.offset
            mstore(point, calldataload(pointer))
            mstore(add(point, 0x20), calldataload(add(pointer, 32)))
            //TODO: PUT THE LOW DEGREENESS PROOF HERE
            degree := shr(224, calldataload(add(pointer, 64)))
            numSys := shr(224, calldataload(add(pointer, 132)))
        }

        return
            DataStoreKZGMetadata({
                c: point,
                degree: degree,
                numSys: numSys,
                numPar: numPar
            });
    }

    function getNumSysFromHeader(
        // bytes calldata header
        bytes calldata header
    ) external pure returns (uint32) {
        uint32 numSys = 0;
        
        assembly {
            numSys := shr(224, calldataload(add(header.offset, 132)))
        }

        return numSys;
    }

    function getLeadingCosetIndexFromHighestRootOfUnity(uint32 i, uint32 numSys, uint32 numPar)
        public
        pure
        returns (uint32)
    {
        uint32 numNode = numSys + numPar;
        uint32 numSysE = uint32(nextPowerOf2(numSys));
        uint32 ratio = numNode / numSys + (numNode % numSys == 0 ? 0 : 1);
        uint32 numNodeE = uint32(nextPowerOf2(numSysE * ratio));

        if (i < numSys) {
            return (reverseBitsLimited(uint32(numNodeE), uint32(i)) * 256) / numNodeE;
        } else if (i < numNodeE - (numSysE - numSys)) {
            return (reverseBitsLimited(uint32(numNodeE), uint32((i - numSys) + numSysE)) * 256) / numNodeE;
        } else {
            revert("Cannot create number of frame higher than possible");
        }
    }

    function reverseBitsLimited(uint32 length, uint32 value) public pure returns (uint32) {
        uint32 unusedBitLen = 32 - uint32(log2(length));
        return reverseBits(value) >> unusedBitLen;
    }

    function reverseBits(uint32 value) public pure returns (uint32) {
        uint256 reversed = 0;
        for (uint256 i = 0; i < 32; i++) {
            uint256 mask = 1 << i;
            if (value & mask != 0) {
                reversed |= (1 << (31 - i));
            }
        }
        return uint32(reversed);
    }

    /// @notice Takes the log base 2 of n and returns it.
    function log2(uint256 n) internal pure returns (uint256) {
        require(n > 0, "Log must be defined");
        uint256 log = 0;
        while (n >> log != 1) {
            log++;
        }
        return log;
    }

    /// @notice Finds the next power of 2 greater than n and returns it.
    function nextPowerOf2(uint256 n) public pure returns (uint256) {
        uint256 res = 1;
        while (1 << res < n) {
            res++;
        }
        res = 1 << res;
        return res;
    }

    // gets the merkle root of a tree where all the leaves are the hashes of the zero/vanishing polynomials of the given multireveal
    // degree at different roots of unity. We are assuming a max of 512 datalayr nodes  right now, so, for merkle root for "degree"
    // will be of the tree where the leaves are the hashes of the G2 kzg commitments to the following polynomials:
    // l = degree (for brevity)
    // w^(512*l) = 1
    // (s^l - 1), (s^l - w^l), (s^l - w^2l), (s^l - w^3l), (s^l - w^4l), ...
    // we have precomputed these values and return them directly because it's cheap. currently we
    // tolerate up to degree 2^11, which means up to (31 bytes/point)(1024 points/dln)(256 dln) = 8 MB in a datastore
    function getZeroPolyMerkleRoot(uint256 degree) public pure returns (bytes32) {
        uint256 log = log2(degree);

        if (log == 0) {
            return 0xe82cea94884b1b895ea0742840a3b19249a723810fd1b04d8564d675b0a416f1;
        } else if (log == 1) {
            return 0x4843774a80fc8385b31024f5bd18b42e62de439206ab9468d42d826796d41f67;
        } else if (log == 2) {
            return 0x092d3e5f87f5293e7ab0cc2ca6b0b5e4adb5e0011656544915f7cea34e69e5ab;
        } else if (log == 3) {
            return 0x494b208540ec8624fbbb3f2c64ffccdaf6253f8f4e50c0d93922d88195b07755;
        } else if (log == 4) {
            return 0xfdb44b84a82893cfa0e37a97f09ffc4298ad5e62be1bea1d03320ae836213d22;
        } else if (log == 5) {
            return 0x3f50cb08231d2a76853ba9dbb20dad45a1b75c57cdaff6223bfe069752cff3d4;
        } else if (log == 6) {
            return 0xbb39eebd8138eefd5802a49d571e65b3e0d4e32277c28fbf5fbca66e7fb04310;
        } else if (log == 7) {
            return 0xf0a39b513e11fa80cbecbf352f69310eddd5cd03148768e0e9542bd600b133ec;
        } else if (log == 8) {
            return 0x038cca2238865414efb752cc004fffec9e6069b709f495249cdf36efbd5952f6;
        } else if (log == 9) {
            return 0x2a26b054ed559dd255d8ac9060ebf6b95b768d87de767f8174ad2f9a4e48dd01;
        } else if (log == 10) {
            return 0x1fe180d0bc4ff7c69fefa595b3b5f3c284535a280f6fdcf69b20770d1e20e1fc;
        } else if (log == 11) {
            return 0x60e34ad57c61cd6fdd8177437c30e4a30334e63d7683989570cf27020efc8201;
        } else if (log == 12) {
            return 0xeda2417e770ddbe88f083acf06b6794dfb76301314a32bd0697440d76f6cd9cc;
        } else if (log == 13) {
            return 0x8cbe9b8cf92ce70e3bec8e1e72a0f85569017a7e43c3db50e4a5badb8dea7ce8;
        } else {
            revert("Log not in valid range");
        }
    }

    /// @notice Opens up kzg commitment c(x) at r and makes sure c(r) = s. proof (pi) is in G2 to allow for calculation of Z in G1
    function openPolynomialAtPoint(BN254.G1Point memory c, BN254.G2Point calldata pi, uint256 r, uint256 s)
        public
        view
        returns (bool)
    {
        //we use and overwrite z as temporary storage
        //g1 = (1, 2)
        BN254.G1Point memory g1Gen = BN254.G1Point({X: 1, Y: 2});
        //calculate -g1*r = -[r]_1
        BN254.G1Point memory z = BN254.scalar_mul(BN254.negate(g1Gen), r);

        //add [x]_1 - [r]_1 = Z and store in first 2 slots of input
        //CRITIC TODO: SWITCH THESE TO [x]_1 of Powers of Tau!
        BN254.G1Point memory firstPowerOfTau = BN254.G1Point({
            X: 15397661830938158195220872607788450164522003659458108417904919983213308643927,
            Y: 4051901473739185471504766068400292374549287637553596337727654132125147894034
        });
        z = BN254.plus(firstPowerOfTau, z);
        //calculate -g1*s = -[s]_1
        BN254.G1Point memory negativeS = BN254.scalar_mul(BN254.negate(g1Gen), s);
        //calculate C-[s]_1
        BN254.G1Point memory cMinusS = BN254.plus(c, negativeS);

        //check e(z, pi)e(C-[s]_1, -g2) = 1
        return BN254.pairing(z, pi, cMinusS, BN254.negGeneratorG2());
    }

    function validateDisclosureResponse(
        DataStoreKZGMetadata memory dskzgMetadata,
        uint32 chunkNumber,
        BN254.G1Point calldata interpolationPoly,
        BN254.G1Point calldata revealProof,
        BN254.G2Point memory zeroPoly,
        bytes calldata zeroPolyProof
    )
        public
        view
        returns (bool)
    {
        require(zeroPolyProof.length/32 == ZERO_POLY_TREE_HEIGHT, "DataLayrChallengeUtils.validateDisclosureResponse: incorrect merkle proof length");

        // check that [zeroPoly.x0, zeroPoly.x1, zeroPoly.y0, zeroPoly.y1] is actually the "chunkNumber" leaf
        // of the zero polynomial Merkle tree
        {
            //deterministic assignment of "y" here
            // @todo
            require(
                Merkle.verifyInclusionKeccak(
                    // Merkle proof
                    zeroPolyProof,
                    // Merkle root hash
                    getZeroPolyMerkleRoot(dskzgMetadata.degree),
                    // leaf
                    keccak256(abi.encodePacked(zeroPoly.X[1], zeroPoly.X[0], zeroPoly.Y[1], zeroPoly.Y[0])),
                    // index in the Merkle tree
                    getLeadingCosetIndexFromHighestRootOfUnity(chunkNumber, dskzgMetadata.numSys, dskzgMetadata.numPar)
                ),
                "Incorrect zero poly merkle proof"
            );
        }

        /**
         * Doing pairing verification  e(Pi(s), Z_k(s)).e(C - I, -g2) == 1
         */
        //get the commitment to the zero polynomial of multireveal degree

        // calculate [C]_1 - [I]_1
        BN254.G1Point memory cMinusI = BN254.plus(dskzgMetadata.c, BN254.negate(interpolationPoly));

        //check e(z, pi)e(C-[s]_1, -g2) = 1
        return BN254.pairing(revealProof, zeroPoly, cMinusI, BN254.negGeneratorG2());
    }

    function nonInteractivePolynomialProof(
        bytes calldata header,
        uint32 chunkNumber,
        bytes calldata poly,
        MultiRevealProof calldata multiRevealProof,
        BN254.G2Point calldata polyEquivalenceProof
    )
        external
        view
        returns (bool)
    {
        DataStoreKZGMetadata memory dskzgMetadata =
            getDataCommitmentAndMultirevealDegreeAndSymbolBreakdownFromHeader(header);

        //verify pairing for the commitment to interpolating polynomial
        require(
            validateDisclosureResponse(
                dskzgMetadata,
                chunkNumber,
                multiRevealProof.interpolationPoly,
                multiRevealProof.revealProof,
                multiRevealProof.zeroPoly,
                multiRevealProof.zeroPolyProof
            ),
            "Reveal failed due to non 1 pairing"
        );

        // TODO: verify that this check is correct!
        // check that degree of polynomial in the header matches the length of the submitted polynomial
        // i.e. make sure submitted polynomial doesn't contain extra points
        require(
            (dskzgMetadata.degree + 1) * 32 == poly.length, "Polynomial must have a 256 bit coefficient for each term"
        );

        //Calculating r, the point at which to evaluate the interpolating polynomial
        uint256 r = uint256(
            keccak256(
                abi.encodePacked(
                    keccak256(poly),
                    multiRevealProof.interpolationPoly.X,
                    multiRevealProof.interpolationPoly.Y
                )
            )
        ) % BN254.FR_MODULUS;
        uint256 s = linearPolynomialEvaluation(poly, r);
        return
            openPolynomialAtPoint(
                multiRevealProof.interpolationPoly,
                polyEquivalenceProof,
                r,
                s
            );
    }

    function verifyPolyEquivalenceProof(
        bytes calldata poly,
        BN254.G1Point calldata interpolationPoly,
        BN254.G2Point calldata polyEquivalenceProof
    ) external view returns (bool) {
        //Calculating r, the point at which to evaluate the interpolating polynomial
        uint256 r = uint256(
            keccak256(
                abi.encodePacked(
                    keccak256(poly),
                    interpolationPoly.X,
                    interpolationPoly.Y
                )
            )
        ) % BN254.FR_MODULUS;
        uint256 s = linearPolynomialEvaluation(poly, r);
        bool ok = openPolynomialAtPoint(
            interpolationPoly,
            polyEquivalenceProof,
            r,
            s
        );
        return ok;
    }

    function verifyBatchPolyEquivalenceProof(
        bytes[] calldata polys,
        BN254.G1Point[] calldata interpolationPolys,
        BN254.G2Point calldata polyEquivalenceProof
    ) external view returns (bool) {
        bytes32[] memory rs = new bytes32[](polys.length);
        //Calculating r, the point at which to evaluate the interpolating polynomial
        for (uint i = 0; i < polys.length; i++) {
            rs[i] = keccak256(
                abi.encodePacked(
                    keccak256(polys[i]),
                    interpolationPolys[i].X,
                    interpolationPolys[i].Y
                )
            );
        }
        //this is the point to open each polynomial at
        uint256 r = uint256(keccak256(abi.encodePacked(rs))) % BN254.FR_MODULUS;
        //this is the offset we add to each polynomial to prevent collision
        //we use array to help with stack
        uint256[2] memory gammaAndGammaPower;
        gammaAndGammaPower[0] =
            uint256(keccak256(abi.encodePacked(rs, uint256(0)))) %
            BN254.FR_MODULUS;
        gammaAndGammaPower[1] = gammaAndGammaPower[0];
        //store I1
        BN254.G1Point memory gammaShiftedCommitmentSum = interpolationPolys[0];
        //store I1(r)
        uint256 gammaShiftedEvaluationSum = linearPolynomialEvaluation(
            polys[0],
            r
        );
        for (uint i = 1; i < interpolationPolys.length; i++) {
            //gammaShiftedCommitmentSum += gamma^i * Ii
            gammaShiftedCommitmentSum = BN254.plus(
                gammaShiftedCommitmentSum,
                BN254.scalar_mul(interpolationPolys[i], gammaAndGammaPower[1])
            );
            //gammaShiftedEvaluationSum += gamma^i * Ii(r)
            uint256 eval = linearPolynomialEvaluation(polys[i], r);
            gammaShiftedEvaluationSum = addmod(
                gammaShiftedEvaluationSum,
                mulmod(gammaAndGammaPower[1], eval, BN254.FR_MODULUS),
                BN254.FR_MODULUS
            );
            // gammaPower = gamma^(i+1)
            gammaAndGammaPower[1] = mulmod(
                gammaAndGammaPower[0],
                gammaAndGammaPower[1],
                BN254.FR_MODULUS
            );
        }

        return
            openPolynomialAtPoint(
                gammaShiftedCommitmentSum,
                polyEquivalenceProof,
                r,
                gammaShiftedEvaluationSum
            );
    }

    function batchNonInteractivePolynomialProofs(
        bytes calldata header,
        uint32 firstChunkNumber,
        bytes[] calldata polys,
        MultiRevealProof[] calldata multiRevealProofs,
        BN254.G2Point calldata polyEquivalenceProof
    ) external view returns (bool) {
        //randomness from each polynomial
        bytes32[] memory rs = new bytes32[](polys.length);
        DataStoreKZGMetadata
            memory dskzgMetadata = getDataCommitmentAndMultirevealDegreeAndSymbolBreakdownFromHeader(
                header
            );
        uint256 numProofs = multiRevealProofs.length;
        for (uint256 i = 0; i < numProofs; ) {
            //verify pairing for the commitment to interpolating polynomial
            require(
                validateDisclosureResponse(
                    dskzgMetadata,
                    firstChunkNumber + uint32(i),
                    multiRevealProofs[i].interpolationPoly,
                    multiRevealProofs[i].revealProof,
                    multiRevealProofs[i].zeroPoly,
                    multiRevealProofs[i].zeroPolyProof
                ),
                "Reveal failed due to non 1 pairing"
            );

            // TODO: verify that this check is correct!
            // check that degree of polynomial in the header matches the length of the submitted polynomial
            // i.e. make sure submitted polynomial doesn't contain extra points
            require(
                dskzgMetadata.degree * 32 == polys[i].length,
                "Polynomial must have a 256 bit coefficient for each term"
            );

            //Calculating r, the point at which to evaluate the interpolating polynomial
            rs[i] = keccak256(
                abi.encodePacked(
                    keccak256(polys[i]),
                    multiRevealProofs[i].interpolationPoly.X,
                    multiRevealProofs[i].interpolationPoly.Y
                )
            );
            unchecked {
                ++i;
            }
        }
        //this is the point to open each polynomial at
        uint256 r = uint256(keccak256(abi.encodePacked(rs))) % BN254.FR_MODULUS;
        //this is the offset we add to each polynomial to prevent collision
        //we use array to help with stack
        uint256[2] memory gammaAndGammaPower;
        gammaAndGammaPower[0] =
            uint256(keccak256(abi.encodePacked(rs, uint256(0)))) %
            BN254.FR_MODULUS;
        gammaAndGammaPower[1] = gammaAndGammaPower[0];
        //store I1
        BN254.G1Point memory gammaShiftedCommitmentSum = multiRevealProofs[0]
            .interpolationPoly;
        //store I1(r)
        uint256 gammaShiftedEvaluationSum = linearPolynomialEvaluation(
            polys[0],
            r
        );
        for (uint i = 1; i < multiRevealProofs.length; i++) {
            //gammaShiftedCommitmentSum += gamma^i * Ii
            gammaShiftedCommitmentSum = BN254.plus(
                gammaShiftedCommitmentSum,
                BN254.scalar_mul(
                    multiRevealProofs[i].interpolationPoly,
                    gammaAndGammaPower[1]
                )
            );
            //gammaShiftedEvaluationSum += gamma^i * Ii(r)
            uint256 eval = linearPolynomialEvaluation(polys[i], r);
            gammaShiftedEvaluationSum = gammaShiftedEvaluationSum = addmod(
                gammaShiftedEvaluationSum,
                mulmod(gammaAndGammaPower[1], eval, BN254.FR_MODULUS),
                BN254.FR_MODULUS
            );
            // gammaPower = gamma^(i+1)
            gammaAndGammaPower[1] = mulmod(
                gammaAndGammaPower[0],
                gammaAndGammaPower[1],
                BN254.FR_MODULUS
            );
        }

        return
            openPolynomialAtPoint(
                gammaShiftedCommitmentSum,
                polyEquivalenceProof,
                r,
                gammaShiftedEvaluationSum
            );
    }

    //evaluates the given polynomial "poly" at value "r" and returns the result
    function linearPolynomialEvaluation(bytes calldata poly, uint256 r)
        public
        pure
        returns (uint256)
    {
        uint256 sum;
        uint256 length = poly.length;
        uint256 rPower = 1;
        for (uint i = 0; i < length; ) {
            uint256 coefficient = uint256(bytes32(poly[i:i + 32]));
            sum = addmod(sum, mulmod(coefficient, rPower, BN254.FR_MODULUS), BN254.FR_MODULUS);
            rPower = mulmod(rPower, r, BN254.FR_MODULUS);
            i += 32;
        }
        return sum;
    }

    /// @notice This function evaluates the zero polynomial (x^ChunkLenE - (⍵^index * φ)^{ChunkLenE}) at point alpha
    ///         Recall that `φ = ⍵^NumSysE` and `⍵` is the `NumSysE * ChunkLenE`-th root of unity. 
    ///         So, we are now left with evaluating `⍵^{index * ChunkLenE}`.  
    function constructZeroPolyEval(uint256 index, uint32 chunkLenE, uint32 numNodeE, uint256 alpha) public returns (uint256) {
        // getting the mapping of the coset index
        index = mulmod(chunkLenE, (reverseBitsLimited(numNodeE, uint32(index))), BN254.FR_MODULUS);

        // uint256 modulus = BN254.FR_MODULUS;
        // uint256 omegaPower = BN254.OMEGA;
        uint256 omegaPower;

        // computing ⍵^{index * ChunkLenE}
        omegaPower = BN254.expMod(BN254.OMEGA, index, BN254.FR_MODULUS);
 
        // computing alpha^ChunkLenE
        alpha = BN254.expMod(alpha, chunkLenE, BN254.FR_MODULUS);

        return addmod(alpha, BN254.FR_MODULUS - omegaPower,  BN254.FR_MODULUS);
    }
    
}
