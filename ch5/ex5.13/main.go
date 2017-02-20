package main
/*
修改crawl，使其能保存发现的页面，必要时，可以创建目录来保存这些页面。
只保存来自原始域名下的页面。假设初始页面在golang.org下，就不要保存vimeo.com下的页面。
*/
import (
	"fmt"
	"log"
	"./links"
	"net/url"
	"net/http"
	"io/ioutil"

	"strings"
	"os"
)

// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string, hosts map[string]bool) []string, worklist []string) {
	seen := make(map[string]bool)
	hosts := make(map[string]bool)
	for _, k := range worklist {
		s,_ := url.Parse(k)
		hosts[s.Host] = true;
		_ = os.Mkdir(BasePath+ s.Host, 0755)
	}
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item, hosts)...)
			}
		}
	}
}

const BasePath = "./output/"
func crawl(u string, hosts map[string]bool) []string {
	currentURL, e := url.Parse(u)
	if !hosts[currentURL.Host] {
		return nil
	}
	list, err := links.Extract(u)

	if err != nil {
		log.Print(err)
		return list
	}

	if hosts[currentURL.Host] && e == nil && currentURL.Path != ""{
		resp, err := http.Get(u)
		if err != nil {
			resp.Body.Close()
			//			log.Print(err)
			return list
		}
		body, e := ioutil.ReadAll(resp.Body)
		if e != nil {
			resp.Body.Close()
			log.Print(err)
			return list
		}
		path := BasePath + currentURL.Host + currentURL.Path
		filename := path + ".html"
		if strings.HasSuffix(currentURL.Path, "/") {
			// os
			e = os.MkdirAll(path, 0755)
			if e != nil {
				resp.Body.Close()
				log.Print(err)
				return list
			}
			filename = path + "index.html"
		}
		e = ioutil.WriteFile(filename, body, 0755)
		if e != nil {
			log.Print(err)
		}
		fmt.Printf("Crawling...\t%s\t%s\t%s\n", currentURL.Path, filename, u)
		resp.Body.Close()
		return list
	}
	return list
}

func main() {
	_ = url.Parse
	breadthFirst(crawl, []string{"https://golang.org"})
}
