package main

import (
	"fmt"
	"strings"
)

func String() {
	var message string = "hello friends, how are you?"
	fmt.Println(message)
	fmt.Println(strings.Contains(message, "hello"))
	fmt.Println(strings.ReplaceAll(message, "hello", "Hi"))
	fmt.Println(strings.ToUpper(message))
	fmt.Println(strings.ToLower(message))
	fmt.Println(strings.Split(message, " "))
	fmt.Println(strings.Index(message, "how"))
}
