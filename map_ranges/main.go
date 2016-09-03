// map_ranges project main.go
package main

import (
	"fmt"
)

func main() {
	//create map
	m := make(map[string]string)
	m["hello"] = "world"
	m["u"] = "there"
	m["get"] = "me"

	val, err := m["hello"]
	fmt.Println(val, err)

	val, err = m["me"]
	if !err {
		fmt.Println(val)
	}

	//use range for loop

	for k, v := range m {
		fmt.Printf("%s -> %s\n", k, v)
	}

}
