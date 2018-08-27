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
