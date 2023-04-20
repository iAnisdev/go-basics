package main

import (
	"fmt"
	"sort"
)

func Sort() {
	ages := []int{28, 43, 31, 19, 25, 29}

	fmt.Println(ages)
	sort.Ints(ages)
	fmt.Println(ages)
	index := sort.SearchInts(ages, 19)
	fmt.Println(index)
}
