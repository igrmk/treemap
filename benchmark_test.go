//
// Copyright (c) 2018 Igor Mikushkin <igor.mikushkin@gmail.com>.
// All rights reserved. This program is free software. It comes without
// any warranty, to the extent permitted by applicable law. You can
// redistribute it and/or modify it under the terms of the Do What
// The Fuck You Want To Public License, Version 2, as published by
// Sam Hocevar. See LICENSE file for more details or see below.
//

//
//        DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
//                    Version 2, December 2004
//
// Copyright (C) 2004 Sam Hocevar <sam@hocevar.net>
//
// Everyone is permitted to copy and distribute verbatim or modified
// copies of this license document, and changing it is allowed as long
// as the name is changed.
//
//            DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
//   TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION
//
//  0. You just DO WHAT THE FUCK YOU WANT TO.
//

package treemap

import (
	"math/rand"
	"testing"
)

func BenchmarkSeqSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tr := New(less)
		for j := 0; j < MapSize; j++ {
			tr.Set(j, "")
		}
	}
	b.ReportAllocs()
}

func BenchmarkSeqGet(b *testing.B) {
	b.StopTimer()
	tr := New(less)
	for i := 0; i < MapSize; i++ {
		tr.Set(i, "")
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < MapSize; j++ {
			tr.Get(j)
		}
	}
	b.ReportAllocs()
}

func BenchmarkSeqSetDel(b *testing.B) {
	b.StopTimer()
	tr := New(less)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < MapSize; j++ {
			tr.Set(j, "")
		}
		for j := 0; j < MapSize; j++ {
			tr.Del(j)
		}
	}
	b.ReportAllocs()
}

func BenchmarkSeqExs(b *testing.B) {
	b.StopTimer()
	tr := New(less)
	for i := 0; i < MapSize; i++ {
		tr.Set(i, "")
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < MapSize; j++ {
			tr.Contains(j)
		}
	}
	b.ReportAllocs()
}

func BenchmarkSeqMin(b *testing.B) {
	b.StopTimer()
	tr := New(less)
	for i := 0; i < MapSize; i++ {
		tr.Set(i, "")
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tr.Min()
	}
	b.ReportAllocs()
}

func BenchmarkSeqMax(b *testing.B) {
	b.StopTimer()
	tr := New(less)
	for i := 0; i < MapSize; i++ {
		tr.Set(i, "")
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tr.Max()
	}
	b.ReportAllocs()
}

func BenchmarkSeqIter(b *testing.B) {
	b.StopTimer()
	tr := New(less)
	for i := 0; i < MapSize; i++ {
		tr.Set(i, "")
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for it := tr.Iterator(); it.HasNext(); {
			it.Next()
		}
	}
	b.ReportAllocs()
}

// random

func shuffle(ary []Key) {
	for i := range ary {
		j := rand.Intn(i + 1)
		ary[i], ary[j] = ary[j], ary[i]
	}
}

func BenchmarkRndSet(b *testing.B) {
	t := int64(MapSize * 2)
	b.StopTimer()
	ks := make([]Key, MapSize)
	for i := range ks {
		ks[i] = int(rand.Int63n(t))
	}
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
	ks := make([]Key, MapSize)
	t := int64(MapSize * 2)
	for i := range ks {
		k := int(rand.Int63n(t))
		ks[i] = k
		tr.Set(k, "")
	}
	shuffle(ks)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for _, k := range ks {
			tr.Get(k)
		}
	}
	b.ReportAllocs()
}

func BenchmarkRndSetDel(b *testing.B) {
	b.StopTimer()
	ks := make([]Key, MapSize)
	t := int64(MapSize * 2)
	for i := range ks {
		ks[i] = int(rand.Int63n(t))
	}
	shuffle(ks)
	tr := New(less)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for _, k := range ks {
			tr.Set(k, "")
		}
		for _, k := range ks {
			tr.Del(k)
		}
	}
	b.ReportAllocs()
}

func BenchmarkRndExs(b *testing.B) {
	b.StopTimer()
	tr := New(less)
	ks := make([]Key, MapSize)
	t := int64(MapSize * 2)
	for i := range ks {
		ks[i] = int(rand.Int63n(t))
		tr.Set(ks[i], "")
	}
	shuffle(ks)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for _, t := range ks {
			tr.Contains(t)
		}
	}
	b.ReportAllocs()
}

func BenchmarkRndMin(b *testing.B) {
	b.StopTimer()
	tr := New(less)
	t := int64(MapSize * 2)
	for i := 0; i < MapSize; i++ {
		k := int(rand.Int63n(t))
		tr.Set(k, "")
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tr.Min()
	}
	b.ReportAllocs()
}

func BenchmarkRndMax(b *testing.B) {
	b.StopTimer()
	tr := New(less)
	t := int64(MapSize * 2)
	for i := 0; i < MapSize; i++ {
		k := int(rand.Int63n(t))
		tr.Set(k, "")
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tr.Max()
	}
	b.ReportAllocs()
}

func BenchmarkRndIter(b *testing.B) {
	b.StopTimer()
	tr := New(less)
	t := int64(MapSize * 2)
	for i := 0; i < MapSize; i++ {
		k := int(rand.Int63n(t))
		tr.Set(k, "")
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for it := tr.Iterator(); it.HasNext(); {
			it.Next()
		}
	}
	b.ReportAllocs()
}
