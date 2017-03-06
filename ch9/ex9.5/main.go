package main
/*
练习 9.5:
写一个有两个goroutine的程序，两个goroutine会向两个无buffer channel反复地发送ping-pong消息。这样的程序每秒可以支持多少次通信？
*/
import (
	"fmt"
	"time"
)

var done = make(chan struct{}, 2)

func ping(in, out chan uint64) {
	for {
		select {
		case s := <-in:
			// fmt.Println("ping")
			out <- s+1
		case <-done:
			break;
		}
	}
}

func pong(out, in chan uint64) {
	for {
		select {
		case s:= <-out:
			// fmt.Println("pong")
			in <- s
		case <- done:
			break;
		}
	}
}

func main() {
	in := make(chan uint64)
	out := make(chan uint64)

	go ping(in, out)
	go pong(out, in)
	ticker := time.NewTicker(1 * time.Second)
	in <- 0
	for {
		select {
		case <-ticker.C:
			fmt.Println("total: ", <-in)
			in <- 0

		}
	}
	// done<-struct{}{}
	// done<-struct{}{}
	// close(done)
	// close(in)
	// close(out)
}
