package main

import (
	"bufio"
	"fmt"
	"strings"
)

// 练习 7.1：使用来自ByteCounter的思路，实现一个针对对单词和行数的计数器。

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

type WordsCounter struct {
	line int
	word int
}

func (w *WordsCounter) Write(p string) (int, int, error) {
	reader := strings.NewReader(p)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		w.line++
		lineReader := strings.NewReader(scanner.Text())
		wordScanner := bufio.NewScanner(lineReader)
		wordScanner.Split(bufio.ScanWords)
		for wordScanner.Scan() {
			w.word++
		}
	}
	return w.line, w.word, nil
}

func main() {
	// var c ByteCounter
	// c.Write([]byte("hello"))
	// fmt.Println(c) // "5", = len("hello")
	// c = 0          // reset the counter
	// var name = "Dolly"
	// fmt.Fprintf(&c, "hello, %s", name)
	// fmt.Println(c) // "12", = len("hello, Dolly")

	var w WordsCounter

	w.Write("hello world\ndfdf")
	fmt.Println(w)
}
