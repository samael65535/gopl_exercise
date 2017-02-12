// 修改dup2，出现重复的行时打印文件名称。

package main

import "os"
import "fmt"
import "bufio"

func countLines(f *os.File, counts map[string]int, filename map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		// 如果有集合型的数据结构可能更好
		if filename[input.Text()] != nil {
			filename[input.Text()][f.Name()]++
		} else {
			filename[input.Text()] = make(map[string]int)
			filename[input.Text()][f.Name()] = 1
		}

	}
}

func main() {
	counts := make(map[string]int)
	dupLineFileName := make(map[string]map[string]int)
	files := os.Args[1:]

	fmt.Println(files)
	if len(files) == 0 {
		countLines(os.Stdin, counts, dupLineFileName)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, dupLineFileName)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%v\n", n, line, dupLineFileName[line])
		}
	}
}
