// 在fatchall中尝试使用长一些的参数列表，比如使用在alexa.com的上百万网站里排名靠前的。如果一个网站没有回应，程序将采取怎样的行为？

// 如果一个网站长时间没有响应, 那么会把http err写入管道, 并跳过

package main

import (
	"fmt"
	"net/http"
	"time"
	"io"
	"io/ioutil"
)

var urlList = []string{
	"https://www.baidu.com",
	"https://www.tmall.com/",
	"http://www.qq.com/",
	"http://www.sohu.com/",
	"https://www.taobao.com/"}


func main() {
	start := time.Now()

	ch := make(chan string)
	for _, url := range urlList {
		go fetchAll(url, ch)
	}

	for range urlList {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetchAll(url string, ch chan<- string) {
	start := time.Now()

	// 在这可以设置超时时间
	c := &http.Client {
		Timeout: 10*time.Second,
	}
	resp, err := c.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close();
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
