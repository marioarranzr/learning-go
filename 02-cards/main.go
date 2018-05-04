package main

import (
	"fmt"
)

func main() {
	cards := newDeck()
	hand, remainingCards := deal(cards, 5)
	fmt.Println(hand.toString())
	fmt.Println(remainingCards.toString())
	hand.saveToFile("my_hand")
	remainingCards.saveToFile("remaining_cards")
}
