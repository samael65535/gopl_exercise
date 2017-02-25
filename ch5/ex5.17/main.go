package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

/*
练习5.17:
编写多参数版本的ElementsByTagName，函数接收一个HTML结点树以及任意数量的标签名
返回与这些标签名匹配的所有元素。
*/

func isExist(doc *html.Node, name ...string) bool {
	for _, c := range name {
		if c == (*doc).Data {
			return true
		}
	}
	return false
}
func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	if doc == nil {
		return nil
	}

	res := []*html.Node{}
	if doc.Type == html.ElementNode {
		//		fmt.Println(doc.Data)
		// res = append(res, ElementsByTagName(doc, name...)...)
		if isExist(doc, name...) {
			res = append(res, doc)
		}
	}
	for n := doc.FirstChild; n != nil; n = n.NextSibling {
		res = append(res, ElementsByTagName(n, name...)...)
	}

	return res
}

func find(url string, name ...string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return
	}

	doc, e := html.Parse(resp.Body)
	if e != nil {
		log.Fatal(e)
		return
	}
	res := ElementsByTagName(doc, name...)

	for _, n := range res {
		fmt.Print(n.Data)
		for _, a := range n.Attr {
			fmt.Printf("\t%s=%s", a.Key, a.Val)
		}
		fmt.Println("\n------")
	}
}

func main() {
	find("http://www.ituring.com.cn", "div", "iframe")
}
