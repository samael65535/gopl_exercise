package main

/*
练习 5.1：
修改findlinks代码中遍历n.FirstChild链表的部分，将循环调用visit，改成递归调用。
*/
import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	list, err := find(os.Args[1:])
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
		for _, link := range visit(nil, doc) {
			fmt.Println(link)
		}
	}

}

// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}

	links = visit(links, n.FirstChild)
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	links = visit(links, n.NextSibling)

	return links
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
