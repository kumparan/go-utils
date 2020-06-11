package utils

import "testing"

func BenchmarkGenerateRandomString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = GenerateRandomString(100)
	}
}

func BenchmarkGenerateRandomAlphanumeric(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GenerateRandomAlphanumeric(100)
	}
}

func BenchmarkGenerateRandomStringURLSafe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = GenerateRandomStringURLSafe(100)
	}
}
