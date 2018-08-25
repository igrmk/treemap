package main

import "fmt"

//go:generate gotemplate "github.com/igrmk/treemap" "intStringTreeMap(int, string)"

func less(x, y int) bool { return x < y }

func main() {
	tr := newIntStringTreeMap(less)
	tr.Set(0, "Hello")
	tr.Set(1, "World")

	for it := tr.Iterator(); it.HasNext(); {
		k, v := it.Next()
		fmt.Println(k, v)
	}
}
