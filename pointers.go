package main

import "fmt"

func Pointers() {
	var a int = 27
	var b *int = &a
	fmt.Println(a, *b)
	a = 42
	fmt.Println(a, *b)
}
