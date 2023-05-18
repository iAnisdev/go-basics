package main

import "fmt"

func testDefer() {
	defer fmt.Println("!")
	defer fmt.Println("World")
	defer fmt.Println("Hello")
	a := "start"
	defer fmt.Println(a)
	a = "end"
}
