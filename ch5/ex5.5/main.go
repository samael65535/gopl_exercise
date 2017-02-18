package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	//list, err := find(os.Args[1:])
	urlList := []string{"http://www.baidu.com"}
	list, err := find(urlList)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	for idx, res := range *list {
		doc, err := html.Parse(bytes.NewReader(res))
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
			os.Exit(1)
		}

		w, i := visit(0, 0, doc)
		fmt.Println(urlList[idx])
		fmt.Printf("world: %d\nimg: %d\n", w, i)
	}
}

func countWords(s string) int {
	count := 0
	in := bufio.NewScanner(strings.NewReader(s))
	in.Split(bufio.ScanWords)
	for in.Scan() {
		count++
	}
	return count

}

// visit appends to links each link found in n and returns the result.
func visit(wordsCount int, imgCount int, n *html.Node) (int, int) {
	if n == nil {
		return 0, 0
	}
	w, i := visit(wordsCount, imgCount, n.FirstChild)

	if n.Type == html.TextNode {
		w += countWords(n.Data)
	}

	if n.Type == html.ElementNode && n.Data == "img" {
		i++
	}

	w_, i_ := visit(wordsCount, imgCount, n.NextSibling)

	return wordsCount + w + w_, imgCount + i + i_
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
