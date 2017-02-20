package main

import (
	"fmt"
)
// 使用panic和recover编写一个不包含return语句但能返回一个非零值的函数。

func foo() {
	panic("testing")
}

func main() {
	defer func() {
		if p:=recover(); p != nil {
			fmt.Println(p)
		}
	}()
	foo()

}
