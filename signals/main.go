// signals project main.go
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//create a channel for capturing signals
	sigs := make(chan os.Signal)

	done := make(chan bool)

	//notify the signal to the channel
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

	//spwan go routine to read the system signals
	go func() {
		s := <-sigs
		fmt.Println("received signal:", s)
		done <- true
	}()

	fmt.Println("waiting for signal")
	<-done
	fmt.Println("exit")
}
