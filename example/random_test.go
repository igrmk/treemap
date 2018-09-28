package main

import (
	"math/rand"
	"reflect"
	"sort"
	"strconv"
	"testing"
)

const NumIters = 10000
const RandMax = 20

func min(kv map[int]string) int {
	if len(kv) == 0 {
		return 0
	}
	var key *int
	for k := range kv {
		if key == nil || k < *key {
			t := k
			key = &t
		}
	}
	return *key
}

func max(kv map[int]string) int {
	if len(kv) == 0 {
		return 0
	}
	var key *int
	for k := range kv {
		if key == nil || k > *key {
			t := k
			key = &t
		}
	}
	return *key
}

type pair struct {
	k int
	v string
}

func kv() []pair {
	var kv []pair
	for i := 0; i < NumIters; i++ {
		k := int(rand.Int63n(RandMax))
		v := value(strconv.Itoa(int(rand.Int63n(RandMax))))
		kv = append(kv, pair{k, v})
	}
	return kv
}

// nolint: gocyclo
func TestRandom(t *testing.T) {
	tr := newIntStringTreeMap(less)
	mp := make(map[int]string)
	kvs := kv()
	for i, kv := range kvs {
		k, v := kv.k, kv.v
		expV, expOK := mp[k]
		if got, gotOK := tr.Get(k); got != expV || gotOK != expOK {
			t.Errorf("wrong returned value, expected %s, got %s", expV, got)
		}

		if i%3 == 0 && (i/200)%2 == 0 {
			tr.Set(k, v)
			mp[k] = v
		} else {
			delete(mp, k)
			tr.Del(k)
		}

		if len(mp) != tr.Count() {
			t.Errorf("wrong count, expected %d, got %d", len(mp), tr.Count())
			return
		}

		var gotints []int
		for it := tr.Iterator(); it.Valid(); it.Next() {
			gotints = append(gotints, it.Key())
		}

		var expints []int
		for k := range mp {
			expints = append(expints, k)
		}
		sort.Ints(expints)

		if !reflect.DeepEqual(gotints, expints) {
			t.Errorf("wrong keys, expected %v, got %v", expints, gotints)
			return
		}

		exp := min(mp)
		var got int
		if it := tr.Iterator(); it.Valid() {
			got = it.Key()
		}
		if exp != got {
			t.Errorf("wrong min, expected %d, got %d", exp, got)
		}

		exp = max(mp)
		got = 0
		if it := tr.Reverse(); it.Valid() {
			got = it.Key()
		}
		if exp != got {
			t.Errorf("wrong max, expected %d, got %d", exp, got)
		}

		for it := tr.Reverse(); it.Valid(); it.Next() {
			if mp[it.Key()] != it.Value() {
				t.Errorf("wrong value, expected %s, got %s", mp[it.Key()], it.Value())
			}
		}
	}
}
