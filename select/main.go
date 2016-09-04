// select project main.go
package main

import (
	"fmt"
	"time"
)

func main() {
	//read from channels using select
	//1. blocking calls no timeout

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "hello first"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "hello second"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)

		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}

	//2. blocking calls with timeout
	go func() {
		time.Sleep(time.Second * 5)
		c1 <- "hello first"
	}()

	select {
	case msg1 := <-c1:
		fmt.Println(msg1)
	case <-time.After(time.Second * 2):
		fmt.Println("time out in reading from c1")
	}

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "hello second success"
	}()

	select {
	case msg2 := <-c2:
		fmt.Println(msg2)
	case <-time.After(time.Second * 4):
		fmt.Println("time out in reading from c2")
	}

	//3. non blocking calls

	go func() {
		//time.Sleep(time.Second * 2)
		c1 <- "c1 msg"
		c2 <- "c2 msg"
	}()

	time.Sleep(time.Second * 2)

	for i := 0; i < 10; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		default:
			fmt.Println("no msg")
		}
	}
}
