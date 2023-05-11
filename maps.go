package main

import "fmt"

func Maps() {
	user := map[string]string{
		"name": "ABC",
		"age":  "DEF",
	}

	fmt.Println(user)
	fmt.Println(user["name"])
	user["name"] = "GHI"
	fmt.Println(user["name"])
	user["Address"] = "JKL"
	fmt.Println(user)
	delete(user, "age")
	fmt.Println(user)

	name, isNameAvailable := user["name"]
	if isNameAvailable {
		fmt.Println(name)
	}
	fmt.Println(len(user))

	menu := map[string]int{
		"Rice":  12,
		"Water": 5,
	}
	for k, v := range menu {
		fmt.Println(k, v)
	}
}
