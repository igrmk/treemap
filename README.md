TreeMap v2
==========

[![PkgGoDev](https://pkg.go.dev/badge/github.com/igrmk/treemap/v2)](https://pkg.go.dev/github.com/igrmk/treemap/v2)
[![Unlicense](https://img.shields.io/badge/license-Unlicense-brightgreen.svg)](http://unlicense.org/)
[![Build Status](https://api.travis-ci.com/igrmk/treemap.svg?branch=master)](https://app.travis-ci.com/github/igrmk/treemap)
[![Coverage Status](https://coveralls.io/repos/igrmk/treemap/badge.svg?branch=master)](https://coveralls.io/github/igrmk/treemap)
[![GoReportCard](https://goreportcard.com/badge/github.com/igrmk/treemap/v2)](https://goreportcard.com/report/github.com/igrmk/treemap/v2)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)

`TreeMap` is a generic key-sorted map using a red-black tree under the hood.
It requires and relies on [Go 1.18](https://tip.golang.org/doc/go1.18) generics feature.
Iterators are designed after C++.

### Usage

```go
package main

import (
	"fmt"

	"github.com/igrmk/treemap/v2"
)

func main() {
	tr := treemap.New[int, string]()
	tr.Set(1, "World")
	tr.Set(0, "Hello")
	for it := tr.Iterator(); it.Valid(); it.Next() {
		fmt.Println(it.Key(), it.Value())
	}
}

// Output:
// 0 Hello
// 1 World
```

### Install

```bash
go get github.com/igrmk/treemap/v2
```

### Complexity

|              Name              |   Time    |
|:------------------------------:|:---------:|
|             `Set`              | O(log*N*) |
|             `Del`              | O(log*N*) |
|             `Get`              | O(log*N*) |
|           `Contains`           | O(log*N*) |
|             `Len`              |   O(1)    |
|            `Clear`             |   O(1)    |
|            `Range`             | O(log*N*) |
|           `Iterator`           |   O(1)    |
|           `Reverse`            | O(log*N*) |
| Iterate through the entire map |  O(*N*)   |

### Memory usage

TreeMap uses O(*N*) memory.

### TreeMap v1

The previous version of this package used [gotemplate](https://github.com/ncw/gotemplate) library to generate a type specific file in your local directory.
Here is the link to this version [treemap v1](https://github.com/igrmk/treemap/tree/v1.0.0).

### Licensing

Copyright &copy; 2022 igrmk.
This work is free. You can redistribute it and/or modify it under the
terms of the Unlicense. See the LICENSE file for more details.

### Thanks to

[![JetBrains](svg/jetbrains.svg)](https://www.jetbrains.com/?from=treemap)
