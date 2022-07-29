package main

func main() {
	cards := newDeck()
	// hand, remainingDeck := deal(cards, 13)
	// hand.print()
	// remainingDeck.print()
	cards.saveToFile("cards.txt")
	newCards := newDeckFromFile("cards2.txt")
	newCards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
