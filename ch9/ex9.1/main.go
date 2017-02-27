package main

/*
练习 9.1：
给gopl.io/ch9/bank1程序添加一个Withdraw(amount int)取款函数。
其返回结果应该要表明事务是成功了还是因为没有足够资金失败了。
这条消息会被发送给monitor的goroutine，且消息需要包含取款的额度和一个新的channel，
这个新channel会被monitor goroutine来把boolean结果发回给Withdraw。
*/

import (
	"fmt"
	"time"
)

type WithDrawInfo struct {
	quota int
	flag  bool
}

var deposits = make(chan int)   // send amount to deposit
var balances = make(chan int)   // receive balance
var withdraws = make(chan WithDrawInfo) //
var withdrawresult = make(chan WithDrawInfo) //
func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func WithDraw(amount int) bool {
	withdraws <- WithDrawInfo{amount, false}
	info := <-withdrawresult
	return info.flag
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			fmt.Printf("Deposit = %d\n", amount)
			balance += amount
		case balances <- balance:
		case info := <-withdraws:
			fmt.Printf("Withdraw = %d\n", info.quota)
			if balance >= info.quota {
				info.flag = true
				balance -= info.quota
			} else {
				info.flag = false
				info.quota = balance
			}
			withdrawresult <- info
		}

	}
}

func init() {
	go teller() // start the monitor goroutine
}

func main() {
	done := make(chan struct{})

	// Alice
	go func() {
		num := 200
		Deposit(num)
		done <- struct{}{}
	}()

	// Bob
	go func() {
		num := 500
		Deposit(num)
		done <- struct{}{}
	}()

	go func() {
		_ = time.Second
		num := 100
		_ = WithDraw(num)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done
	<-done
	if got, want := Balance(), 300; got != want {
		fmt.Printf("Final Balance = %d\n", got)
	}
}
