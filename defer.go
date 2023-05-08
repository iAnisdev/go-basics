package main

import "fmt"

func testDefer() {
	defer fmt.Println("!")
	defer fmt.Println("World")
	defer fmt.Println("Hello")
}
