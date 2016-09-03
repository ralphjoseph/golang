// functions project main.go
package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}
func add3(a, b, c int) int {
	return a + b + c
}

func muliplereturn() (int, int) {
	return 6, 10
}

//varidiac function
func sum(nums ...int) int {
	fmt.Println(nums)

	//iterate using range for
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
	return total

}

//anonymous functions

func nextNum() func() int {
	i := 1
	return func() int {
		i += 1
		return i
	}
}

//recursive function
func factorial(num int) int {
	if num == 0 {
		return 1
	}

	return num * factorial(num-1)
}

func main() {
	fmt.Println(add(1, 4))
	fmt.Println(add3(1, 4, 9))

	x, y := muliplereturn()
	fmt.Println(x, y)

	sum(9, 0)
	sum(100, 2300, 3434)
	sum(7686, 121, 121, 11, 11, 11)

	fun := nextNum()
	fmt.Println(fun())
	fmt.Println(fun())
	fmt.Println(fun())
	fmt.Println(fun())
	fmt.Println(fun())
	fmt.Println(fun())

	fmt.Println(factorial(8))
	fmt.Println(factorial(6))

}
