package main

/*
练习 7.18：
 使用基于标记的解码API，编写一个可以读取任意XML文档和构造这个文档所代表的普通节点树的程序。
节点有两种类型：CharData节点表示文本字符串，和 Element节点表示被命名的元素和它们的属性。每一个元素节点有一个字节点的切片。
*/

import (
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

type Node interface{} // CharData or *Element

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func visit(n Node, depth int) {
	switch n := n.(type) {
	case *Element:
		fmt.Printf("%*s", depth, n.Type.Local)
		for _, child := range n.Children {
			visit(child, depth+1)
		}
	case CharData:
		fmt.Printf("%*s", depth, n)
	default:

	}

}


func parse(r io.Reader) (Node, error) {
	dec := xml.NewDecoder(r)
	var stack []*Element
	var root Node
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			el := &Element{tok.Name, tok.Attr, nil}
			if len(stack) == 0 {
				root = el
			} else {
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, el)
			}
			stack = append(stack, el) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			parent := stack[len(stack)-1]
			parent.Children = append(parent.Children, CharData(tok))
		}
	}
	return root, nil
}

func main() {
	in := strings.NewReader("<A><B><C>hello</C><D>abc</D></B><C>world</C></A>")
	node,_ := parse(in)
	visit(node, 0)
}
