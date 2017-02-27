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
	IP       string
	LastTime time.Time
	NickName string
}

var (
	entering = make(chan client) // 管道的管道
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	// ex8.15
	clients := make(map[string]client) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			fmt.Println(msg)
			for _, cli := range clients {
				select {
				case cli.ch <- msg:
					default:
				}
			}
		case cli := <-entering:
			// cli 是个内存地址
			ip := cli.IP
			clients[ip] = cli
			cli.ch <- "welcome: "
			for _, cli := range clients {
				cli.ch <- cli.NickName
			}

		case cli := <-leaving:
			nick := cli.NickName
			fmt.Println(nick + " has left")
			delete(clients, cli.IP)
			close(cli.ch)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	closed := make(chan struct{})
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	cli := client{ch, who, time.Now(), "guest"}

	// ex8.14
	input := bufio.NewScanner(conn)
	fmt.Fprintln(conn, "input nick name")
	input.Scan()
	cli.NickName = input.Text()
	ch <- "You are " + cli.NickName
	messages <- cli.NickName + " has arrived"
	entering <- cli

	go func() {
		// ex8.13
		timeout := 60.0 * time.Second
		ticker := time.NewTicker(timeout)
		for {
			select {
			case <-ticker.C:
				dur := time.Now().Sub(cli.LastTime)
				fmt.Println(dur.Seconds(), timeout.Seconds())
				if dur.Seconds() > timeout.Seconds() {
					closed <- struct{}{}
				}
			case <-closed: 
				messages <- cli.NickName + " has left"
				leaving <- cli
				conn.Close()
				break;
			}
		}

	}()

	for input.Scan() {
		messages <- cli.NickName + ": " + input.Text()
		cli.LastTime = time.Now()
	}
	closed <- struct{}{}
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}
