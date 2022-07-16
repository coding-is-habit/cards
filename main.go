package main

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")

	// card := newDeckFromFile("my_cards")
	// card.print()

	card := newDeck()
	card.shuffle()
	card.print()
}
