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

package main

import (
	"math/rand"
	"strconv"
	"testing"
)

const COUNT = 10000
const RMAX = 20000

func TestRandomSetGetDel(t *testing.T) {
	tr := newIntStringTreeMap(less)
	kv := make(map[int]string)
	for i := 0; i < COUNT; i++ {
		k := int(rand.Int63n(RMAX))
		v := value(strconv.Itoa(int(rand.Int63n(RMAX))))
		tr.Set(k, v)
		kv[k] = v
		if len(kv) != tr.Count() {
			t.Errorf("[random set get] wrong count, expected %d, got %d", len(kv), tr.Count())
		}
	}
	for k := range kv {
		if v, ok := tr.Get(k); v != kv[k] || !ok {
			t.Errorf("[random set get] wrong returned value, expected %s, got %s", kv[k], v)
		}
		delete(kv, k)
		tr.Del(k)
		if len(kv) != tr.Count() {
			t.Errorf("[random set get] wrong count, expected %d, got %d",
				len(kv), tr.Count())
		}
	}
}

func TestRandomSetIter(t *testing.T) {
	tr := newIntStringTreeMap(less)
	kv := make(map[int]string)
	for i := 0; i < COUNT; i++ {
		k := int(rand.Int63n(RMAX))
		v := value(strconv.Itoa(int(rand.Int63n(RMAX))))
		tr.Set(k, v)
		kv[k] = v
		if len(kv) != tr.Count() {
			t.Errorf("[random set walk] wrong count, expected %d, got %d", len(kv), tr.Count())
		}
	}
	var count int
	for it := tr.Iterator(); it.HasNext(); {
		count++
		k, v := it.Next()
		if kv[k] != v {
			t.Errorf("wrong value, expected %s, got %s", kv[k], v)
		}
	}
	if count != len(kv) {
		t.Errorf("[random set walk] direct order: wrong walking count, expected, %d, got %d", len(kv), count)
	}
	count = 0
	for it := tr.Reverse(); it.HasNext(); {
		count++
		k, v := it.Next()
		if kv[k] != v {
			t.Errorf("wrong value, expected %s, got %s", kv[k], v)
		}
	}
	if count != len(kv) {
		t.Errorf("[random set walk] reverse order: wrong walking count, expected, %d, got %d", len(kv), count)
	}
}
