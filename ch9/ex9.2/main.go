package main
/*
练习 9.2：
重写2.6.2节中的PopCount的例子，使用sync.Once，只在第一次需要用到的时候进行初始化。
*/
import (
	"sync"
	"fmt"
)
var pc [256]byte
func loadPC() {
	fmt.Println("init")
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}
var loadPCOnce sync.Once
func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(PopCount(0x1234567890ABCDEF + uint64(i)))
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	loadPCOnce.Do(loadPC)
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

//!-
