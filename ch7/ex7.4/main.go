package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		fmt.Println(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

type MyHTMLReader struct {
	str string
}

func (m *MyHTMLReader) Read(b []byte) (int, error) {
	n := copy(b, m.str)
	m.str = m.str[n:]
	if len(m.str) == 0 {
		return n, io.EOF
	}
	return n, nil
}

func NewReader(s string) io.Reader {
	return &MyHTMLReader{s}
}

func main() {
	doc, err := html.Parse(NewReader("<h1>Hello!</h1>"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}
