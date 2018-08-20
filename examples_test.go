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
	"fmt"
)

func ExampleTreeMap_Set() {
	tr := New(less)
	tr.Set(0, "hello")
	v, _ := tr.Get(0)
	fmt.Println(v)
	// Output:
	// hello
}

func ExampleTreeMap_Del() {
	tr := New(less)
	tr.Set(0, "hello")
	tr.Del(0)
	fmt.Println(tr.Contains(0))
	// Output:
	// false
}

func ExampleTreeMap_Get() {
	tr := New(less)
	tr.Set(0, "hello")
	v, _ := tr.Get(0)
	fmt.Println(v)
	// Output:
	// hello
}

func ExampleTreeMap_Contains() {
	tr := New(less)
	tr.Set(0, "hello")
	fmt.Println(tr.Contains(0))
	// Output:
	// true
}

func ExampleTreeMap_Count() {
	tr := New(less)
	tr.Set(0, "hello")
	fmt.Println(tr.Count())
	tr.Set(0, "hello")
	fmt.Println(tr.Count())
	tr.Set(1, "hi")
	fmt.Println(tr.Count())
	// Output:
	// 1
	// 1
	// 2
}

func ExampleTreeMap_Min() {
	tr := New(less)
	tr.Set(0, "hello")
	tr.Set(1, "hi")
	fmt.Println(tr.Min())
	// Output:
	// 0 hello
}

func ExampleTreeMap_Max() {
	tr := New(less)
	tr.Set(0, "hello")
	tr.Set(1, "hi")
	fmt.Println(tr.Max())
	// Output:
	// 1 hi
}

func ExampleTreeMap_Clear() {
	tr := New(less)
	tr.Set(0, "hello")
	tr.Set(1, "world")
	tr.Clear()
	fmt.Println(tr.Count())
	// Output:
	// 0
}

func ExampleTreeMap_Iterator() {
	tr := New(less)
	tr.Set(1, "one")
	tr.Set(2, "two")
	tr.Set(3, "three")
	for it := tr.Iterator(); it.HasNext(); {
		key, value := it.Next()
		fmt.Println(key, "-", value)
	}
	// Output:
	// 1 - one
	// 2 - two
	// 3 - three
}

func ExampleTreeMap_Reverse() {
	tr := New(less)
	tr.Set(1, "one")
	tr.Set(2, "two")
	tr.Set(3, "three")
	for it := tr.Reverse(); it.HasNext(); {
		key, value := it.Next()
		fmt.Println(key, "-", value)
	}
	// Output:
	// 3 - three
	// 2 - two
	// 1 - one
}
