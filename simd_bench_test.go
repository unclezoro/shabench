package shabench

import (
	sha2562 "crypto/sha256"
	"github.com/minio/sha256-simd"
	"hash"
	"testing"
)

func benchmarkSize(b *testing.B, size int,bench hash.Hash) {
	var buf = make([]byte, size)
	b.SetBytes(int64(size))
	sum := make([]byte, bench.Size())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bench.Reset()
		bench.Write(buf[:size])
		bench.Sum(sum[:0])
	}
}


func BenchmarkHash(b *testing.B) {
	var ssidSha = sha256.New()
	var crypSha = sha2562.New()

	sizes:= []struct{
		n string
		s int
		b hash.Hash
	}{
		{"simd: 200bytes",200, ssidSha},
		{"simd: 400bytes",400, ssidSha},
		{"simd: 800bytes",800,ssidSha},
		{"simd: 1000bytes",1000,ssidSha},
		{"simd: 1600bytes",1600,ssidSha},
		{"crypSha: 200bytes",200, crypSha},
		{"crypSha: 400bytes",400, crypSha},
		{"crypSha: 800bytes",800,crypSha},
		{"crypSha: 1000bytes",1000,crypSha},
		{"crypSha: 1600bytes",1600,crypSha},
	}
	for _,s:=range sizes{
		b.Run(s.n, func(b *testing.B) {
			benchmarkSize(b,s.s,s.b)
		})
	}
}

