package main

import (
	"fmt"
)

func main() {
	cards := newDeck()
	cards.shuffle()
	hand, remainingCards := deal(cards, 5)
	fmt.Println(hand.toString())
	fmt.Println(remainingCards.toString())
	hand.saveToFile("my_hand")
	remainingCards.saveToFile("remaining_cards")
	cards.saveToFile("cards")
	cards = newDeckFromFile("cards")
	cards.print()
}
