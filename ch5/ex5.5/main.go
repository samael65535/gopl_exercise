package main

import (
	"bufio"
	"strings"
)

func main() {

}

func wordCount(s string) int {
	count := 0
	in := bufio.NewScanner(strings.NewReader(s))
	in.Split(bufio.ScanWords)
	for in.Scan() {
		count++
	}
	return count

}
