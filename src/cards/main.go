package main

func main() {
	cards := newDeck()

	hand, remainingDeck := deal(cards, 5)

	hand.saveToFile("my_cards")
	remainingDeck.saveToFile("my_cards")

	remainingDeck.shuffle()
	remainingDeck.print()
}
