package main

import (
	"bufio"
	"fmt"
	"os"
)

func Input() {
	reader := bufio.NewReader((os.Stdin))
	fmt.Println("Enter your name please")
	input, _ := reader.ReadString('\n')
	fmt.Println("Welcome ", input)
}
