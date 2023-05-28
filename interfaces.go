package main

import "fmt"

type Book struct {
	name  string
	price int
}

type Dish struct {
	name  string
	price int
}

func (b Book) PrintInfo() {
	fmt.Printf("The Book %v is currently selling at %v$", b.name, b.price)
}

func (d Dish) PrintInfo() {
	fmt.Printf("The Dish %v is currently selling at %v$", d.name, d.price)
}

func InterfaceTest() {
	rdPd := Book{
		name:  "Rich Dad Poor Dad",
		price: 25,
	}

	rice := Dish{
		name:  "Rice",
		price: 2,
	}

	rdPd.PrintInfo()
	fmt.Println()
	rice.PrintInfo()
	fmt.Println()
}
