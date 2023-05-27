package main

import "fmt"

func TestFunction() {
	result, error := Divide(5.0, 0.0)
	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Println(result)
}

func Divide(a, b float64) (float64, error) {
	if a > 0 && b > 0 {
		return a / b, nil
	}
	return 0.0, fmt.Errorf("both number must be greater than 0")
}
