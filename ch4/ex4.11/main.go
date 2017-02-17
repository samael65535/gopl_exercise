package main

import (
	"fmt"
	"./github"
	"bufio"
	"strconv"
	"os"
	"golang.org/x/crypto/ssh/terminal"
	"os/exec"
	"log"
	"io/ioutil"
)
// 只能在*nix系统运行

//  编写一个工具，允许用户在命令行创建、读取、更新和关闭GitHub上的issue，当必要的时候自动打开用户默认的编辑器用于输入文本信息。

func main() {
	// if (github.USERNAME == "" || github.PASSWORD == "") {
	//	github.USERNAME, github.PASSWORD = getUserPassword()
	// }


	switch(os.Args[1]) {
	case "create":
		title := os.Args[2]
		body, err := getIssueBodyByEditor()
		if err != nil {
			log.Fatal(err)
		}
		issue := github.Issue{
			Title: title,
			Body: string(*body),
			State: "open",
			Labels: []string{"ex4.11"},
		}
		err = github.CreateIssue(&issue)
		if err != nil {
			fmt.Println(err.Error())
		}
	case "close":
		err := github.CloseIssue(os.Args[2])
		if err != nil {
			fmt.Println(err.Error())

		}
	case "edit":
		strconv.ParseBool("false");
	case "load":
		github.LoadIssue(1)
	}
}


func getIssueBodyByEditor() (*[]byte, error){

	tempfile,err := ioutil.TempFile("/tmp", ".temp_issue")
	defer tempfile.Close()
	defer os.Remove(tempfile.Name())

	if err != nil {
		log.Fatal(err)
	}


	editorPath, err := exec.LookPath("vim")
	if err != nil {
		log.Fatal(err)
	}
	cmd := &exec.Cmd{
		Path:   editorPath,
		Args:   []string{"vim", tempfile.Name()},
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	err = cmd.Run()

	b, err := ioutil.ReadFile(tempfile.Name())
	return &b, err
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
