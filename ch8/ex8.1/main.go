package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

/*
练习 8.1：
修改clock2来支持传入参数作为端口号，然后写一个clockwall的程序，
这个程序可以同时与多个clock服务器通信，从多服务器中读取时间，
并且在一个表格中一次显示所有服务传回的结果，
*/

func handleConn(c net.Conn, timezone string) {
	defer c.Close()
	fmt.Println(timezone)
	for {
		loc, err := time.LoadLocation(timezone)
		if err != nil {
			log.Fatal(err)
			return
		}
		now := time.Now().In(loc)
		timeStr := now.Format("15:04:05\n")
		_, err = io.WriteString(c, timeStr)
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func CreateClock(server string, timezone string) {
	listener, err := net.Listen("tcp", server)
	if err != nil {
		log.Fatal(err)
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleConn(conn, timezone)
	}
}

func main() {
	if len(os.Args) < 2 {
		go CreateClock("localhost:8000", "")
	} else {
		fmt.Println(os.Args)
		for _, i := range os.Args[1:] {
			info := strings.Split(i, "=")
			if len(info) != 2 {
				log.Fatal(fmt.Errorf("格式错误: %s", i))
				return
			}
			timezone := info[0]
			server := info[1]
			go CreateClock(server, timezone)
		}
	}

	for {
		
	}
}
