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

/*
Package treemap uses red-black tree under the hood.
This is gotemplate ready package.
You can use it as a template to generate TreeMap with specific Key and Value types.

Generating TreeMap with int keys and string values
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
It is not thread safe.
*/
package treemap

import (
	"runtime"
)

type color bool

const (
	red   color = true
	black color = false
)

// template type TreeMap(Key, Value)

// Key is a generic key type of the map
type Key interface{}

// Value is a generic value type of the map
type Value interface{}

type node struct {
	left   *node
	right  *node
	parent *node
	color  color
	key    Key
	value  Value
}

var sentinel = &node{left: nil, right: nil, parent: nil, color: black}

func init() {
	sentinel.left, sentinel.right = sentinel, sentinel
}

// TreeMap is the red-black tree based map
type TreeMap struct {
	root  *node
	count int
	less  func(Key, Key) bool
}

// New creates the new red-black tree based TreeMap.
// Parameter less is a function returning a < b.
func New(less func(a Key, b Key) bool) *TreeMap {
	return &TreeMap{
		root: sentinel,
		less: less,
	}
}

// Set the value. Silently override previous value if exists. This will overwrite the existing value.
// Complexity: O(log N).
func (t *TreeMap) Set(id Key, value Value) {
	t.insertNode(id, value)
}

// Del deletes the value.
// Complexity: O(log N).
func (t *TreeMap) Del(id Key) {
	t.deleteNode(t.findNode(id))
}

// Get retrieves a value from a map for specified key and reports if it exists.
// Complexity: O(log N).
func (t *TreeMap) Get(id Key) (Value, bool) {
	node := t.findNode(id)
	if node == sentinel {
		return node.value, false
	}
	return node.value, true
}

// Contains checks if key exists in a map.
// Complexity: O(log N)
func (t *TreeMap) Contains(id Key) bool {
	return t.findNode(id) != sentinel
}

// Count returns total count of elements in a map.
// Complexity: O(1).
func (t *TreeMap) Count() int {
	return t.count
}

// Clear clears the map.
// Complexity: O(1).
func (t *TreeMap) Clear() *TreeMap {
	t.root = sentinel
	t.count = 0
	runtime.GC()
	return t
}

// Max returns maximum key and associated value.
// Complexity: O(log N).
func (t *TreeMap) Max() (Key, Value) {
	current := t.root
	for current.right != sentinel {
		current = current.right
	}
	return current.key, current.value
}

// Min returns minimum key and associated value.
// Complexity: O(log N).
func (t *TreeMap) Min() (Key, Value) {
	current := t.root
	for current.left != sentinel {
		current = current.left
	}
	return current.key, current.value
}

// Range returns an iterator such that it goes through all the keys in the range [from, to].
// Complexity: O(log N).
func (t *TreeMap) Range(from, to Key) ForwardIterator {
	lower := t.LowerBound(from)
	upper := t.UpperBound(to)
	return ForwardIterator{tree: t, node: lower.node, end: upper.node}
}

// LowerBound returns an iterator such that it goes through all the keys in the range [key, max(key)] by analogy with C++.
// Complexity: O(log N).
func (t *TreeMap) LowerBound(key Key) ForwardIterator {
	node := t.root
	result := sentinel
	if node == sentinel {
		return ForwardIterator{tree: t, node: sentinel, end: sentinel}
	}
	for {
		if t.less(node.key, key) {
			if node.right != sentinel {
				node = node.right
			} else {
				return ForwardIterator{tree: t, node: result, end: sentinel}
			}
		} else {
			result = node
			if node.left != sentinel {
				node = node.left
			} else {
				return ForwardIterator{tree: t, node: result, end: sentinel}
			}
		}
	}
}

// UpperBound returns an iterator such that it goes through all the keys in the range (key, max(key)] by analogy with C++.
// Complexity: O(log N).
func (t *TreeMap) UpperBound(key Key) ForwardIterator {
	node := t.root
	result := sentinel
	if node == sentinel {
		return ForwardIterator{tree: t, node: sentinel, end: sentinel}
	}
	for {
		if !t.less(key, node.key) {
			if node.right != sentinel {
				node = node.right
			} else {
				return ForwardIterator{tree: t, node: result, end: sentinel}
			}
		} else {
			result = node
			if node.left != sentinel {
				node = node.left
			} else {
				return ForwardIterator{tree: t, node: result, end: sentinel}
			}
		}
	}
}

// Iterator returns an iterator for tree map.
// It starts at the one-before-the-start position and goes to the end.
// You can iterate a map at O(N) complexity.
func (t *TreeMap) Iterator() ForwardIterator {
	node := t.root
	for node.left != sentinel {
		node = node.left
	}
	return ForwardIterator{tree: t, node: node, end: sentinel}
}

// Reverse returns a reverse iterator for tree map.
// It starts at the one-past-the-end position and goes to the beginning.
// You can iterate a map at O(N) complexity.
func (t *TreeMap) Reverse() ReverseIterator {
	node := t.root
	for node.right != sentinel {
		node = node.right
	}
	return ReverseIterator{tree: t, node: node, end: sentinel}
}

func (t *TreeMap) rotateLeft(x *node) {
	y := x.right
	x.right = y.left
	if y.left != sentinel {
		y.left.parent = x
	}
	if y != sentinel {
		y.parent = x.parent
	}
	if x.parent != nil {
		if x == x.parent.left {
			x.parent.left = y
		} else {
			x.parent.right = y
		}
	} else {
		t.root = y
	}
	y.left = x
	if x != sentinel {
		x.parent = y
	}
}

func (t *TreeMap) rotateRight(x *node) {
	y := x.left
	x.left = y.right
	if y.right != sentinel {
		y.right.parent = x
	}
	if y != sentinel {
		y.parent = x.parent
	}
	if x.parent != nil {
		if x == x.parent.right {
			x.parent.right = y
		} else {
			x.parent.left = y
		}
	} else {
		t.root = y
	}
	y.right = x
	if x != sentinel {
		x.parent = y
	}
}

func (t *TreeMap) insertFixup(x *node) {
	for x != t.root && x.parent.color == red {
		if x.parent == x.parent.parent.left {
			y := x.parent.parent.right
			if y.color == red {
				x.parent.color = black
				y.color = black
				x.parent.parent.color = red
				x = x.parent.parent
			} else {
				if x == x.parent.right {
					x = x.parent
					t.rotateLeft(x)
				}
				x.parent.color = black
				x.parent.parent.color = red
				t.rotateRight(x.parent.parent)
			}
		} else {
			y := x.parent.parent.left
			if y.color == red {
				x.parent.color = black
				y.color = black
				x.parent.parent.color = red
				x = x.parent.parent
			} else {
				if x == x.parent.left {
					x = x.parent
					t.rotateRight(x)
				}
				x.parent.color = black
				x.parent.parent.color = red
				t.rotateLeft(x.parent.parent)
			}
		}
	}
	t.root.color = black
}

func (t *TreeMap) insertNode(id Key, value Value) {
	current := t.root
	var parent *node
	for current != sentinel {
		if id == current.key {
			current.value = value
			return
		}
		parent = current
		if t.less(id, current.key) {
			current = current.left
		} else {
			current = current.right
		}
	}
	x := &node{
		value:  value,
		parent: parent,
		left:   sentinel,
		right:  sentinel,
		color:  red,
		key:    id,
	}
	if parent != nil {
		if t.less(id, parent.key) {
			parent.left = x
		} else {
			parent.right = x
		}
	} else {
		t.root = x
	}
	t.insertFixup(x)
	t.count++
}

// nolint: gocyclo
func (t *TreeMap) deleteFixup(x *node) {
	for x != t.root && x.color == black {
		if x == x.parent.left {
			w := x.parent.right
			if w.color == red {
				w.color = black
				x.parent.color = red
				t.rotateLeft(x.parent)
				w = x.parent.right
			}
			if w.left.color == black && w.right.color == black {
				w.color = red
				x = x.parent
			} else {
				if w.right.color == black {
					w.left.color = black
					w.color = red
					t.rotateRight(w)
					w = x.parent.right
				}
				w.color = x.parent.color
				x.parent.color = black
				w.right.color = black
				t.rotateLeft(x.parent)
				x = t.root
			}
		} else {
			w := x.parent.left
			if w.color == red {
				w.color = black
				x.parent.color = red
				t.rotateRight(x.parent)
				w = x.parent.left
			}
			if w.right.color == black && w.left.color == black {
				w.color = red
				x = x.parent
			} else {
				if w.left.color == black {
					w.right.color = black
					w.color = red
					t.rotateLeft(w)
					w = x.parent.left
				}
				w.color = x.parent.color
				x.parent.color = black
				w.left.color = black
				t.rotateRight(x.parent)
				x = t.root
			}
		}
	}
	x.color = black
}

// nolint: gocyclo
func (t *TreeMap) deleteNode(z *node) {
	var x, y *node
	if z == nil || z == sentinel {
		return
	}
	if z.left == sentinel || z.right == sentinel {
		y = z
	} else {
		y = z.right
		for y.left != sentinel {
			y = y.left
		}
	}
	if y.left != sentinel {
		x = y.left
	} else {
		x = y.right
	}
	x.parent = y.parent
	if y.parent != nil {
		if y == y.parent.left {
			y.parent.left = x
		} else {
			y.parent.right = x
		}
	} else {
		t.root = x
	}
	if y != z {
		z.key = y.key
		z.value = y.value
	}
	if y.color == black {
		t.deleteFixup(x)
	}
	t.count--
}

func (t *TreeMap) findNode(id Key) *node {
	current := t.root
	for current != sentinel {
		if id == current.key {
			return current
		}
		if t.less(id, current.key) {
			current = current.left
		} else {
			current = current.right
		}
	}
	return sentinel
}

// ForwardIterator represents a position in a tree map.
// It starts at the one-before-the start position and goes to the end.
type ForwardIterator struct {
	tree *TreeMap
	node *node
	end  *node
}

// HasNext reports if we have elements after current position
func (i ForwardIterator) HasNext() bool { return i.node != i.end }

// Next returns next element from a tree map.
// It panics if goes out of bounds.
func (i *ForwardIterator) Next() (key Key, value Value) {
	if i.node == i.end {
		panic("out of bound iteration")
	}
	key, value = i.node.key, i.node.value
	if i.node.right != sentinel {
		i.node = i.node.right
		for i.node.left != sentinel {
			i.node = i.node.left
		}
		return
	}
	for i.node.parent != nil {
		i.node = i.node.parent
		if !i.tree.less(i.node.key, key) {
			return
		}
	}
	i.node = i.end
	return
}

// ReverseIterator represents a position in a tree map.
// It starts at the one-past-the-end position and goes to the beginning.
type ReverseIterator struct {
	tree *TreeMap
	node *node
	end  *node
}

// HasNext reports if we have elements after current position
func (i ReverseIterator) HasNext() bool { return i.node != i.end }

// Next returns next element from a tree map
func (i *ReverseIterator) Next() (key Key, value Value) {
	if i.node == i.end {
		panic("out of bound iteration")
	}
	key, value = i.node.key, i.node.value
	if i.node.left != i.end {
		i.node = i.node.left
		for i.node.right != i.end {
			i.node = i.node.right
		}
		return
	}
	for i.node.parent != nil {
		i.node = i.node.parent
		if !i.tree.less(key, i.node.key) {
			return
		}
	}
	i.node = i.end
	return
}
