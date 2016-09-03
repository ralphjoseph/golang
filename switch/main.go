// switch project main.go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Is it today a weekday of weekend")

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}

	current_time := time.Now()
	switch {
	case current_time.Hour() < 12:
		fmt.Println("before noon")
	default:
		fmt.Println("after noon")
	}
}
