TreeMap
=======

This repository is fully merged into [gotemplate](https://github.com/ncw/gotemplate).

[![GoDoc](https://godoc.org/github.com/igrmk/treemap?status.svg)](https://godoc.org/github.com/igrmk/treemap)
[![Unlicense](https://img.shields.io/badge/license-Unlicense-blue.svg)](http://unlicense.org/)
[![Build Status](https://travis-ci.org/igrmk/treemap.svg?branch=master)](https://travis-ci.org/igrmk/treemap)
[![Coverage Status](https://coveralls.io/repos/igrmk/treemap/badge.svg?branch=master)](https://coveralls.io/r/igrmk/treemap?branch=master)
[![GoReportCard](http://goreportcard.com/badge/igrmk/treemap)](http://goreportcard.com/report/igrmk/treemap)

`TreeMap` is a generic key-sorted map. It uses red-black tree under the hood.
This is [gotemplate](https://github.com/ncw/gotemplate) ready package.
You can use it as a template to generate a sorted map with specific key and value types.
See example folder for an illustration of generating a map with `int` key and `string` value.
The package is useful without a generation as well.
It uses `interface{}` type for both a key and a value in this case.
Iterators are designed after C++.

### Usage

```go
package main

import "fmt"

//go:generate gotemplate "github.com/igrmk/treemap" "intStringTreeMap(int, string)"

func less(x, y int) bool { return x < y }

func main() {
	tr := newIntStringTreeMap(less)
	tr.Set(0, "Hello")
	tr.Set(1, "World")

	for it := tr.Iterator(); it.Valid(); it.Next() {
		fmt.Println(it.Key(), it.Value())
	}
}
```

To build it you need to run

```bash
go generate
go build
```

Command `go generate` will generate a file `gotemplate_intStringTreeMap.go` in the same directory.
This file will contain ready to use tree map.

### Install

```bash
go get github.com/ncw/gotemplate/...
go get github.com/igrmk/treemap
```

### Complexity

| Name       | Time      |
|:----------:|:---------:|
| `Set`      | O(log*N*) |
| `Del`      | O(log*N*) |
| `Get`      | O(log*N*) |
| `Contains` | O(log*N*) |
| `Len`      | O(1)      |
| `Clear`    | O(1)      |
| `Range`    | O(log*N*) |
| `Iterator` | O(1)      |
| `Reverse`  | O(log*N*) |
| Iteration  | O(*N*)    |

### Memory usage

TreeMap uses O(*N*) memory.

### Licensing

Copyright &copy; 2018 Igor Mikushkin &lt;igor.mikushkin@gmail.com&gt;.
This work is free. You can redistribute it and/or modify it under the
terms of the Unlicense. See the LICENSE file for more details.

### Thanks to

[![JetBrains](svg/jetbrains.svg)](https://www.jetbrains.com/?from=treemap)
