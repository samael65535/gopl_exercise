// 打印出HTTP协议的状态码，可以从resp.Status变量得到该状态码。

package main

import (
	"fmt"
	"net/http"
	"io"
	"os"
	"strings"
)
func main() {
	for _, url :=  range os.Args[1:] {
		if strings.HasPrefix(url, "http://") == false {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %v\n", err)
			os.Exit(1)
		}

		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Fprintf(os.Stderr, "fetch reading %s: %v\n", url, err)
			resp.Body.Close()
			os.Exit(1)
		}
		resp.Body.Close()

		fmt.Println(resp.Status)
	}
}
