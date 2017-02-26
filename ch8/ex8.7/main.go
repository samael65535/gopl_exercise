package main

/*
练习 8.7：
完成一个并发程序来创建一个线上网站的本地镜像，把该站点的所有可达的页面都抓取到本地硬盘。
为了省事，我们这里可以只取出现在该域下的所有页面
当然了，出现在页面里的链接你也需要进行一些处理，使其能够在你的镜像站点上进行跳转，而不是指向原始的链接。
*/

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"

	"golang.org/x/net/html"

	"../links"

	"bytes"
	"os"
	"strings"
)

// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string, hosts map[string]bool, depth int) []string, worklist [][]string) {
	seen := make(map[string]bool)
	hosts := make(map[string]bool)
	for _, k := range worklist[0] {
		s, _ := url.Parse(k)
		hosts[s.Host] = true
		_ = os.Mkdir(BasePath+s.Host, 0755)
	}

	for len(worklist) > 0 {
		for depth, items := range worklist {
			worklist = append(worklist, []string{})
			for _, item := range items {
				if !seen[item] {
					seen[item] = true
					list := f(item, hosts, depth)
					worklist[depth+1] = append(worklist[depth+1], list...)

				}
			}
		}
	}
}

const BasePath = "./output/"

func crawl(u string, hosts map[string]bool, depth int) []string {
	currentURL, e := url.Parse(u)
	if !hosts[currentURL.Host] {
		return nil
	}
	list, err := links.Extract(u)

	if err != nil {
		log.Print(err)
		return list
	}

	if hosts[currentURL.Host] && e == nil && currentURL.Path != "" {
		resp, err := http.Get(u)
		if err != nil {
			resp.Body.Close()
			return list
		}
		body, e := ioutil.ReadAll(resp.Body)
		if e != nil {
			resp.Body.Close()
			log.Print(err)
			return list
		}

		node, _ := html.Parse(bytes.NewReader(body))
		forEachNode(node, replaceLink, nil, depth)
		e = writeFile(node, currentURL, resp)

		if e != nil {
			resp.Body.Close()
			log.Print(err)
			return list
		}
		//fmt.Printf("Crawling...\t%s\t%s\t%s\n", currentURL.Path, filename, u)
		resp.Body.Close()
		return list
	}
	return list
}

func writeFile(node *html.Node, currentURL *url.URL, resp *http.Response) error {
	filename := BasePath + currentURL.Host + currentURL.Path
	e := os.MkdirAll(path.Dir(filename), 0755)
	if filename == "./"+path.Dir(filename)+"/" {
		filename = filename + "/index.html"
	}
	// os
	if e != nil {
		resp.Body.Close()
		log.Print(e)
	}

	b := &bytes.Buffer{}
	e = html.Render(b, node)
	if e != nil {
		return e
	}
	e = ioutil.WriteFile(filename, b.Bytes(), 0775)
	if e != nil {
		return e
	}
	return nil
}

// forEachNode针对每个结点x,都会调用pre(x)和post(x)。
// pre和post都是可选的。
// 遍历孩子结点之前,pre被调用
// 遍历孩子结点之后，post被调用
func replaceLink(n *html.Node, depth int) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for i, a := range n.Attr {
			if a.Key == "href" {
				strings.Repeat("../")
				n.Attr[i].Val = n.Attr[i].Val
			}
		}
	}
}
func forEachNode(n *html.Node, pre, post func(n *html.Node, d int), depth int) {
	if pre != nil {
		pre(n, depth)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post, depth)
	}

	if post != nil {
		post(n, depth)
	}
}

func main() {
	_ = url.Parse
	worklist := [][]string{[]string{"http://www.baidu.com"}}
	breadthFirst(crawl, worklist)
}
