// stateful_goroutines project main.go
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync/atomic"
	"time"
)

type read struct {
	key  int
	resp chan int
}

type write struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	var cnt int64 = 0

	//make two channels one for reading request and one for writing
	reads := make(chan *read)
	writes := make(chan *write)

	//process the requests here
	go func() {
		m := make(map[int]int)
		for {
			select {
			case r := <-reads:
				r.resp <- m[r.key]
			case w := <-writes:
				m[w.key] = w.val
				w.resp <- true
			}

		}
	}()

	//send read requests
	for i := 0; i < 100; i++ {
		go func() {
			for {
				r := &read{key: rand.Intn(10), resp: make(chan int)}
				reads <- r
				//get the response
				<-r.resp
				atomic.AddInt64(&cnt, 1)
				runtime.Gosched()
			}
		}()
	}

	//send write requests
	for i := 0; i < 10; i++ {
		go func() {
			for {
				w := &write{key: rand.Intn(10), val: rand.Intn(1000), resp: make(chan bool)}
				writes <- w
				//get the response
				<-w.resp
				atomic.AddInt64(&cnt, 1)
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)
	//get the atomicVal
	final := atomic.LoadInt64(&cnt)
	fmt.Println(final)

}
