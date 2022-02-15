package treemap

import (
	"math/rand"
	"testing"
)

func BenchmarkSeqSet(b *testing.B) {
	tr := New[int, string]()
	for i := 0; i < b.N; i++ {
		for j := 0; j < NumIterations; j++ {
			tr.Set(j, "")
		}
		tr.Clear()
	}
	b.ReportAllocs()
}

func BenchmarkSeqGet(b *testing.B) {
	tr := New[int, string]()
	for i := 0; i < NumIterations; i++ {
		tr.Set(i, "")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tr.Get(i % NumIterations)
	}
	b.ReportAllocs()
}

func BenchmarkSeqIter(b *testing.B) {
	tr := New[int, string]()
	for i := 0; i < NumIterations; i++ {
		tr.Set(i, "")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for it := tr.Iterator(); it.Valid(); it.Next() {
		}
	}
	b.ReportAllocs()
}

func BenchmarkRndSet(b *testing.B) {
	keys, _ := benchmarksRandomData()
	tr := New[int, string]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, k := range keys {
			tr.Set(k, "")
		}
		tr.Clear()
	}
	b.ReportAllocs()
}

func BenchmarkRndGet(b *testing.B) {
	tr := New[int, string]()
	keys, max := benchmarksRandomData()
	for _, k := range keys {
		tr.Set(k, "")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tr.Get(i % max)
	}
	b.ReportAllocs()
}

func BenchmarkRndIter(b *testing.B) {
	tr := New[int, string]()
	keys, _ := benchmarksRandomData()
	for _, k := range keys {
		tr.Set(k, "")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for it := tr.Iterator(); it.Valid(); it.Next() {
		}
	}
	b.ReportAllocs()
}

func benchmarksRandomData() ([]int, int) {
	keys := make([]int, NumIterations)
	max := NumIterations * 100
	for i := range keys {
		keys[i] = int(rand.Int63n(int64(max)))
	}
	return keys, max
}
