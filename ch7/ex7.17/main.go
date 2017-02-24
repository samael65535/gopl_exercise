// Xmlselect prints the text of selected elements of an XML document.
package main
/*
练习 7.17：
 扩展xmlselect程序以便让元素不仅仅可以通过名称选择，也可以通过它们CSS样式上属性进行选择；
*/
import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack [][]string // stack of element names
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			dic := []string{}
			dic = append(dic, tok.Name.Local)
			for _, val := range tok.Attr {
				if val.Name.Local == "id" || val.Name.Local == "class" {
					dic = append(dic, val.Value)
				}
			}
			stack = append(stack, dic) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				for _, v := range stack {
					fmt.Printf("%s ", strings.Join(v, "|"))
				}
				fmt.Printf(": %s\n", tok)
			}
		}
	}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x [][]string, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}

		for _, element := range x[0] {
			if element == y[0] {
				y = y[1:]
				break
			}
		}
		x = x[1:]
	}
	return false
}
