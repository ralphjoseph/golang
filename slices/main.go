// slices project main.go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	//slices
	s := make([]string, 6)
	fmt.Println("empty string", s)
	s[2] = "a"
	fmt.Println(s)
	s = append(s, "b")
	fmt.Println("s: ", s, "len", len(s))

	dest := make([]string, len(s))
	copy(dest, s)
	fmt.Println(dest)

	for i := 0; i < len(s); i++ {
		s[i] = strconv.Itoa(i)
	}

	//substring
	fmt.Println(s)

	//from index 1 to 5
	fmt.Println("print from index 1 to 5", s[1:5])
	fmt.Println("print from index begin to 5", s[:5])
	fmt.Println("print from index 3 to end", s[3:])
	fmt.Println("print from index begin to end", s[:])

	//create multidimensional array of varying length

	twod := make([][]int, 4)

	for i := 0; i < 4; i++ {
		innerlen := i + 1
		twod[i] = make([]int, innerlen)
		for j := 0; j < innerlen; j++ {
			twod[i][j] = i*10 + j
		}
	}

	fmt.Println(twod)
}
