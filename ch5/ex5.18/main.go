package main

import (
	"path"
	"net/http"
	"os"
	"io"
	"fmt"
)
// 不修改fetch的行为，重写fetch函数，要求使用defer机制关闭文件。

// Fetch downloads the URL and returns the
// name and length of the local file.
func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	local := path.Base(resp.Request.URL.Path)
	fmt.Println(local)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)

	defer func() {
		// Close file, but prefer error from Copy, if any.
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()
	return local, n, err
}


func main() {
	fetch("http://www.ituring.com.cn/index.html")
}
