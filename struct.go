package main

import "fmt"

type Doctor struct {
	id        int
	name      string
	available []string
}

type Surgeon struct {
	Doctor
	manageDoctor []int
}

func useStruct() {
	john := Doctor{
		id:        1,
		name:      "John",
		available: []string{"21", "22", "23"},
	}
	max := Doctor{
		2,
		"Max",
		[]string{"19", "20"},
	}
	fmt.Println(john)
	fmt.Println(john.id)
	fmt.Println(max)
	fmt.Println(max.available)

	DoctorList := []Doctor{john, max}
	fmt.Println((DoctorList))

	Alex := struct{ name string }{name: "Alex"}
	fmt.Println(Alex)

	mathew := Surgeon{
		Doctor: Doctor{id: 4,
			name:      "Mathew",
			available: []string{"21", "22", "23"}},
		manageDoctor: []int{1, 2},
	}
	fmt.Println(mathew)
	fmt.Println(mathew.id)
	fmt.Println(mathew.name)
	fmt.Println(mathew.manageDoctor)
}
