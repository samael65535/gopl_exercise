package main

import (
	"fmt"
)

func main() {
	arr := []string{
		"abcd",
		"aaaa",
		"aaaa",
		"abcd"}
	fmt.Println(clearSpace(arr))
}

func clearSpace(arr []string) []string {
	end := len(arr)
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[i] == arr[j] {

			}
		}
	}
	return arr[0:end]
}
