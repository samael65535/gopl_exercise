//找一个数据量比较大的网站，用本小节中的程序调研网站的缓存策略，对每个URL执行两遍请求
//查看两次时间是否有较大的差别，并且每次获取到的响应内容是否一致，修改本节中的程序，将响应结果输出，以便于进行对比。

// 观察结果: 第二次时间大大小于第一次时间

/*
2.93s    10791  https://www.google.com
2.94s    10745  https://www.google.com
3.12s      227  https://www.baidu.com
3.29s      227  https://www.baidu.com
3.29s elapsed
*/

package main

import (
	"net/http"
	"fmt"
	"io"
	"os"
	"io/ioutil"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	fmt.Println("start")
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
		fmt.Println(<-ch) // receive from channel ch
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
