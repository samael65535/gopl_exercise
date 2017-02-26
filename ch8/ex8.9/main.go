package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sync"
	"time"
)
/*
练习 8.9：
编写一个du工具，每隔一段时间将root目录下的目录大小计算并显示出来。
*/

const rootPath = "../../"
var wg *sync.WaitGroup
var sema chan struct{}
func main() {
	wg = &sync.WaitGroup{}
	sema = make(chan struct{}, 20)
	var totalSize uint64 = 0
	var fileSize chan uint64
	ticker := time.NewTicker(5 * time.Second)
	fileSize = make(chan uint64)
	for {
		select {
		case <-ticker.C:
			go walkDir(rootPath, fileSize, wg)
			go func() {
				wg.Wait()
				fmt.Println(totalSize)
				totalSize = 0
			}()

		case size := <-fileSize:
			totalSize += size
		}
	}

}

func walkDir(dir string, fileSize chan<- uint64, wg *sync.WaitGroup) {
	sema<- struct{}{}
	wg.Add(1)
	defer func() {
		wg.Done()
		<-sema
	}()
	items, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, file := range items {
		if file.IsDir() {
			p := filepath.Join(dir, file.Name())
			go walkDir(p, fileSize, wg)
		} else {
			fileSize <- uint64(file.Size())
		}
	}
}
