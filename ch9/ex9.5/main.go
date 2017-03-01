package main

import (
	"fmt"
	"time"
)

/*
练习 9.5:
写一个有两个goroutine的程序，两个goroutine会向两个无buffer channel反复地发送ping-pong消息。这样的程序每秒可以支持多少次通信？
*/
func ping(in, out chan string) {
	for {
		//select {
		s := <-in
		fmt.Println(s)
		out <- "pong"
//		default:
//			fmt.Println("ping")
//		}
	}
}

func pong(out, in chan string) {
	for {
		s:= <-out
		fmt.Println(s)
		in <- "ping"
	}
}

func main() {
	in := make(chan string)
	out := make(chan string)
	done := make(chan struct{})
	go ping(in, out)
	go pong(out, in)
	time.Sleep(2)
	in <- "ping"
	fmt.Println("pong")
	_ =fmt.Print
	<-done
}
