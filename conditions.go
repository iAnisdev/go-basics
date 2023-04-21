package main

import "fmt"

func Conditions() {
	age := 53

	fmt.Println(age <= 30)
	fmt.Println(age >= 30)
	fmt.Println(age == 30)
	fmt.Println(age != 30)

	if age > 40 && age < 50 {
		fmt.Println("Age is between 40 and 50")
	} else if age > 50 {
		fmt.Println("Age is over 50")
	} else {
		fmt.Println("Age is below 40")
	}
}
