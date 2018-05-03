package main

import "fmt"

func main() {
	// var card = "Ace of Spades"
	// card := "Ace of Spades"
	card := newCard()
	fmt.Println(card)
}

func newCard() string {
	return "Ace of Spades"
}
