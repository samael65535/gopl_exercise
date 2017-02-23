package main

/*
练习 7.6：
对tempFlag加入支持开尔文温度。
*/

/*
练习 7.7：
解释为什么帮助信息在它的默认值是20.0没有包含°C的情况下输出了°C。


默认tempconv.CelsiusFlag, 输入的类型是Celsius
*/
import (
	"flag"
	"fmt"
	"time"

	"./tempconv"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")
var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println("ddd")
	flag.Parse()
	fmt.Println(*temp)
}
