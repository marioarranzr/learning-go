package main

import "fmt"

// Saiyan stuct
type Saiyan struct {
	Name   string
	Power  int
	Father *Saiyan
}

func main() {
	gohan := Saiyan{ // gohan := &Saiyan{...}
		Name:  "Gohan",
		Power: 1000,
		Father: &Saiyan{
			Name:   "Goku",
			Power:  9001,
			Father: nil,
		},
	}
	gohan.Introduce()

}

// Introduce Saiyan
func (p *Saiyan) Introduce() {
	fmt.Printf("Hi, I'm %s\n", p.Name)
}
