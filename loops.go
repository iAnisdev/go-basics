package main

import "fmt"

func Loops() {

	x := 0

	for x < 5 {
		fmt.Printf("X is currently %v \n", x)
		x++
	}

	for i := 0; i < 5; i++ {
		fmt.Printf("i is currently %v \n", i)
	}

	ages := []int{28, 43, 31, 19, 25, 29}

	for i := 0; i < len(ages); i++ {
		fmt.Printf("age is currently %v \n", ages[i])
	}

	for index, value := range ages {
		fmt.Printf("age at index #%v  is %v\n", index, value)
	}

	for _, value := range ages {
		fmt.Printf("age is %v\n", value)
	}
}
