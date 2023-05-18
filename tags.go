package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Name   string `required max: "100"`
	Origin string
}

func Tags() {
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")

	fmt.Println(field)
	fmt.Println(field.Tag)
}
