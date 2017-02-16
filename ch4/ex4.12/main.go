package main

/*
流行的web漫画服务xkcd也提供了JSON接口。例如，一个 https://xkcd.com/571/info.0.json 请求将返回一个很多人喜爱的571编号的详细描述。下载每个链接（只下载一次）然后创建一个离线索引。编写一个xkcd工具，使用这些离线索引，打印和命令行输入的检索词相匹配的漫画的URL。
*/

// NOTICE: 还没用到并发与管理, 以后会重构一版
// NOTICE: https://xkcd.com/404/info.0.json 其实不存在
import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

const XKCDURL = "https://xkcd.com/"

type Info struct {
	SafeTitle  string `json:"safe_title"`
	Title      string
	Year       string
	Month      string
	Day        string
	Num        int
	Img        string
	Transcript string
}

var items []Info

func main() {
	file, _ := os.OpenFile("xkcd.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	i := 405
	for {
		res, err := getJson(i)
		if i == 404 {
			i++
		}
		fmt.Println("Downloading.... " + XKCDURL + strconv.Itoa(i) + "/info.0.json")
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		fmt.Fprintf(file, "%d %s-%s-%s\t%s\n%s\n\n",
			(*res).Num, (*res).Year, (*res).Month, (*res).Day, (*res).Img, (*res).Transcript)
		fmt.Println("Finish.... " + XKCDURL + strconv.Itoa(i) + "/info.0.json")
		i++
	}
	defer file.Close()
}

func getJson(idx int) (*Info, error) {
	resp, err := http.Get(XKCDURL + strconv.Itoa(idx) + "/info.0.json")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed: %s", resp.Status)
	}
	var result Info
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	return &result, nil
}
