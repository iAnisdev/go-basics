package main

import "fmt"

func Slice() {
	var scores = []int{1, 3, 1, 4, 5, 6}
	fmt.Println(scores, len(scores))
	fmt.Println(scores, cap(scores))
	scores = append(scores, 2)
	fmt.Println(scores, len(scores))
	fmt.Println(scores, cap(scores))

	topThree := scores[:3]
	lastFour := scores[4:]
	between := scores[1:3]
	fmt.Println(topThree)
	fmt.Println(lastFour)
	fmt.Println(between)
}
