package main

import (
	"flag"
	"fmt"
	"sync"

	"github.com/bradfitz/iter"
	"github.com/wkharold/fnosvc/fno"
)

var fnos chan int
var done chan struct{}
var wg sync.WaitGroup

func init() {
	fnos = make(chan int)
	done = make(chan struct{})

	go fno.Generator(fnos, done)
}

func main() {
	calls := flag.Int("calls", 128, "number of calls per client")
	cc := flag.Int("clients", 8, "number of concurrent clients")
	flag.Parse()

	for cid := range iter.N(*cc) {
		wg.Add(1)
		go client(cid, fnos, *calls)
	}

	wg.Wait()

	close(done)
}

func client(cid int, fnos chan int, calls int) {
	defer wg.Done()
	for _ = range iter.N(calls) {
		fno := <-fnos
		fmt.Printf("client %d got fno %d\n", cid, fno)
	}
}
