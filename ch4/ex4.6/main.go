package main
//编写一个函数，原地将一个UTF-8编码的[]byte类型的slice中相邻的空格（参考unicode.IsSpace）替换成一个空格返回
import (
	"fmt"
	"unicode/utf8"
	"unicode"
)

func main() {
	s := []byte("我有    很   多  的   空  格      ")
	s1 :=[]byte("我没有空格")
	s2 := []byte("    ")
	s3 := []byte("    我只有前面有空格")
	s4 := []byte("我只有后面有空格    ")

	clearSpace(s)
	clearSpace(s1)
	clearSpace(s2)
	clearSpace(s3)
	clearSpace(s4)
}


func clearSpace(b []byte) {
	count := 0
	for i := 0; i < len(b); {
		r, size := utf8.DecodeRune(b[i:])
		if (unicode.IsSpace(r)) {
			start_ := i + size
			end_ := start_
			for j := start_; j < start_ + len(b[start_:]); {
				c, cSize := utf8.DecodeRune(b[j:])
				if unicode.IsSpace(c) {
					end_ += cSize
				} else {
					break
				}
				j += cSize
			}
			len_ := end_- start_
			// 多余的空格数
			if len_ > count {
				count = len_
			}
			// 优化copy代码
			copy(b[start_:], b[end_:])
			// for j := start_; j < len(b); j++ {
			//	if (j+len_ < len(b)) {
			//		b[j] = b[j+len_]
			//	}
			// }
		}
		i += size
	}
	fmt.Println(string(b[:len(b) - count]), utf8.RuneCount(b[:len(b) - count]))
}
