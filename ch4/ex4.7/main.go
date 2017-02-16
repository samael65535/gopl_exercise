package main
// 修改reverse函数用于原地反转UTF-8编码的[]byte。是否可以不用分配额外的内存？

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := []byte("1234")
	s1 :=[]byte("123")
	s2 := []byte("1")
	s3 := []byte("")
	s4 := []byte("123456")

	s5 := []byte("你好12345")
	reverse(s)
	reverse(s1)
	reverse(s2)
	reverse(s3)
	reverse(s4)
	reverse(s5)


	fmt.Println(string(s))
	fmt.Println(string(s1))
	fmt.Println(string(s2))
	fmt.Println(string(s3))
	fmt.Println(string(s4))
	fmt.Println(string(s5))
}

func reverse(b []byte) {
	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:])
		mid := size/2 + i
		end := i + size - 1

		for j:=i; j < mid; j++ {
			b[j], b[end -(j - i)] = b[end - (j - i)], b[j]
		}
		i += size
	}
	end := len(b) - 1
	for j:=0; j< end/2; j++ {
		b[j], b[end-j] = b[end - j], b[j]
	}
}
