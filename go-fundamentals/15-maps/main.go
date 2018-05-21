package main

import "fmt"

// Friends collection
type Friends map[string]*Saiyan

// Saiyan struct
type Saiyan struct {
	Name    string
	Friends Friends
}

func main() {

	goku := &Saiyan{
		Name:    "Goku",
		Friends: make(Friends), // make(map[string]*Saiyan),
	}
	vegeta := &Saiyan{
		Name: "Vegeta",
	}
	krillin := &Saiyan{
		Name: "Krillin",
		Friends: Friends{
			"Goku": goku,
		},
	}
	fmt.Println(goku)
	fmt.Println(krillin)
	fmt.Println(vegeta)

}

// Redefine String method to Saiyan struct
func (s Saiyan) String() string {
	return fmt.Sprintf("%s %s", s.Name, s.Friends)
}

// Redefine String method to Friends type
func (f Friends) String() string {
	if len(f) == 0 {
		return ""
	}
	s := "(Friends:"
	isFirst := true
	for k := range f {
		if !isFirst {
			s += ", "
		} else {
			isFirst = false
		}
		s += k
	}
	s += ")"
	return s
}
