package deckstringslice

import (
	"fmt"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 20 {
		t.Errorf("Expected 20, but got %d", len(d))
	}

	if d[0] != "Ace of Diamonds" {
		t.Errorf("Expected the first element as Ace of Diamonds but got %s", d[0])
	}

	if d[len(d)-1] != "Five of Clubs" {
		t.Errorf("Expected the last element as Five of Clubs but got %s", d[len(d)-1])
	}
}

// an example of a unit test
func TestToJsonString(t *testing.T) {
	d := deck{"a", "b", "c"}

	actualJsonStr := d.toJsonString()

	expectedJsonStr := fmt.Sprintf("[%q,%q,%q]", "a", "b", "c")
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

	expectedDeck := deck{"a", "b", "c"}
	err = expectedDeck.saveToFile(testFileName)
	if err != nil {
		t.Errorf("Cannot save deck to file")
	}

	var actualDeck deck
	actualDeck, err = ReadFromFile(testFileName)
	if err != nil {
		t.Errorf("Cannot read deck from file")
	}

	if len(actualDeck) != 3 || actualDeck[0] != expectedDeck[0] || actualDeck[2] != expectedDeck[2] {
		t.Errorf("Cannot read deck from file")
	}

	err = os.Remove(testFileName)
	if err != nil && !os.IsNotExist(err) {
		t.Errorf("Error while deleting test file %v", err)
	}
}
