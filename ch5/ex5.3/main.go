package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var textDoc []string

func main() {
	list, err := find(os.Args[1:])
	//list, err := find([]string{"http://www.baidu.com"})
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
		visit(doc)
	}
}

func visit(n *html.Node) {
	if n == nil {
		return
	}

	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "sytle") {
		return
	}
	visit(n.FirstChild)
	if n.Type == html.TextNode {
		fmt.Println(n.Data)
	}
	visit(n.NextSibling)
}

func find(urlList []string) (*[][]byte, error) {
	list := [][]byte{{}}
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
