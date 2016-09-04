// mutex project main.go
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	//read/write from map safely using mutex
	m := make(map[int]int)
	mut := &sync.Mutex{}

	var atomiVal int64

	go func() {
		tot := 0
		for {
			mut.Lock()
			tot += m[rand.Intn(5)]
			mut.Unlock()
			atomic.AddInt64(&atomiVal, 1)

		}
	}()

	for i := 0; i < 10; i++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(1000)

				mut.Lock()
				m[key] = val
				mut.Unlock()
				atomic.AddInt64(&atomiVal, 1)
			}
		}()
	}

	time.Sleep(time.Second)
	//get the atomicVal
	final := atomic.LoadInt64(&atomiVal)
	fmt.Println(final)
	mut.Lock()
	fmt.Println(m)
	mut.Unlock()

}
