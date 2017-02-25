package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"strings"

	"golang.org/x/net/html"
)

/*
练习5.7:
完善startElement和endElement函数，使其成为通用的HTML输出器。要求：输出注释结点，文本结点以及每个元素的属性（< a href='...'>）。
使用简略格式输出没有孩子结点的元素（即用<img/>代替<img></img>）。编写测试，验证程序输出的格式正确。
*/

var depth int

func main() {
	//list, err := find(os.Args[1:])
	list, err := find([]string{"http://www.baidu.com"})
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	for _, res := range *list {
		doc, err := html.Parse(bytes.NewReader(res))
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
			os.Exit(1)
		}

		forEachNode(doc, startElement, endElement)
	}
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func startElement(n *html.Node) {
	depth++
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s", depth*2, "", n.Data)
		if n.FirstChild != nil {
			for _, a := range n.Attr {
				fmt.Printf(" %s='%s'", a.Key, a.Val)
			}
			fmt.Printf(">\n")
		}
	}

	if n.Type == html.TextNode {
		fmt.Printf("%*s", depth*2, strings.TrimSpace(n.Data))
	}
}

func endElement(n *html.Node) {
	depth--
	if n.Type == html.ElementNode {
		if n.FirstChild == nil {
			fmt.Printf("/>\n")
		} else {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}

func find(urlList []string) (*[][]byte, error) {
	list := [][]byte{}
	for _, url := range urlList {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			return nil, err
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return nil, err
		}
		list = append(list, b)
	}
	return &list, nil
}
