package main

func main() {
	cards := deck{"ace of diamons", newCard()}
	cards = append(cards, "six of spades")

	cards.print()
}

func newCard() string {
	return "five of diamons"
}
