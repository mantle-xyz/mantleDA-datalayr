package kzg

import (
	"fmt"
	bls "github.com/Layr-Labs/datalayr/common/crypto/go-kzg-bn254/bn254"

	//"github.com/protolambda/go-kzg/bls"
	"testing"
)

func benchFFT(scale uint8, b *testing.B) {
	fs := NewFFTSettings(scale)
	data := make([]bls.Fr, fs.MaxWidth, fs.MaxWidth)
	for i := uint64(0); i < fs.MaxWidth; i++ {
		bls.CopyFr(&data[i], bls.RandomFr())
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		out, err := fs.FFT(data, false)
		if err != nil {
			b.Fatal(err)
		}
		if len(out) != len(data) {
			panic("output len doesn't match input")
		}
	}
}

func BenchmarkFFTSettings_FFT(b *testing.B) {
	for scale := uint8(4); scale < 16; scale++ {
		b.Run(fmt.Sprintf("scale_%d", scale), func(b *testing.B) {
			benchFFT(scale, b)
		})
	}
}
