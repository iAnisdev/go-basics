package main

import (
	"fmt"
	"strings"
)

func AbbrevName(name string) string {
	capitalizeString := strings.ToUpper(name)
	nameArr := strings.Split(capitalizeString, " ")
	var initials []string

	for _, v := range nameArr {
		initials = append(initials, v[:1])
	}
	return strings.Join(initials, ".")
}

func test() {
	fmt.Println(AbbrevName("Sam Harris"))
	fmt.Println(AbbrevName("Patrick Feenan"))
	fmt.Println(AbbrevName("Evan Cole"))
	fmt.Println(AbbrevName("P Favuzzi"))
	fmt.Println(AbbrevName("David Mendieta"))
}
