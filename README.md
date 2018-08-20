TreeMap
=======

[![GoDoc](https://godoc.org/github.com/igrmk/treemap?status.svg)](https://godoc.org/github.com/igrmk/treemap)
[![WTFPL License](https://img.shields.io/badge/license-wtfpl-blue.svg)](http://www.wtfpl.net/about/)
[![Build Status](https://travis-ci.org/igrmk/treemap.svg)](https://travis-ci.org/igrmk/treemap)
[![Coverage Status](https://coveralls.io/repos/igrmk/treemap/badge.svg?branch=master)](https://coveralls.io/r/igrmk/treemap?branch=master)
[![GoReportCard](http://goreportcard.com/badge/igrmk/treemap)](http://goreportcard.com/report/igrmk/treemap)

Generic `TreeMap` uses red-black tree under the hood.
This is gotemplate ready package.
You can use it as a template to generate `TreeMap` with specific `Key` and `Value` types.
See example folder for generating `TreeMap<int, string>`.
The package is based on github.com/logrusorgru/ebony due to outstanding test coverage.
It is not thread safe.

### Usage

```go
package main

import "fmt"

//go:generate gotemplate "github.com/igrmk/treemap" "intStringTreeMap(int, string)"

func less(x int, y int) bool { return x < y }

func main() {
	tr := newIntStringTreeMap(less)
	tr.Set(0, "Hello")
	tr.Set(1, "World")

	for it := tr.Iterator(); it.HasNext(); {
		k, v := it.Next()
		fmt.Println(k, v)
	}
}
```

### Complexity

| Name        | Time      |
|:-----------:|:---------:|
| `Set`       | O(log*N*) |
| `Del`       | O(log*N*) |
| `Get`       | O(log*N*) |
| `Contains`  | O(log*N*) |
| `Count`     | O(1)      |
| `Max`       | O(log*N*) |
| `Min`       | O(log*N*) |
| `Clear`     | O(1)      |
| `Range`     | O(log*N*) |
| Iteration   | O(*N*)    |

### Memory usage

TreeMap uses O(N) memory.

### Install

Get or update

```bash
go get github.com/igrmk/treemap
```

### Licensing

Copyright &copy; 2018 Igor Mikushkin <igor.mikushkin@gmail.com>
This work is free. You can redistribute it and/or modify it under the
terms of the Do What The Fuck You Want To Public License, Version 2,
as published by Sam Hocevar. See the LICENSE file for more details.
