package main
// 编写一个非递归版本的comma函数，使用bytes.Buffer代替字符串链接操作。

import (
	"bytes"
	"fmt"
)

func main() {
	num := "523"
	fmt.Printf("%s\n", comma(num))
}

func comma (s string) string{
	var buf bytes.Buffer
	len := len(s)
	if (len <= 3) {	return s }

	start := len%3
	if (len%3==0) {
		start = 3
	}

	buf.WriteString(s[:start]);
	for i:= start; i < len; i+=3 {
		end := i+3
		if (end > len) {
			end = i
		}
		buf.WriteString(","+s[i:end])
	}

	return buf.String()
}
