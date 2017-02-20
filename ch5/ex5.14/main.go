package main

import (
	"fmt"
)
/*
 使用breadthFirst遍历其他数据结构。
 比如，topoSort例子中的课程依赖关系（有向图）,个人计算机的文件层次结构（树），你所在城市的公交或地铁线路（无向图）。
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

// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
var classMap = map[string][]string{}
func breadthFirst(f func(item []string) []string, worklist []string, k string) {
	seen := make(map[string]bool)

	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, p := range items {
			if !seen[p] {
				seen[p] = true
				classMap[p] = append(classMap[p], k)
			}
		}
	}
}

func main() {
	for k, v := range prereqs {
		breadthFirst(nil, v, k)
	}

	for k, v := range classMap {
		fmt.Println(k)
		for _, c := range v {
			fmt.Println("\t" + c)
		}
		fmt.Println("------")
	}

}
