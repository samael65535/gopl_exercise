package main

/*
练习5.12:
gopl.io/ch5/outline2（5.5节）的startElement和endElement共用了全局变量depth，
将它们修改为匿名函数，使其共享outline中的局部变量。
*/
import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// forEachNode针对每个结点x,都会调用pre(x)和post(x)。
// pre和post都是可选的。
// 遍历孩子结点之前,pre被调用
// 遍历孩子结点之后，post被调用
func forEachNode(n *html.Node, pre, post func(n *html.Node), depth int) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, func(n *html.Node) {
			if n.Type == html.ElementNode {
				fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
			}
		}, func(n *html.Node) {
			if n.Type == html.ElementNode {
				fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
			}
		}, depth+1)
	}

	if post != nil {
		post(n)
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

		forEachNode(doc, nil, nil, 0)
	}
}
