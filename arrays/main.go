// arrays project main.go
package main

import (
	"fmt"
)

func main() {

	var arr [5]int

	arr[3] = 86

	//print arr
	fmt.Println(arr)
	fmt.Println(arr[3])

	arr2 := [4]int{1, 4, 6}

	//print arr2
	fmt.Println(arr2)
	fmt.Println(arr2[3])

	var twod [4][5]int

	for i := 0; i < 4; i++ {
		for j := 0; j < 5; j++ {
			twod[i][j] = i*10 + j
		}
	}
	fmt.Println(twod)
}
