package main

// 编写函数，记录在HTML树中出现的同名元素的次数。
import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

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
		counter := make(map[string]int)
		for k, v := range visit(doc, counter) {
			fmt.Println(k, v)
		}
	}

}

func visit(n *html.Node, counter map[string]int) map[string]int {
	if n == nil {
		return counter
	}

	if n.Type == html.ElementNode {
		counter[n.Data]++
	}
	visit(n.FirstChild, counter)
	visit(n.NextSibling, counter)
	return counter
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
