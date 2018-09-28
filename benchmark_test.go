package treemap

import (
	"math/rand"
	"testing"
)

func BenchmarkSeqSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tr := New(less)
		for j := 0; j < NumIters; j++ {
			tr.Set(j, "")
		}
	}
	b.ReportAllocs()
}

func BenchmarkSeqGet(b *testing.B) {
	b.StopTimer()
	tr := New(less)
	for i := 0; i < NumIters; i++ {
		tr.Set(i, "")
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < NumIters; j++ {
			tr.Get(j)
		}
	}
	b.ReportAllocs()
}

func BenchmarkSeqIter(b *testing.B) {
	b.StopTimer()
	tr := New(less)
	for i := 0; i < NumIters; i++ {
		tr.Set(i, "")
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for it := tr.Iterator(); it.Valid(); it.Next() {
		}
	}
	b.ReportAllocs()
}

func randoms() []Key {
	ks := make([]Key, NumIters)
	for i := range ks {
		ks[i] = int(rand.Int63n(NumIters * 100))
	}
	return ks
}

func BenchmarkRndSet(b *testing.B) {
	b.StopTimer()
	ks := randoms()
	tr := New(less)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for _, k := range ks {
			tr.Set(k, "")
		}
		tr.Clear()
	}
	b.ReportAllocs()
}

func BenchmarkRndGet(b *testing.B) {
	b.StopTimer()
	tr := New(less)
	ks := randoms()
	for _, k := range ks {
		tr.Set(k, "")
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for _, k := range ks {
			tr.Get(k)
		}
	}
	b.ReportAllocs()
}

func BenchmarkRndIter(b *testing.B) {
	b.StopTimer()
	tr := New(less)
	ks := randoms()
	for _, k := range ks {
		tr.Set(k, "")
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for it := tr.Iterator(); it.Valid(); it.Next() {
		}
	}
	b.ReportAllocs()
}
