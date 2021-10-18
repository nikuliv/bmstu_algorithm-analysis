package tests

import (
	"lab_01/scripts"
	"lab_01/tools"
	"testing"
)

// Recursive method benchmarks.

func BenchmarkRecursiveLen5(b *testing.B) {
	fWord := tools.RandStringBytesMask(5)
	sWord := tools.RandStringBytesMask(5)

	for i := 0; i < b.N; i++ {
		scripts.LevenshteinRecursive(fWord, sWord)
	}
}

func BenchmarkRecursiveLen10(b *testing.B) {
	fWord := tools.RandStringBytesMask(10)
	sWord := tools.RandStringBytesMask(10)

	for i := 0; i < b.N; i++ {
		scripts.LevenshteinRecursive(fWord, sWord)
	}
}

func BenchmarkRecursiveLen15(b *testing.B) {
	fWord := tools.RandStringBytesMask(15)
	sWord := tools.RandStringBytesMask(15)

	for i := 0; i < b.N; i++ {
		scripts.LevenshteinRecursive(fWord, sWord)
	}
}

// LevenshteinMatrix method becnhmarks.

func BenchmarkLevenshteinMatrixLen10(b *testing.B) {
	fWord := tools.RandStringBytesMask(10)
	sWord := tools.RandStringBytesMask(10)

	for i := 0; i < b.N; i++ {
		scripts.LevenshteinMatrix(fWord, sWord)
	}
}

func BenchmarkLevenshteinMatrixLen20(b *testing.B) {
	fWord := tools.RandStringBytesMask(20)
	sWord := tools.RandStringBytesMask(20)

	for i := 0; i < b.N; i++ {
		scripts.LevenshteinMatrix(fWord, sWord)
	}
}

func BenchmarkLevenshteinMatrixLen30(b *testing.B) {
	fWord := tools.RandStringBytesMask(30)
	sWord := tools.RandStringBytesMask(30)

	for i := 0; i < b.N; i++ {
		scripts.LevenshteinMatrix(fWord, sWord)
	}
}

func BenchmarkLevenshteinMatrixLen40(b *testing.B) {
	fWord := tools.RandStringBytesMask(40)
	sWord := tools.RandStringBytesMask(40)

	for i := 0; i < b.N; i++ {
		scripts.LevenshteinMatrix(fWord, sWord)
	}
}

func BenchmarkLevenshteinMatrixLen50(b *testing.B) {
	fWord := tools.RandStringBytesMask(50)
	sWord := tools.RandStringBytesMask(50)

	for i := 0; i < b.N; i++ {
		scripts.LevenshteinMatrix(fWord, sWord)
	}
}

func BenchmarkLevenshteinMatrixLen100(b *testing.B) {
	fWord := tools.RandStringBytesMask(100)
	sWord := tools.RandStringBytesMask(100)

	for i := 0; i < b.N; i++ {
		scripts.LevenshteinMatrix(fWord, sWord)
	}
}

// LevenshteinRecursiveCache method benchmarks.

func BenchmarkLevenshteinCacheLen10(b *testing.B) {
	fWord := tools.RandStringBytesMask(10)
	sWord := tools.RandStringBytesMask(10)

	for i := 0; i < b.N; i++ {
		scripts.LevenshteinRecursiveCache(fWord, sWord)
	}
}

func BenchmarkLevenshteinCacheLen20(b *testing.B) {
	fWord := tools.RandStringBytesMask(20)
	sWord := tools.RandStringBytesMask(20)

	for i := 0; i < b.N; i++ {
		scripts.LevenshteinRecursiveCache(fWord, sWord)
	}
}

func BenchmarkLevenshteinCacheLen30(b *testing.B) {
	fWord := tools.RandStringBytesMask(30)
	sWord := tools.RandStringBytesMask(30)

	for i := 0; i < b.N; i++ {
		scripts.LevenshteinRecursiveCache(fWord, sWord)
	}
}

func BenchmarkLevenshteinCacheLen40(b *testing.B) {
	fWord := tools.RandStringBytesMask(40)
	sWord := tools.RandStringBytesMask(40)

	for i := 0; i < b.N; i++ {
		scripts.LevenshteinRecursiveCache(fWord, sWord)
	}
}


func BenchmarkLevenshteinCacheLen50(b *testing.B) {
	fWord := tools.RandStringBytesMask(50)
	sWord := tools.RandStringBytesMask(50)

	for i := 0; i < b.N; i++ {
		scripts.LevenshteinRecursiveCache(fWord, sWord)
	}
}

func BenchmarkLevenshteinCacheLen100(b *testing.B) {
	fWord := tools.RandStringBytesMask(100)
	sWord := tools.RandStringBytesMask(100)

	for i := 0; i < b.N; i++ {
		scripts.LevenshteinRecursiveCache(fWord, sWord)
	}
}

// DamerauLevenshtein method benchmarks.

func BenchmarkDamerauLevenshteinLen5(b *testing.B) {
	fWord := tools.RandStringBytesMask(5)
	sWord := tools.RandStringBytesMask(5)

	for i := 0; i < b.N; i++ {
		scripts.DamerauLevenshtein(fWord, sWord)
	}
}

func BenchmarkDamerauLevenshteinLen10(b *testing.B) {
	fWord := tools.RandStringBytesMask(10)
	sWord := tools.RandStringBytesMask(10)

	for i := 0; i < b.N; i++ {
		scripts.DamerauLevenshtein(fWord, sWord)
	}
}

func BenchmarkDamerauLevenshteinLen15(b *testing.B) {
	fWord := tools.RandStringBytesMask(15)
	sWord := tools.RandStringBytesMask(15)

	for i := 0; i < b.N; i++ {
		scripts.DamerauLevenshtein(fWord, sWord)
	}
}
