package deckstructrefactored

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/farshadahmadi/cards/customos"
)

type Deck struct {
	Cards       []string
	fileService customos.FileReaderWriter
}

func NewDeck(FileService customos.FileReaderWriter) *Deck {
	cards := make([]string, 0)

	cardSuits := [4]string{"Diamonds", "Hearts", "Spades", "Clubs"}
	cardValues := [5]string{"Ace", "Two", "Three", "Four", "Five"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return &Deck{
		Cards:       cards,
		fileService: FileService,
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

func (d *Deck) SaveToFile(filename string) error {
	// return d.fileService.WriteFile(filename, []byte(d.toCustomString()), 0666)
	return d.fileService.WriteFile(filename, []byte(d.toJsonString()), 0666)
}

func (d *Deck) ReadFromFile(filename string) error {
	fileBytes, err := d.fileService.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(fileBytes, d)
	if err != nil {
		fmt.Println("Cannot convert string representation to deck")
	}

	return nil
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	cards := d.Cards
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
}
