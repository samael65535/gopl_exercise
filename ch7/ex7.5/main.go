package main

/*
 练习 7.5：
 io包里面的LimitReader函数接收一个io.Reader接口类型的r和字节数n，
 并且返回另一个从r中读取字节但是当读完n个字节后就表示读到文件结束的Reader。
*/
import (
	"io"
	"os"
	"strings"
)

type MyReader struct {
	reader io.Reader
	limit  int
	readed int
}

func (m *MyReader) Read(b []byte) (int, error) {
	n, e := m.reader.Read(b)
	m.readed += n
	if e != nil {
		return n, e
	}
	if m.readed >= m.limit {
		return m.limit, io.EOF
	}
	return n, nil
}

func LimitReader(r io.Reader, n int) io.Reader {
	return &MyReader{r, n, 0}
}

func main() {
	s := "123456789"
	r := LimitReader(strings.NewReader(s), 4)
	io.Copy(os.Stdout, r)
}
