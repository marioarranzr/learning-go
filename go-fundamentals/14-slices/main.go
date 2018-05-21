package main

import (
	"fmt"
	"sort"
)

func main() {

	names := []string{"leto", "jessica", "paul"}
	names = append(names, "mario")
	fmt.Println(names)
	// [leto jessica paul mario]

	checks := make([]bool, 10) // Book 10 spaces and write the default value in them. Same as make([]bool, 10, 10)
	checks = append(checks, true)
	fmt.Println(checks)
	// [false false false false false false false false false false true]

	var names2 []string
	names2 = append(names2, "mario")
	fmt.Println(names2)
	// [mario]

	scores := make([]int, 0, 20) // Book 20 spaces, but the slice is empty
	scores = append(scores, 122)
	fmt.Println(scores)
	// [122]

	scores2 := []int{1, 2, 3, 4, 5}
	slice := scores2[2:4]
	slice[0] = 999
	fmt.Println(scores2)
	sort.Ints(scores2)
	fmt.Println(scores2)
	scores3 := make([]int, 2)
	copy(scores3, scores2[:2])
	fmt.Println(scores3)

}
