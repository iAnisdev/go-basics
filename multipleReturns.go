package main

import "fmt"

func Swap(x, y string) (string, string) {
	return y, x
}

func TestSwap() {
	fmt.Println(Swap("Hello", "World"))
	fmt.Println(Swap("Good", "Morning"))
}
