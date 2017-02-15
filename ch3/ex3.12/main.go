package main

import (
	"strings"
	"fmt"
)
// 编写一个函数，判断两个字符串是否是是相互打乱的，也就是说它们有着相同的字符，但是对应不同的顺序。

// NOTE: 暂没学到sort所以只用strings库里的函数
func main() {
	fmt.Println(isEqual("dfdasfds", "dfdfasda"))
	fmt.Println(isEqual("1223", "3213"))
}

func isEqual(s1 string, s2 string) bool{

	if (len(s1) != len(s2))  {
		return false
	}

	for i:=0; i<len(s1);i++{
		c1 := s1[i]
		idx := strings.IndexByte(s2, c1)
		if idx == -1 {
			return false
		}
		s2 = s2[:idx] + s2[idx+1:]
	}
	return true
}
