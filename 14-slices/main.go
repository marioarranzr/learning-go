package main

import "fmt"

// Saiyan stuct
type Saiyan struct {
	Name   string
	Power  int
	Father *Saiyan
}

func main() {

	names := []string{"leto", "jessica", "paul"}
	names = append(names, "mario")
	fmt.Println(names)
	// [leto jessica paul mario]

	checks := make([]bool, 10) // Book 10 spaces and write the default value in them
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

}
