package main

import (
	"fmt"
)

var test = make(chan int)

func main() {
	go f1()
	go f2()
	test <- 1

	fmt.Println("finish")
}

func f1() {
	fmt.Println("in f1", <-test)
}

func f2() {
	fmt.Println("in f2", <-test)
}
