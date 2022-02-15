package treemap

import (
	"math/rand"
	"reflect"
	"sort"
	"strconv"
	"testing"
)

const NumIterations = 10000
const RandMax = 40

func TestRandom(t *testing.T) {
	tr := New[int, string]()
	mp := make(map[int]string)
	kvs := testRandomData()
	for i, kv := range kvs {
		k, v := kv.k, kv.v
		exp, expOK := mp[k]
		if actual, actualOK := tr.Get(k); actual != exp || actualOK != expOK {
			t.Errorf("wrong returned value, expected %s, actual %s", exp, actual)
		}

		if i%3 == 0 && (i/200)%2 == 0 {
			tr.Set(k, v)
			mp[k] = v
		} else {
			delete(mp, k)
			tr.Del(k)
		}

		if len(mp) != tr.Len() {
			t.Errorf("wrong count, expected %d, actual %d", len(mp), tr.Len())
			return
		}

		testKeys(t, mp, tr)
		testMinMax(t, mp, tr)
		testReverse(t, mp, tr)

		if !treeInvariant(tr.endNode.left) {
			t.Errorf("invariant error")
		}
	}
}

func min(kv map[int]string) *int {
	var key *int
	for k := range kv {
		if key == nil || k < *key {
			temp := k
			key = &temp
		}
	}
	return key
}

func max(kv map[int]string) *int {
	var key *int
	for k := range kv {
		if key == nil || k > *key {
			temp := k
			key = &temp
		}
	}
	return key
}

type pair struct {
	k int
	v string
}

func testRandomData() []pair {
	var kv []pair
	for i := 0; i < NumIterations; i++ {
		k := int(rand.Int63n(RandMax))
		v := strconv.Itoa(int(rand.Int63n(RandMax)))
		kv = append(kv, pair{k, v})
	}
	return kv
}

func testKeys(t *testing.T, mp map[int]string, tr *TreeMap[int, string]) {
	var actualKeys []int
	for it := tr.Iterator(); it.Valid(); it.Next() {
		actualKeys = append(actualKeys, it.Key())
	}

	var expKeys []int
	for k := range mp {
		expKeys = append(expKeys, k)
	}
	sort.Ints(expKeys)

	if !reflect.DeepEqual(actualKeys, expKeys) {
		t.Errorf("wrong keys, expected %v, actual %v", expKeys, actualKeys)
	}
}

func testMinMax(t *testing.T, mp map[int]string, tr *TreeMap[int, string]) {
	exp := min(mp)
	var actual *int
	if it := tr.Iterator(); it.Valid() {
		temp := it.Key()
		actual = &temp
	}
	if (exp == nil) != (actual == nil) {
		t.Errorf("wrong min")
	} else if exp != nil && actual != nil && *exp != *actual {
		t.Errorf("wrong min, expected %d, actual %d", *exp, *actual)
	}

	exp = max(mp)
	actual = nil
	if it := tr.Reverse(); it.Valid() {
		temp := it.Key()
		actual = &temp
	}
	if (exp == nil) != (actual == nil) {
		t.Errorf("wrong max")
	} else if exp != nil && actual != nil && *exp != *actual {
		t.Errorf("wrong max, expected %d, actual %d", *exp, *actual)
	}
}

func testReverse(t *testing.T, mp map[int]string, tr *TreeMap[int, string]) {
	for it := tr.Reverse(); it.Valid(); it.Next() {
		if mp[it.Key()] != it.Value() {
			t.Errorf("wrong value, expected %s, actual %s", mp[it.Key()], it.Value())
		}
	}
}
