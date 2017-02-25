package main

/*
练习 8.6：
为并发爬虫增加深度限制。也就是说，如果用户设置了depth=3，那么只有从首页跳转三次以内能够跳到的页面才能被抓取到。
*/

import (
	"fmt"
	"log"

	"../links"
)

type LinkInfo struct {
	link  []string
	depth int
}

var depth int
var tokens = make(chan struct{}, 20)

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
	list, err := links.Extract(url)
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
