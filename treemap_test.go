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
	"testing"
)

func less(x Key, y Key) bool { return x.(int) < y.(int) }

func value(x string) Value { return Value(x) }

func TestNew(t *testing.T) {
	tr := New(less)
	if tr.count != 0 {
		t.Error("[new] count != 0")
	}
	if tr.root != sentinel {
		t.Error("[new] root != sentinel")
	}
}

func TestSet(t *testing.T) {
	x := value("x")
	tr := New(less)
	tr.Set(0, x)
	if tr.root.key != 0 {
		t.Errorf("[set] wrong id, expected 0, got %d", tr.root.key)
	}
	if v := tr.root.value; v != x {
		t.Errorf(
			"[set] wrong returned value, expected '%s', got '%s'", x, v)
	}
	if tr.count != 1 {
		t.Errorf("[set] wrong count, expected 1, got %d", tr.count)
	}
}

func TestDel(t *testing.T) {
	x := "x"
	tr := New(less)
	tr.Set(0, value(x))
	tr.Del(0)
	if tr.count != 0 {
		t.Errorf("[del] wrong count after del, expected 0, got %d", tr.count)
	}
	if tr.root != sentinel {
		t.Error("[del] wrong tree state after del")
	}
}

func TestGet(t *testing.T) {
	x := value("x")
	tr := New(less)
	tr.Set(0, x)
	v, ok := tr.Get(0)
	if v != x || !ok {
		t.Errorf("[get] wrong returned value, expected 'x', got '%s'", v)
	}
	if tr.count != 1 {
		t.Errorf("[get] wrong count, expected 1, got %d", tr.count)
	}
	if v, ok := tr.Get(579); v != nil || ok {
		t.Errorf("[get] wrong returned value, expected nil, got '%v'", v)
	}
	if tr.count != 1 {
		t.Errorf("[get] wrong count, expected 1, got %d", tr.count)
	}
}

func TestExist(t *testing.T) {
	x := value("x")
	tr := New(less)
	tr.Set(0, x)
	val := tr.Contains(0)
	if !val {
		t.Error("[exist] existing is not exist")
	}
	val = tr.Contains(12)
	if val {
		t.Error("[exist] not existing is exist")
	}
}

func TestCount(t *testing.T) {
	x := value("x")
	tr := New(less)
	if tr.Count() != 0 {
		t.Errorf("[count] wrong count, expected 0, got %d", tr.Count())
	}
	tr.Set(0, x)
	if tr.Count() != 1 {
		t.Errorf("[count] wrong count, expected 1, got %d", tr.Count())
	}
	tr.Set(1, x)
	if tr.Count() != 2 {
		t.Errorf("[count] wrong count, expected 2, got %d", tr.Count())
	}
	tr.Del(1)
	if tr.Count() != 1 {
		t.Errorf("[count] wrong count, expected 1, got %d", tr.Count())
	}
	tr.Del(0)
	if tr.Count() != 0 {
		t.Errorf("[count] wrong count, expected 0, got %d", tr.Count())
	}
}

func TestClear(t *testing.T) {
	tr := New(less)
	tr.Set(0, "x")
	tr.Set(1, "y")
	tr.Set(2, "z")
	tr.Clear()
	if tr.count != 0 {
		t.Error("[empty] count != 0")
	}
	if tr.root != sentinel {
		t.Error("[empty] root != sentinel")
	}
}

func TestMax(t *testing.T) {
	max := value("max")
	maxi := 6
	tr := New(less)
	tr.Set(0, "x")
	tr.Set(1, "y")
	tr.Set(2, "z")
	tr.Set(maxi, max)
	tr.Set(3, "m")
	tr.Set(4, "n")
	tr.Set(5, "o")
	i, v := tr.Max()
	if i != maxi {
		t.Errorf("[max] wrong index of min, expected %d, got %d", maxi, i)
	}
	if v != max {
		t.Errorf(
			"[max] wrong returned value, expected '%s', got '%s'",
			max, v)
	}
}

func TestMin(t *testing.T) {
	min := value("min")
	mini := 0
	tr := New(less)
	tr.Set(1, "x")
	tr.Set(2, "y")
	tr.Set(3, "z")
	tr.Set(mini, min)
	tr.Set(4, "m")
	tr.Set(5, "n")
	tr.Set(6, "o")
	i, v := tr.Min()
	if i != mini {
		t.Errorf("[min] wrong index of min, expected %d, got %d", mini, i)
	}
	if v != min {
		t.Errorf("[min] wrong returned value, expected '%s', got '%s'",
			min, v)
	}
}

func testRange13(tr *TreeMap, t *testing.T) {
	var vls []Value
	for it := tr.Range(1, 3); it.HasNext(); {
		_, v := it.Next()
		vls = append(vls, v)
	}
	if len(vls) != 3 {
		t.Errorf("[range] wrong range length, expected 3, got %d", len(vls))
	}
	r13 := []Value{"y", "z", "m"}
	for i := 0; i < len(vls) && i < len(r13); i++ {
		if vls[i] != r13[i] {
			t.Errorf("[range] wrong value, expected '%s', got '%s'",
				r13[i], vls[i])
		}
	}
}

func testRange19(tr *TreeMap, t *testing.T) {
	var vls []Value
	for it := tr.Range(1, 9); it.HasNext(); {
		_, v := it.Next()
		vls = append(vls, v)
	}
	if len(vls) != 4 {
		t.Errorf("[range] wrong range length, expected 4, got %d", len(vls))
	}
	r19 := []Value{"y", "z", "m", "n"}
	for i := 0; i < len(vls) && i < len(r19); i++ {
		if vls[i] != r19[i] {
			t.Errorf("[range] wrong value, expected '%s', got '%s'",
				r19[i], vls[i])
		}
	}
}

func TestRange(t *testing.T) {
	tr := New(less)
	tr.Set(0, "x")
	tr.Set(1, "y")
	tr.Set(2, "z")
	tr.Set(3, "m")
	tr.Set(4, "n")
	testRange13(tr, t)
	testRange19(tr, t)
}

func TestLowerBound(t *testing.T) {
	tr := New(less)
	it := tr.LowerBound(0)
	if it.HasNext() {
		t.Error("lower bound should not exists")
		return
	}
	tr.Set(2, "a")
	tr.Set(4, "b")
	tr.Set(6, "c")
	tr.Set(8, "d")
	tr.Set(10, "e")
	tr.Set(12, "e")
	tr.Set(14, "e")
	tr.Set(16, "e")
	tr.Set(18, "e")
	tr.Set(20, "e")

	tbl := [][2]int{
		{0, 2},
		{2, 2},
		{3, 4},
		{4, 4},
		{9, 10},
		{10, 10},
		{11, 12},
		{19, 20},
		{20, 20},
	}

	for _, tb := range tbl {
		it = tr.LowerBound(tb[0])
		if !it.HasNext() {
			t.Error("lower bound should exists")
			return
		}
		if k, _ := it.Next(); k != tb[1] {
			t.Errorf("lower bound should be %v", tb[1])
			return
		}
	}

	it = tr.LowerBound(21)
	if it.HasNext() {
		t.Error("lower bound should not exists")
		return
	}
}

func TestUpperBound(t *testing.T) {
	tr := New(less)
	it := tr.UpperBound(0)
	if it.HasNext() {
		t.Error("lower bound should not exists")
		return
	}
	tr.Set(2, "a")
	tr.Set(4, "b")
	tr.Set(6, "c")
	tr.Set(8, "d")
	tr.Set(10, "e")
	tr.Set(12, "e")
	tr.Set(14, "e")
	tr.Set(16, "e")
	tr.Set(18, "e")
	tr.Set(20, "e")

	tbl := [][2]int{
		{0, 2},
		{2, 4},
		{3, 4},
		{4, 6},
		{9, 10},
		{10, 12},
		{11, 12},
		{19, 20},
	}

	for _, tb := range tbl {
		it = tr.UpperBound(tb[0])
		if !it.HasNext() {
			t.Error("lower bound should exists")
			return
		}
		if k, _ := it.Next(); k != tb[1] {
			t.Errorf("lower bound should be %v", tb[1])
			return
		}
	}

	it = tr.UpperBound(20)
	if it.HasNext() {
		t.Error("lower bound should not exists")
		return
	}
	it = tr.UpperBound(21)
	if it.HasNext() {
		t.Error("lower bound should not exists")
		return
	}
}

func TestEmptyRange(t *testing.T) {
	tr := New(less)
	tr.Set(0, "x")
	tr.Set(1, "y")
	tr.Set(2, "z")
	tr.Set(3, "m")
	tr.Set(4, "n")
	if rng := tr.Range(5, 10); rng.HasNext() {
		t.Error("[empty range] range should be empty")
	}
}

func TestDelNil(t *testing.T) {
	x := "x"
	tr := New(less)
	tr.Set(0, value(x))
	tr.Del(1)
	if tr.count != 1 {
		t.Errorf("[del nil] wrong count after del, expected 1, got %d",
			tr.count)
	}
}

func TestIteration(t *testing.T) {
	kvs := []struct {
		key   Key
		value Value
	}{
		{0, "a"},
		{1, "b"},
		{2, "c"},
		{3, "d"},
		{4, "e"},
	}
	tr := New(less)
	for _, kv := range kvs {
		tr.Set(kv.key, kv.value)
	}
	count := 0
	for it := tr.Iterator(); it.HasNext(); {
		k, v := it.Next()
		if kvs[count].key != k || kvs[count].value != v {
			t.Errorf("expected %v, %s, got %v, %s", kvs[count].key, kvs[count].value, k, v)
		}
		count++
	}
	for it := tr.Reverse(); it.HasNext(); {
		count--
		k, v := it.Next()
		if kvs[count].key != k || kvs[count].value != v {
			t.Errorf("expected %v, %s, got %v, %s", kvs[count].key, kvs[count].value, k, v)
		}
	}
}

func TestOutOfBoundsIteration(t *testing.T) {
	tr := New(less)
	tr.Set(0, "a")
	tr.Set(1, "b")
	tr.Set(2, "c")
	tr.Set(3, "d")
	tr.Set(4, "e")
	it := tr.Iterator()
	for it.HasNext() {
		it.Next()
	}
	defer func() {
		if r := recover(); r == nil {
			t.Error("should have panicked!")
		}
	}()
	it.Next()
}

func TestRangeSingle(t *testing.T) {
	tr := New(less)
	tr.Set(0, "a")
	tr.Set(1, "b")
	tr.Set(2, "c")
	visited := false
	for it := tr.Range(1, 1); it.HasNext(); {
		_, v := it.Next()
		if visited || v != "b" {
			t.Error("only single element 'b' should be found")
		}
		visited = true
	}
	if !visited {
		t.Error("single element 'b' should be found")
	}
}
