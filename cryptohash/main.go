// cryptohash project main.go
package main

import (
	"crypto/sha1"
	"fmt"
)

var p = fmt.Println

func main() {
	s := "hello there! pleased to meet u"
	hash := sha1.New()
	hash.Write([]byte(s))
	hashVal := hash.Sum(nil)
	p(s)
	p(hashVal)
	fmt.Printf("%x", hashVal)

}
