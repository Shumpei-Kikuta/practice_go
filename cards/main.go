package main

func main() {
	// var card string = "Ace of Spades" // var val_name type
	// bool, string, int, float64
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades") // slice is immutable

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
