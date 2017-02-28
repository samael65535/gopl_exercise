package main

import (
	"fmt"
)

/*
练习 9.4:
 创建一个流水线程序，支持用channel连接任意数量的goroutine，在跑爆内存之前，可以创建多少流水线阶段？一个变量通过整个流水线需要用多久？
*/

func main() {
	out := make(chan int)
	head := out
	var in chan int
	maxGoroutine := 100
	for i := 0 ; i< maxGoroutine; i ++ {
		in = out
		out = make(chan int)
		go handle(in, out)
	}
	fmt.Println("finish")
	head <- 1
	<-out
	close(head)
}

func handle(in chan int, out chan int) {
	num:=<-in
	out <- num+1
	fmt.Println(num)
	close(out)
}
