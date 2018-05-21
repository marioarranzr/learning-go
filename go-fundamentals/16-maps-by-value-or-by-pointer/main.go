package main

import (
	"github.com/fatih/color"
)

// Person type
type Person struct {
	firstName string
	lastName  string
}

func changeNameByPointer(p *Person) {
	p.firstName = "BobPointer"
	customColor := color.New(color.FgRed)
	customColor.Println("Inside the function:", *p)

}

func changeNameByValue(p Person) {
	p.firstName = "BobValue"
	customColor := color.New(color.FgGreen)
	customColor.Println("Inside the function:", p)
}

func main() {
	person := Person{
		firstName: "Alice",
		lastName:  "Dow",
	}

	changeNameByPointer(&person) // change the real object
	customColor := color.New(color.FgRed)
	customColor.Println(person) // It changes. It returns {BobPointer Dow} instead of {Alice Dow}

	changeNameByValue(person)   // makes a copy
	customColor.Println(person) // Desn't change. It still returns {BobPointer Dow}

}
