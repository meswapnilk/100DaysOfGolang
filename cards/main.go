package main

import "fmt"

func main() {
	cards := newDeck()
	hand, remainingDeck := deal(cards, 13)
	// hand.print()
	// remainingDeck.print()
	fmt.Println(hand.toString())
	fmt.Println(remainingDeck.toString())
}

func newCard() string {
	return "Five of Diamonds"
}
