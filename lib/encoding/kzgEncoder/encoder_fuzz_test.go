package kzgEncoder_test

import (
	"context"
	"testing"

	kzgRs "github.com/Layr-Labs/datalayr/lib/encoding/kzgEncoder"
	"github.com/stretchr/testify/assert"
)

func FuzzOnlySystematic(f *testing.F) {

	f.Add(GETTYSBURG_ADDRESS_BYTES)
	f.Fuzz(func(t *testing.T, input []byte) {

		group, _ := kzgRs.NewKzgEncoderGroup(kzgConfig)
		enc, err := group.NewKzgEncoder(10, 3, uint64(len(input)))
		if err != nil {
			t.Errorf("Error making rs: %q", err)
		}

		//encode the data
		_, _, frames, _, err := enc.EncodeBytes(context.Background(), input)

		for _, frame := range frames {
			assert.NotEqual(t, len(frame.Coeffs), 0)
		}

		if err != nil {
			t.Errorf("Error Encoding:\n Data:\n %q \n Err: %q", input, err)
		}

		//sample the correct systematic frames
		samples, indices := sampleFrames(frames, uint64(len(frames)))

		data, err := enc.DecodeSafe(samples, indices, uint64(len(input)))
		if err != nil {
			t.Errorf("Error Decoding:\n Data:\n %q \n Err: %q", input, err)
		}
		assert.Equal(t, input, data, "Input data was not equal to the decoded data")
	})
}
