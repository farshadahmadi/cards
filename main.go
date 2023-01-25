package main

import (
	"github.com/farshadahmadi/cards/customos"
	"github.com/farshadahmadi/cards/deckstructrefactored"
)

func main() {
	/* use deck string slice */

	// cards := newDeck()

	// handDeck, remainingDeck := deal(cards, 4)
	// handDeck.print()
	// remainingDeck.print()
	// fmt.Println(handDeck)

	//err := cards.saveToFile("my_cards")
	//if err != nil {
	//	fmt.Println(err)
	//}

	/*	myCards, err := deckstringslice.ReadFromFile("my_cards")
		if err != nil {
			fmt.Println(err)
			// os.Exit(1)
			return
		}

		myCards.Shuffle()

		myCards.Print()*/

	/* use deck struct */

	// fileService := customos.NewMockFileService()
	fileService := customos.NewFileService()

	d := deckstructrefactored.NewDeck(fileService)

	d.SaveToFile("asdfdsasd")

	//myAnotherCards, err := deckstruct.ReadFromFile("my_cards")
	//if err != nil {
	//	fmt.Println(err)
	//	// os.Exit(1)
	//	return
	//}
	//
	//myAnotherCards.Shuffle()
	//
	//myAnotherCards.Print()
}
