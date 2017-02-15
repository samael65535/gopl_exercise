package main

// 编写一个rotate函数，通过一次循环完成旋转。
import "fmt"

func main() {
	fmt.Println(rotate([]int{1, 2, 3, 4, 5, 6}, 1))
}

func rotate(slice []int, start int) []int {
	for i := 0; i < start; i++ {
		slice = append(slice, slice[i])
	}
	return slice[start:]
}
