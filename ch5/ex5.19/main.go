package main

import (
	"fmt"
)

/*
练习5.19;
使用panic和recover编写一个不包含return语句但能返回一个非零值的函数。
*/

func foo(num int) (ret int) {
	defer func() {
		recover()
		ret = 618 + num
	}()
	panic("test")
}

func main() {
	fmt.Println(foo(306))
}
