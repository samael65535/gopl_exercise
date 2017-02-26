package main

/*
练习 8.8：
使用select来改造8.3节中的echo服务器，为其增加超时，这样服务器可以在客户端10秒中没有任何喊话时自动断开连接。
*/

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

//!+
func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	defer c.Close()

	typing := make(chan int)
	go func() {
		for {
			t := time.NewTicker(10 * time.Second)
			select {
			case <-t.C:
				fmt.Fprintln(c, "timeout!")
				close(typing)
				c.Close()
				return
			case <-typing:
				t.Stop()
			}
		}
	}()
	for input.Scan() {
		typing <- 1
		go echo(c, input.Text(), 1*time.Second)
	}

}

//!-

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
