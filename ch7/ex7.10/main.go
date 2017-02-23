package main

import (
	"fmt"
	"sort"
)

/*
练习 7.10：
编写一个IsPalindrome(s sort.Interface) bool函数表明序列s是否是回文序列，换句话说反向排序不会改变这个序列。假设如果!s.Less(i, j) && !s.Less(j, i)则索引i和j上的元素相等。
*/

type MyString string

func (x MyString) Less(i, j int) bool { return x[i] < x[j] }
func (x MyString) Len() int { return len(x) }
func (x MyString) Swap(i, j int) {}

func IsPalindrome(s sort.Interface) bool {
	length := s.Len()
	for i, j := 0, length-1; i < length; {
		if !s.Less(i, j) && !s.Less(j, i) == false {
			return false
		}
		i++
		j--
	}
	return true
}
func main() {
	fmt.Println(IsPalindrome(MyString("1111132323")))
	fmt.Println(IsPalindrome(MyString("1112111")))
}
