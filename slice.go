package main

import "fmt"

func Slice() {
	var scores = []int{1, 3, 1, 4, 5, 6}
	fmt.Println(scores, len(scores))
	scores = append(scores, 2)
	fmt.Println(scores, len(scores))

	topThree := scores[1:4]
	fmt.Println(topThree)
}
