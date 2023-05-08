package main

import (
	"fmt"
	"runtime"
)

func findOS() {
	fmt.Print("Checking OS \n")
	var os string = runtime.GOOS

	switch os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.\n", os)
	}
}
