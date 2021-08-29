package main

import "fmt"

func main() {
	cards := newDeck()
	// hand, cards := deal(cards, 5)

	// cards.print()
	// hand.print()
	fmt.Println(cards.saveToFile("test"))
}
