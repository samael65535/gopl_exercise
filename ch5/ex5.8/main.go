package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

/*
练习5.8:
改pre和post函数，使其返回布尔类型的返回值。返回false时，中止forEachNode的遍历。
使用修改后的代码编写ElementByID函数，根据用户输入的id查找第一个拥有该id元素的HTML元素，查找成功后，停止遍历。
*/
func pre(id string, n *html.Node) bool {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				fmt.Println(a.Key, a.Val, n.Data)
				return true
			}
		}
		return false
	}
	return false
}

func ElementByID(id string, n *html.Node, pre, post func(id string, n *html.Node) bool) {
	if n != nil && pre(id, n) {
		return
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ElementByID(id, c, pre, post)
	}
	if n != nil && post != nil {
		post(id, n)
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
	list, err := find([]string{"https://www.github.com/"})
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
		ElementByID("ajax-error-message", doc, pre, nil)
	}
}
