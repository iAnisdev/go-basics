package main

import "fmt"

type Number interface {
	int64 | float64
}

func sumOf[T Number](list []T) T {
	var result T

	for _, val := range list {
		result += val
	}

	return result
}

func Generics() {
	fmt.Printf("Sum is %v", sumOf([]int64{1, 2, 3, 4, 5, 6, 7}))
	fmt.Println()
	fmt.Printf("Sum is %v", sumOf([]float64{3.2, 2.1, 3.8, 4.5, 5.9, 6.3, 7.4}))
	fmt.Println()
}
