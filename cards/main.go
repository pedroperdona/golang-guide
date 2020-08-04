package main

func main() {
	cards := newDeck()
	cards.print()
	cards.suffle()
	cards.print()
}
