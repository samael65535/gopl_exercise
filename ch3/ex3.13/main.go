package main

import (
	"fmt"

)
const (
	KB = 8 * 1024
	MB = 8 * 1024 << (10 * iota)
	TB
	PB
	EB
	ZB  //  uint64
	YB // overflows uint64
)
func main() {
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(TB)
	fmt.Println(PB)
	fmt.Println(EB)
	fmt.Println(uint64(ZB))
	//fmt.Println(YB)
}
