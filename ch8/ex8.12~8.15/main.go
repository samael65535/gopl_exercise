package main

/*
练习 8.12： 使broadcaster能够将arrival事件通知当前所有的客户端。为了达成这个目的，你需要有一个客户端的集合，并且在entering和leaving的channel中记录客户端的名字。

练习 8.13： 使聊天服务器能够断开空闲的客户端连接，比如最近五分钟之后没有发送任何消息的那些客户端。提示：可以在其它goroutine中调用conn.Close()来解除Read调用，就像input.Scanner()所做的那样。

练习 8.14： 修改聊天服务器的网络协议这样每一个客户端就可以在entering时可以提供它们的名字。将消息前缀由之前的网络地址改为这个名字。

练习 8.15： 如果一个客户端没有及时地读取数据可能会导致所有的客户端被阻塞。修改broadcaster来跳过一条消息，而不是等待这个客户端一直到其准备好写。或者为每一个客户端的消息发出channel建立缓冲区，这样大部分的消息便不会被丢掉；broadcaster应该用一个非阻塞的send向这个channel中发消息。
*/

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

//type client chan<- string // an outgoing message channel
// ex8.12
type client struct {
	ch       chan<- string
	Name     string
	LastTime time.Time
}

var (
	entering = make(chan client) // 管道的管道
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli.ch <- msg
			}
		case cli := <-entering:
			// cli 是个地址
			clients[cli] = true
			cli.ch <- "welcome: "
			for cli := range clients {
				cli.ch <- cli.Name
			}

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	cli := client{ch, who, time.Now()}
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- cli

	input := bufio.NewScanner(conn)

	go func() {
		// ex8.13
		timeout := 8.0 * time.Second
		ticker := time.NewTicker(timeout)
		for {
			<-ticker.C
			dur := time.Now().Sub(cli.LastTime)
			fmt.Println(dur.Seconds(), timeout.Seconds())
			if dur.Seconds() > timeout.Seconds(){
				conn.Close()
				break;
			}
		}
	}()
	for input.Scan() {
		messages <- who + ": " + input.Text()
		cli.LastTime = time.Now()
	}

	// NOTE: ignoring potential errors from input.Err()
	leaving <- cli
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}
