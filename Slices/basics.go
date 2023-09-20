package main

import (
	"fmt"
	"slices"
)

// func Contains[T comparable](s []T, e T) bool {
// 	for _, v := range s {
// 		if v == e {
// 			return true
// 		}
// 	}
// 	return false
// }

func main() {
	list := []int{1, 2, 3, 4, 5}

	for _, v := range list {
		fmt.Println(v)
	}

	fmt.Println(slices.Contains(list, 3))

	hasNegative := slices.ContainsFunc(list, func(i int) bool { return i < 0 })
	fmt.Println("hasNegative: ", hasNegative)

	hasOdd := slices.ContainsFunc(list, func(i int) bool { return i%2 != 0 })

	fmt.Println("hasOdd:", hasOdd)

}
