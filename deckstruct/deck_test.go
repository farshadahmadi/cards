package deckstruct

import (
	"fmt"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d.Cards) != 20 {
		t.Errorf("Expected 20, but got %d", len(d.Cards))
	}

	if d.Cards[0] != "Ace of Diamonds" {
		t.Errorf("Expected the first element as Ace of Diamonds but got %s", d.Cards[0])
	}

	if d.Cards[len(d.Cards)-1] != "Five of Clubs" {
		t.Errorf("Expected the last element as Five of Clubs but got %s", d.Cards[len(d.Cards)-1])
	}
}

// an example of a unit test
func TestToJsonString(t *testing.T) {
	d := Deck{
		Cards: []string{"a", "b", "c"},
	}

	actualJsonStr := d.toJsonString()

	expectedJsonStr := fmt.Sprintf("{%q:[%q,%q,%q]}", "Cards", "a", "b", "c")
	// expectedJsonStr := "[\"a\",\"b\",\"c\"]"

	if actualJsonStr != expectedJsonStr {
		t.Errorf("Expected the json string to be %s but got %s", expectedJsonStr, actualJsonStr)
	}
}

// an example of a integration test
func TestSaveToFileAndReadFromFile(t *testing.T) {
	testFileName := "_testdeck"

	err := os.Remove(testFileName)
	if err != nil && !os.IsNotExist(err) {
		t.Errorf("Error while deleting test file %v", err)
	}

	expectedDeck := Deck{Cards: []string{"a", "b", "c"}}
	err = expectedDeck.saveToFile(testFileName)
	if err != nil {
		t.Errorf("Cannot save Deck to file")
	}

	var actualDeck *Deck
	actualDeck, err = ReadFromFile(testFileName)
	if err != nil {
		t.Errorf("Cannot read Deck from file")
	}

	if len(actualDeck.Cards) != 3 || actualDeck.Cards[0] != expectedDeck.Cards[0] || actualDeck.Cards[2] != expectedDeck.Cards[2] {
		t.Errorf("Cannot read Deck from file")
	}

	err = os.Remove(testFileName)
	if err != nil && !os.IsNotExist(err) {
		t.Errorf("Error while deleting test file %v", err)
	}
}
