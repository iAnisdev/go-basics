package main

import (
	"fmt"
	"sync"
)

var waitgroup = sync.WaitGroup{}

func testChannels() {
	ch := make(chan string)
	waitgroup.Add(2)

	go func() {
		i := <-ch
		fmt.Println(i)
		waitgroup.Done()
	}()

	go func() {
		ch <- "Sending positive wibes"
		waitgroup.Done()
	}()
	waitgroup.Wait()
}
