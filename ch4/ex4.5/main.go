package main
// 写一个函数在原地完成消除[]string中相邻重复的字符串的操作。
import (
	"fmt"
)

func main() {
	arr := []string{
		"aaaa",
		"a",
		"a",
		"aaaa",
		"a",
	}

	fmt.Println(clearDuplicate(arr))

	arr1 := []string{
		"a",
		"a",
	}
	fmt.Println(clearDuplicate(arr1))


	arr2 := []string{
		"a",
		"aaa",
		"a",
	}
	fmt.Println(clearDuplicate(arr2))

	arr3 := []string{
		"a",
		"a",
	}
	fmt.Println(clearDuplicate(arr3))


	arr4 := []string{
		"aaa",
		"a",
		"a",
	}
	fmt.Println(clearDuplicate(arr4))

}

func clearDuplicate(arr []string) []string {
	end := len(arr)
	for i := 0; i < len(arr) - 1; i++ {
		if arr[i] == arr[i+1] {
			for j := i; j < len(arr)- 1; j++ {
				arr[j] = arr[j+1]
				end--
			}
		}
	}
	return arr[:end]
}
