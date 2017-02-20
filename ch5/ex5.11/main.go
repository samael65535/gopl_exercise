package main

import "fmt"

// 重写topoSort函数，用map代替切片并移除对key的排序代码。验证结果的正确性（结果不唯一）。

// prereqs记录了每个课程的前置课程
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	//"linear algebra": {"calculus"},
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
	"computer organization": {"networks"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

var subsequent = map[string][]string{}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
	for k, v := range subsequent {
		fmt.Println(k, v)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)

	var visitParent func(root string) error

	visitParent = func(root, parent, current string) error {
		if root == current {
			return fmt.Errorf("err")
		}
		subsequent[root] = append(subsequent[root], current)
		for _, c := range subsequent[parent] {
			visitParent(root, current, subsequent[current])
		}
		return nil
	}

	for k, v := range prereqs {
		for _, c := range v {
			visitParent(c, k, c)
		}
	}

	var visitAll func(items map[string]bool)
	visitAll = func(items map[string]bool) {
		for k, v := range items {
			if v == false {
				flag := true
				for _, pk := range m[k] {
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
	if len(order) != len(subsequent) {
		return nil
	}
	return order
}
