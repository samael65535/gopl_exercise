// Issues prints a table of GitHub issues matching the search terms.
// 修改issues程序，根据问题的时间进行分类，比如不到一个月的、不到一年的、超过一年。


package main

import (
	"fmt"
	"log"
	"./github"
	"os"
	"time"
)


func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	items := make(map[string][]*github.Issue)
	now := time.Now()
	for _, item := range result.Items {
		create_at := item.CreatedAt
		// 本月内的
		if create_at.Month() == now.Month() {
			items["this month"] = append(items["this month"], item)
		} else if create_at.Year() == now.Year() {
			// 本年内的
			items["this year"] = append(items["this year"], item)
		} else {
			// 不是本年的
			items["over year"] = append(items["over year"], item)
		}

		// fmt.Printf("#%-5d %9.9s %.55s\n",
		//	item.Number, item.User.Login, item.Title)
	}
	for k, results := range items {
		fmt.Println(k)
		for _, item := range results {
			fmt.Printf("#%-5d %9.9s %.55s\n",
						item.Number, item.User.Login, item.Title)
		}
	}
}
