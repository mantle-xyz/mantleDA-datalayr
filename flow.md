# Message Flow 
![flow](docs/comm-flow.png)
# Components
- Disperser
  - Listener for taking data blob, B
  - Encoder for adding Reed-Solomon code to the data blob, C
  - Data dispersal for sending data to DL node (has payment or approved balances)
    - Disperse(chunk, duration, approve-proof)
  - Listener for collecting data signature from data layr nodes
  - Signatures aggregator for uploading on main chain
  - Main chain tx submitter
- Data Layr(DL) node
  - Listener for dispersal request, verify paid fee and accept connections
  - Verifier for checking dispersed chunk is consistent with the merkle root(commitment) and returned signed commitment
  - Listener for data retrieval 
  - Challenge Handler for verifying if local symbols is a part of commited data
    - verifier for commitment consistent aginst local symbol
		- sign challenge-commit to challenger
  - Key-Value Database (commitment, verified symbol)
- Data Layr smart contract (DLSC)
  - DL node data structure
    - registration, (pubkey, ip, stake)
  - Verify Aggregate/Threshold signature and store commitment if success
  - Challenge process
    - Challenger instantiation, (commitment, challenger stake, expiration); Emit an instantiation event(pubkey)
    - Challenger closing, (aggregate signature)
  - Fee managements, redistribution
- Challenger
  - DLSC Event Listener for logging all commitments
  - Block Assembler for retrieval from DL nodes
  - Commitment Verifier for checking if reassembled block is consistent to hash. 
    - If yes, do nothing; 
    - If no, instantiate DLSC Challenge process on smart contract
    - OffChain, send a challenge request to all DL nodes along with the challenge-commit and all received symbols
    - OffChain, aggreate signatures from all DL nodes
    - If sufficient signatures are accumulated, send to smart contract, close 
