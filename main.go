package main

import (
	"fmt"

	"github.com/farshadahmadi/cards/deckstringslice"
)

func main() {
	// cards := newDeck()

	// handDeck, remainingDeck := deal(cards, 4)
	// handDeck.print()
	// remainingDeck.print()
	// fmt.Println(handDeck)

	//err := cards.saveToFile("my_cards")
	//if err != nil {
	//	fmt.Println(err)
	//}

	myCards, err := deckstringslice.ReadFromFile("my_cards")
	if err != nil {
		fmt.Println(err)
		// os.Exit(1)
		return
	}

	myCards.Shuffle()

	myCards.Print()
}
