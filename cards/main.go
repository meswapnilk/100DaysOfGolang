package main

func main() {
	cards := newDeck()
	hand, remainingDeck := deal(cards, 13)
	hand.print()
	remainingDeck.print()
}

func newCard() string {
	return "Five of Diamonds"
}
