// jobs project main.go
package main

import (
	"fmt"
)

func main() {
	//send jobs via channel and close it when done
	jobs := make(chan int)
	done := make(chan bool)

	go func() {
		//read jobs
		for {
			job, more := <-jobs
			if more {
				fmt.Println("got job", job)
			} else {
				fmt.Println("got all jobs")
				done <- true
				return // break out
			}
		}
	}()

	//create jobs
	for i := 0; i < 5; i++ {
		jobs <- i + 1
	}
	close(jobs) //this will help inn identifying the last msg
	<-done      //wait till everything is done

	//range on channels
	nums := make(chan int, 10)

	for i := 0; i < 10; i++ {
		nums <- i
	}
	close(nums)

	for data := range nums {
		fmt.Println(data)
	}
}
