package main
/*
练习 11.1:
为4.3节中的charcount程序编写测试。
*/
import (
	"testing"
	"strings"
	"bufio"
	"fmt"
)

func TestCharcount(t *testing.T) {
	var tests = []struct {
		input string
		char  rune
		count int
	}{
		{"abcd", 'a', 1},
		{"⌘⌘⌘⌘⌘", '⌘', 5},
		{"été", 'é', 2},
	}

	for _, test := range tests {
		in := bufio.NewReader(strings.NewReader(test.input))
		counts, _, _, _ := CountChar(in)
		fmt.Println(counts)
		if counts[test.char] != test.count {
			t.Errorf("counts of %v = %v", test.char, counts[test.char])
		}
	}
}
