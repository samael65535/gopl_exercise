package main

/*
练习 7.3：
为在gopl.io/ch4/treesort (§4.4)的*tree类型实现一个String方法去展示tree类型的值序列。
*/

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
)

// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 101.

// Package treesort provides insertion sort using an unbalanced binary tree.
//!+
type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func (t *tree) String() string {

	ret := ""
	if t == nil {
		return ret
	}

	var queue []*tree

	queue = append(queue, t)
	ret += strconv.FormatInt(int64(t.value), 10)
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur.right == nil && cur.left == nil {
			continue
		}

		if cur.left != nil {
			queue = append(queue, cur.left)
			ret += " " + strconv.FormatInt(int64(cur.left.value), 10)
		} else {
			ret += " x"
		}

		if cur.right != nil {
			queue = append(queue, cur.right)
			ret += " " + strconv.FormatInt(int64(cur.right.value), 10)
		} else {
			ret += " x"
		}
	}

	return ""
}

//!-

func main() {
	data := make([]int, 10)
	var	root *tree
	for i := range data {
		data[i] = rand.Int() % 50
		root = add(root, data[i])
	}

//	appendValues(data, &root)
	Sort(data)
	fmt.Println(root)
	if !sort.IntsAreSorted(data) {
		fmt.Printf("not sorted: %v", data)
	}
	fmt.Println(data)
}
