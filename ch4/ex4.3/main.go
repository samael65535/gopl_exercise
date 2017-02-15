package main

//  重写reverse函数，使用数组指针代替slice。

import (
	"fmt"
)

func main() {
	arr := [4]int{1, 3, 3, 4}
	reverse(&arr)
	fmt.Println(arr)
}

func reverse(s *[4]int) {
	arr := s
	len := len(arr) - 1
	for i := 0; i < len/2; i++ {
		arr[i], arr[len-i] = arr[len-i], arr[i]
	}
}
