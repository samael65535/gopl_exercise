package main

import (
	"fmt"
	"./github"
	"bufio"
	"os"
	"golang.org/x/crypto/ssh/terminal"
)
// 只能在*nix系统运行

//  编写一个工具，允许用户在命令行创建、读取、更新和关闭GitHub上的issue，当必要的时候自动打开用户默认的编辑器用于输入文本信息。

func main() {
	if (github.USERNAME == "" || github.PASSWORD == "") {
		github.USERNAME, github.PASSWORD = getUserPassword()
	}


	// TODO: 打开默认编辑器并输出文本
	issue := github.Issue{
		Title: "Testing",
		Body: "Body",
	}
	err := github.CreateIssue(&issue)
	if err != nil {
		fmt.Println(err.Error())
	}
}
func getUserPassword() (string, string){
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Username: ")
	username, _ := reader.ReadString('\n')

	fmt.Print("Enter Password: ")
	bytePassword, err := terminal.ReadPassword(0)
	if err == nil {
		fmt.Println("\nPassword typed: " + string(bytePassword))
	}
	password := string(bytePassword)

	return username,password
}
