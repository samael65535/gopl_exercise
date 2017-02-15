package main
// 完善comma函数，以支持浮点数处理和一个可选的正负号的处理。

import (
	"bytes"
	"fmt"
	"strings"
	"strconv"
)

func main() {

	fmt.Printf("%s\n", comma("-"))
	fmt.Printf("%s\n", comma("0"))
	fmt.Printf("%s\n", comma("-0"))
	fmt.Printf("%s\n", comma("+0"))
	fmt.Printf("%s\n", comma("12321312321"))
	fmt.Printf("%s\n", comma("12.020"))
	fmt.Printf("%s\n", comma("12..020"))
	fmt.Printf("%s\n", comma("123459345.020"))
	fmt.Printf("%s\n", comma("-123459345.020"))
	fmt.Printf("%s\n", comma("12345945.020"))
	fmt.Printf("%s\n", comma("-12345345.020"))

}

func comma (s string) string{
	var buf bytes.Buffer
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err.Error()
	}

	s = strconv.FormatFloat(v, 'f', -1, 64)
	len := len(s)

	// 处理小数点
	end := strings.Index(s, ".")
	if end == -1 {
		end = len
	}
	// 处理符号
	start := 0
	if s[0] == '-' || s[0] == '+' {
		start = 1
	}

	// 数值部分
	len = end - start
	if (len <= 3) {	return s }

	if (len%3==0) {
		start += 3
	} else {
		start += len%3
	}

	buf.WriteString(s[:start]);
	for i:= start; i < end; i+=3 {
		end_ := i+3
		if (end_ >= end) {
			end_ = end
		}
		buf.WriteString(","+s[i:end_])
	}
	buf.WriteString(s[end:])
	return buf.String()
 }
