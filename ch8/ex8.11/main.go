package main

/*
练习 8.11：
紧接着8.4.4中的mirroredQuery流程，实现一个并发请求url的fetch的变种。当第一个请求返回时，直接取消其它的请求。
*/

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	done := make(chan struct{})
	fmt.Println(mirroredQuery([]string{
		"http://www.baidu.com",
		"http://www.163.com",
		"http://www.sina.com",
		"http://www.qq.com",
		"http://www.360.com",
	}, &done))
}
func mirroredQuery(urls []string, done *chan struct{}) string {
	responses := make(chan string)
	for _, u := range urls {
		go func(url string, done *chan struct{}) {
			responses <- request(url, done)
			close(responses)
		}(u, done)
	}
	return <-responses
}

func request(url string, done *chan struct{}) (response string) {
	rest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return "failed"
	}
	rest.Cancel = *done
	resp, err := http.DefaultClient.Do(rest)

	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}
	if resp.StatusCode != http.StatusOK {
		return strconv.FormatInt(int64(resp.StatusCode), 10)
	}
	*done <- struct{}{}
	return url
}
