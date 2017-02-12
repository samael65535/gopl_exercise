// 修改echo程序，使其打印每个参数的索引和值，每个一行。
package main

import "os"
import "fmt"

func main() {
	for idx, val := range os.Args {
		fmt.Println(idx, val)
	}
}
