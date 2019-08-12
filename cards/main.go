package main

func main() {
	// var card string = "Ace of Spades" // var val_name type
	// bool, string, int, float64
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades") // slice is immutable

	// cards := newDeck()
	// cards.saveToFile("my_cards")

	// // cards.print()
	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
