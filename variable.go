package main

import (
	"fmt"
	"strconv"
)

var (
/*
name   string = "A"
age    int    = 21
status bool   = false
size   byte   = 255
phone  rune   = 03452674175
*/
)

func Variables() {
	var a int
	a = 50

	var b string = "Bangkok"

	c := true
	d := 99.
	var e float32 = 43.12

	fmt.Printf("%v , %T \n", a, a)
	a = 55
	fmt.Printf("%v , %T \n", a, a)
	fmt.Printf("%v , %T \n", b, b)
	fmt.Printf("%v , %T \n", c, c)
	fmt.Printf("%v , %T \n", d, d)
	fmt.Printf("%v , %T \n", e, e)
	var f string = strconv.Itoa(a)
	fmt.Printf("%v , %T \n", f, f)

}
