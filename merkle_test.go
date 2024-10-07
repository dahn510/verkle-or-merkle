package main

import (
	"testing"
)

func BenchmarkMerkleConstruct300Byte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := benchCreateMerkleTree(300, b); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkVerkleConstruct300Byte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := benchCreateMerkleTree(300, b); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkMerkleConstruct500Byte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := benchCreateMerkleTree(500, b); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkVerkleConstruct500Byte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := benchCreateMerkleTree(500, b); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkMerkleConstruct10000Byte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := benchCreateMerkleTree(10000, b); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkVerkleConstruct10000Byte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := benchCreateMerkleTree(10000, b); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkMerkleConstruct1000000Byte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := benchCreateMerkleTree(1000000, b); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkVerkleConstruct1000000Byte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := benchCreateMerkleTree(1000000, b); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkJustCreateMerkleProof300Byte(b *testing.B) {
	benchJustCreateMerkleProof(300, b)
}

func BenchmarkJustCreateVerkleProof300Byte(b *testing.B) {
	benchJustCreateVerkleProof(300, b)
}

func BenchmarkJustCreateMerkleProof10000Byte(b *testing.B) {
	benchJustCreateMerkleProof(10000, b)
}

func BenchmarkJustCreateVerkleProof10000Byte(b *testing.B) {
	benchJustCreateVerkleProof(10000, b)
}

func BenchmarkCreateAndSerializeMerkleProof300Byte(b *testing.B) {
	if err := benchCreateAndSerializeMerkleProof(300, b); err != nil {
		b.Fatal(err)
	}
}
func BenchmarkCreateAndSerializeVerkleProof300Byte(b *testing.B) {
	if err := benchCreateAndSerializeVerkleProof(300, b); err != nil {
		b.Fatal(err)
	}
}

func BenchmarkCreateAndSerializeMerkleProof100000Byte(b *testing.B) {
	if err := benchCreateAndSerializeMerkleProof(100000, b); err != nil {
		b.Fatal(err)
	}
}
func BenchmarkCreateAndSerializeVerkleProof100000Byte(b *testing.B) {
	if err := benchCreateAndSerializeVerkleProof(100000, b); err != nil {
		b.Fatal(err)
	}
}

func BenchmarkCreateAndSerializeMerkleProof10000000Byte(b *testing.B) {
	if err := benchCreateAndSerializeMerkleProof(10000000, b); err != nil {
		b.Fatal(err)
	}
}

func BenchmarkCreateAndSerializeVerkleProof10000000Byte(b *testing.B) {
	if err := benchCreateAndSerializeVerkleProof(10000000, b); err != nil {
		b.Fatal(err)
	}
}
