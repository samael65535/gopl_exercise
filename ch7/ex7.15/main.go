package main

/*
练习 7.15：
编写一个从标准输入中读取一个单一表达式的程序，用户及时地提供对于任意变量的值，然后在结果环境变量中计算表达式的值。优雅的处理所有遇到的错误。
*/

import (
	"../ex7.14/eval"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	fmt.Println("表达式:")
	in.Scan()
	expr := in.Text()
	fmt.Println("变量:x=3 y=3")
	in.Scan()
	var e = eval.Env{}

	vals := strings.Fields(in.Text())
	for _, s := range vals {
		v := strings.Split(s, "=")
		if len(v) != 2 {
			fmt.Fprintf(os.Stderr, "变量分配错误, 格式为 x=3 y=3")
			os.Exit(2)
		}
		name := v[0]
		var num float64
		var err error
		if num, err = strconv.ParseFloat(v[1], 64); err != nil {
			fmt.Fprintf(os.Stderr, "数值错误, %s", v[1])
			os.Exit(2)
		}
		e[eval.Var(name)] = num
	}

	result, err := eval.Parse(expr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "错误 %s", err.Error())
	}

	fmt.Println(result.Eval(e))
}
