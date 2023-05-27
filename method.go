package main

import "fmt"

type person struct {
	name string
}

func Method() {
	anis := person{name: "Anis"}

	anis.greet()
}

func (u person) greet() {
	fmt.Printf("Hello %v", u.name)
}
