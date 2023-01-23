package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := [4]string{"Diamonds", "Hearts", "Spades", "Clubs"}
	cardValues := [5]string{"Ace", "Two", "Three", "Four", "Five"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toJsonString() string {
	dBytes, _ := json.Marshal(d)
	return string(dBytes)
}

func (d deck) toCustomString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toCustomString()), 0666)
}

func readFromFile(filename string) (deck, error) {
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	fileString := string(fileBytes)
	return strings.Split(fileString, ","), nil
}

func (d deck) shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}
