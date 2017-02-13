package main

/*
练习 2.3： 重写PopCount函数，用一个循环代替单一的表达式。比较两个版本的性能。（11.4节将展示如何系统地比较两个不同实现的性能。）

练习 2.4： 用移位算法重写PopCount函数，每次测试最右边的1bit，然后统计总数。比较和查表算法的性能差异。

练习 2.5： 表达式x&(x-1)用于将x的最低的一个非零的bit位清零。使用这个算法重写PopCount函数，然后比较性能。
*/
import (
	"fmt"
	"time"

	"./popcount"
)

func main() {
	var num uint64
	num = 9099
	start := time.Now()
	fmt.Println(popcount.PopCount1(num))
	fmt.Printf("elapsed %f\n", time.Since(start).Seconds())

	start = time.Now()
	fmt.Println(popcount.PopCount2(num))
	fmt.Printf("elapsed %f\n", time.Since(start).Seconds())

	start = time.Now()
	fmt.Println(popcount.PopCount3(num))
	fmt.Printf("elapsed %f\n", time.Since(start).Seconds())

	start = time.Now()
	fmt.Println(popcount.PopCount4(num))
	fmt.Printf("elapsed %f\n", time.Since(start).Seconds())

}
