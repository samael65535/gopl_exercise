package main

/*
练习 8.10：
HTTP请求可能会因http.Request结构体中Cancel channel的关闭而取消。修改8.6节中的web crawler来支持取消http请求
*/

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

type LinkInfo struct {
	link  []string
	depth int
}

var depth int
var tokens = make(chan struct{}, 20)
var cancel = make(chan struct{})

func main() {
	depth = 1
	worklist := make(chan LinkInfo)
	var n int
	// Start with the command-line arguments.
	n++
	go func() {
		info := LinkInfo{
			link:  []string{"http://www.baidu.com"},
			depth: 0,
		}
		worklist <- info
	}()

	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		fmt.Println("close")
		close(cancel)
	}()
	// Crawl the web concurrently.
	seen := make(map[string]bool)

	for ; n > 0; n-- {
		info := <-worklist
		for _, link := range info.link {
			if !seen[link] && depth >= info.depth {
				seen[link] = true
				n++
				go func(link string, depth int) {
					worklist <- crawl(link, depth)
				}(link, info.depth)
			}
		}
	}

}

func crawl(url string, curDepth int) LinkInfo {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := Extract(url)
	info := LinkInfo{
		link:  list,
		depth: curDepth + 1,
	}
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return info
}

// Extract makes an HTTP GET request to the specified URL, parses
// the response as HTML, and returns the links in the HTML document.
func Extract(url string) ([]string, error) {
	rest, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	rest.Cancel = cancel
	resp, err := http.DefaultClient.Do(rest)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

//!-Extract

// Copied from gopl.io/ch5/outline2.
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
