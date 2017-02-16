package main
//  编写一个程序wordfreq程序，报告输入文本中每个单词出现的频率。在第一次调用Scan前先调用input.Split(bufio.ScanWords)函数，这样可以按单词而不是按行输入。

import (
	"bufio"
	"fmt"
	"os"
)


func main() {
	counts := make(map[string]int)
	in := bufio.NewScanner(bufio.NewReader(os.Stdin))
	in.Split(bufio.ScanWords)
	for in.Scan(){
		w := in.Text()
		counts[w]++
	}
	for w, n := range counts {
		fmt.Printf("%q\t%d\n", w, n)
	}

}
