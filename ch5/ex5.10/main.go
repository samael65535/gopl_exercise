package main

import "fmt"

/*
练习5.10:
重写topoSort函数，用map代替切片并移除对key的排序代码。验证结果的正确性（结果不唯一）。
*/

// prereqs记录了每个课程的前置课程
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	for k := range m {
		seen[k] = false
	}
	var visitAll func(items map[string]bool)
	visitAll = func(items map[string]bool) {
		for k := range items {
			if items[k] == false {
				prev := m[k]
				flag := true
				for _, pk := range prev {
					if m[pk] == nil && !seen[pk] {
						seen[pk] = true
						order = append(order, pk)
					}
					flag = flag && seen[pk]

				}

				if flag {
					seen[k] = true
					order = append(order, k)
					visitAll(seen)
				}
			}
		}
	}
	visitAll(seen)
	return order
}
