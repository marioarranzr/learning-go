package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	contactInfo
}
type contactInfo struct {
	email 	string
	zipCode int
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email: "a@a.es",
			zipCode: 90000,
		},
	}

	// alexPointer := &alex
	// alexPointer.updateName("Tony")
	alex.updateName("Tony")
	alex.print()
}

func (pointerToPerson *person) updateName(firstName string) {
	(*pointerToPerson).firstName = firstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
