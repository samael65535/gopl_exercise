package main

// 编写函数expand，将s中的"foo"替换为f("foo")的返回值。
import (
	"fmt"
	"strings"
)

func main() {
	str := "foo foo foo"
	str = expand(str, func(s string) string {
		return fmt.Sprintf("foo(\"%s\")", s)
	})
	fmt.Println(str)
}

func expand(s string, f func(string) string) string {
	return strings.Replace(s, "foo", f("foo"), -1)
}
