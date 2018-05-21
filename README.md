# Learning Go (golang)

[Go: The Complete Developer's Guide (Golang)](https://www.udemy.com/go-the-complete-developers-guide/)

## 01-helloworld
>`go run main.go`

## 02-cards
>Libraries used: `fmt`, `io/ioutil`, `math/rand`, `os`, `strings`,`time`

>`go run main.go deck.go`

```go
var card1 = "Ace of Spades"
card1 = "Ace of Spades"

card2 := "Ace of Spades"
card2 = "Ace of Spades2"

card3 := newCard()

// slice
cards := []string{newCard(), "Ace of Diamonds"}
cards = append(cards, "Six of Spades")

// for loop
for i, card := range cards {
    fmt.Print(i, " "+card+", ")
}
```

## 04-structs (and pointers)
>`go run main.go deck.go`

```go
type person struct {
	firstName string
	lastName  string
}
alex1 := person{
    firstName: "Alex",
    lastName:  "Anderson",
}
var alex2 person
alex2.firstName = "Alex"
alex2.lastName = "Anderson"
```
>**go** is a *pass by value* language. If we need to access to an object by reference, it is necessary to specify it
>`&variable` is the **memory address** of the value this variable is pointing at. i.e.: x000A
>`variablePointer := &variable`

>`*variable` is the **value** this memory address is pointing at. i.e.: "Mario"

>`variable *type` means that `variable` is a pointer to a `person`
>`func (variable *type) function {...}`

## 05-maps
```go
colors1 := map[string]string {
    "red": "#ff0000",
    "green": "#00ff00",
}
var colors2 map[string]string
colors3 := make(map[string]string)
colors3["red"] = "#ff0000"
delete(colors3, "red")
```

## 06-07-interfaces
```go
type shape interface {
	printArea() float64
}
func (t triangle) printArea() float64 {
    // ...
}
func (s square) printArea() float64 {
    // ...
}
```

## 08-http

>Libraries used: `fmt`, `io`, `net/http`, `os`
