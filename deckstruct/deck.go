package deckstruct

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Deck struct {
	Cards []string
}

func newDeck() *Deck {
	cards := make([]string, 0)

	cardSuits := [4]string{"Diamonds", "Hearts", "Spades", "Clubs"}
	cardValues := [5]string{"Ace", "Two", "Three", "Four", "Five"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return &Deck{
		Cards: cards,
	}
}

func (d *Deck) Print() {
	for i, card := range d.Cards {
		fmt.Println(i, card)
	}
}

func deal(d *Deck, handSize int) (*Deck, *Deck) {
	return &Deck{Cards: d.Cards[:handSize]}, &Deck{Cards: d.Cards[handSize:]}
}

func (d *Deck) toJsonString() string {
	dBytes, _ := json.Marshal(d)
	return string(dBytes)
}

func (d *Deck) toCustomString() string {
	return strings.Join(d.Cards, ",")
}

func (d *Deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toCustomString()), 0666)
}

func ReadFromFile(filename string) (*Deck, error) {
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	fileString := string(fileBytes)
	return &Deck{Cards: strings.Split(fileString, ",")}, nil
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	cards := d.Cards
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
}
