package main

/*
练习5.15:
编写类似sum的可变参数函数max和min。
考虑不传参时，max和min该如何处理，再编写至少接收1个参数的版本。
*/
import (
	"fmt"
)

func min(vals ...int) int {
	if len(vals) == 0 {
		panic()
	}
	return 0
}

func max(vals ...int) int {
	if len(vals) == 0 {
		panic()
	}
	return 0
}

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total

}
func main() {
	vals := []int{}
	fmt.Println(sum(vals...))
}
