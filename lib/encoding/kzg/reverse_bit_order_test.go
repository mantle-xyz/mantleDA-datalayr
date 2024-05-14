package kzg

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseBitOrder(t *testing.T) {
	for s := 2; s < 2048; s *= 2 {
		t.Run(fmt.Sprintf("size_%d", s), func(t *testing.T) {
			data := make([]uint32, s, s)
			for i := 0; i < s; i++ {
				data[i] = uint32(i)
			}

			reverseBitOrder(uint32(s), func(i, j uint32) {
				data[i], data[j] = data[j], data[i]
			})

			for i := 0; i < s; i++ {
				assert.Equal(t, reverseBitsLimited(uint32(s), uint32(i)), data[i], "bad reversal at %d", i)
				expected := fmt.Sprintf("%0"+fmt.Sprintf("%d", s)+"b", i)
				got := fmt.Sprintf("%0"+fmt.Sprintf("%d", s)+"b", data[i])
				assert.Equal(t, len(expected), len(got), "bad length: %d, expected %d", len(got), len(expected))

				for j := 0; j < len(expected); j++ {
					// TODO: add check
				}
			}
		})
	}
}

func TestRevBitorderBitIndex(t *testing.T) {
	for i := 0; i < 32; i++ {
		got := bitIndex(uint32(1 << i))
		assert.Equal(t, got, uint8(i), "bit index %d is wrong: %d", i, got)
	}
}

func TestReverseBits(t *testing.T) {
	rng := rand.New(rand.NewSource(1234))
	for i := 0; i < 10000; i++ {
		v := rng.Uint32()
		expected := revStr(fmt.Sprintf("%032b", v))
		out := reverseBits(v)
		got := fmt.Sprintf("%032b", out)
		assert.Equal(t, expected, got, "bit mismatch: expected: %s, got: %s ", expected, got)
	}
}

func revStr(v string) string {
	out := make([]byte, len(v))
	for i := 0; i < len(v); i++ {
		out[i] = v[len(v)-1-i]
	}
	return string(out)
}
