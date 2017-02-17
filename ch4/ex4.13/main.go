package main

import (
	"net/http"
	"strings"
	"fmt"
	"os"
	"net/url"
	"encoding/json"
	"io"
)
const OMDBAPI = "http://www.omdbapi.com/?t="

type Item struct {
	Title string
	Poster string
}
func main() {
	// 去空格
	moviename := strings.Join(os.Args[1:], " ")
	poster, _ := getMoviePoster(moviename)
	if *poster == (Item{}) {
		fmt.Println("The movie is not found!")
		return
	}
	if err := downloadPoster(poster); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Finish Downlading Poster...")
}

func getMoviePoster(moviename string) (*Item, error) {
	resp, err := http.Get(OMDBAPI + url.QueryEscape(moviename))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result Item
	if err:=json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func downloadPoster(item *Item) (error) {
	f,err := os.Create((*item).Title + ".jpg")
	if err != nil {
		return err
	}

	resp, err := http.Get((*item).Poster)
	io.Copy(f, resp.Body)
	defer f.Close()

	return nil
}
