package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}

var mx = sync.RWMutex{}

var counter = 0

func Routines() {
	runtime.GOMAXPROCS(100)
	for i := 0; i <= 10; i++ {
		wg.Add(2)
		mx.RLock()
		go sayHello(i)
		mx.RLock()
		go incrementCounter()
		wg.Wait()

	}
}

func sayHello(i int) {
	fmt.Printf("Hello #%v  Counter: %v \n", i, counter)
	fmt.Println()
	mx.RUnlock()
	wg.Done()
}

func incrementCounter() {
	counter++
	mx.RUnlock()
	wg.Done()
}
