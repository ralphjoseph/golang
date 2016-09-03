// goroutines_channels project main.go
package main

import (
	"fmt"
)

func fun(str string) {
	for i := 0; i < 5; i++ {
		fmt.Println(str, i)
	}
}
func main() {
	//go routines
	go fun("first")

	//spawn anonymous function
	go func(str string) {
		fmt.Println(str)
	}("second")

	fun("simple")

	var input string
	fmt.Scanln(&input)

	//try channels
	mesg := make(chan string, 2)

	mesg <- "hello"
	mesg <- "there"
	go func() {
		mesg <- "last write"
	}()

	fmt.Println(<-mesg)
	fmt.Println(<-mesg)
	fmt.Println(<-mesg)

	fmt.Scanln(&input)

}
