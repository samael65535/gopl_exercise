package main

/*
练习 2.2：
写一个通用的单位转换程序，用类似cf程序的方式从命令行读取参数，如果缺省的话则是从标准输入读取参数，
然后做类似Celsius和Fahrenheit的单位转换，长度单位可以对应英尺和米，重量单位可以对应磅和公斤等
*/
import (
	"fmt"
	"os"
	"strconv"

	"./lenconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}

		ft := lenconv.Foot(t)
		m := lenconv.Meter(t)

		fmt.Printf("%s = %s, %s = %s\n", ft, lenconv.FTToM(ft).String(), m, lenconv.MToFT(m).String())
	}
}
