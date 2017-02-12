// 测量潜在低效的版本和使用了strings.Join的版本的运行时间差异。
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var res, sep string
	args := os.Args
	var start = time.Now()
	for _, val := range args {
		res += sep + val
		sep = " "
	}
	fmt.Println(res)
	fmt.Printf("%d elapsed\n", time.Since(start).Nanoseconds())

	start = time.Now()
	for _, val := range args {
		fmt.Print(sep + val)
		sep = " "
	}
	fmt.Printf("\n%d elapsed\n", time.Since(start).Nanoseconds())

	start = time.Now()
	fmt.Println(strings.Join(args, " "))
	fmt.Printf("%d elapsed\n", time.Since(start).Nanoseconds())
}
