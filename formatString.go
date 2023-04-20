package main

import "fmt"

func FormatString() {
	var name string = "Adam"
	var age float32 = 12.50

	fmt.Println("My Name is", name, "and i am", age, "years old!")
	fmt.Printf("My Name is %q and i am %0.1f years old \n", name, age)

	var message string = fmt.Sprintf("My Name is %q and i am %0.1f years old", name, age)

	fmt.Println(message)
}
