package main

//  向tempconv包添加类型、常量和函数用来处理Kelvin绝对温度的转换，Kelvin 绝对零度是−273.15°C，Kelvin绝对温度1K和摄氏度1°C的单位间隔是一样的。

import "fmt"
import "./tempconv"

func main() {

	fmt.Printf("Brrr! %v\n", tempconv.KToC(298.15).String())
	fmt.Printf("Brrr! %v\n", tempconv.CToK(25).String())
}
