package main

/*
练习 7.6：
对tempFlag加入支持开尔文温度。
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
