package main

import (
	"fmt"
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

func TestToJsonString(t *testing.T) {
	d := deck{"a", "b", "c"}

	actualJsonStr := d.toJsonString()

	expectedJsonStr := fmt.Sprintf("[%q,%q,%q]", "a", "b", "c")
	// expectedJsonStr := "[\"a\",\"b\",\"c\"]"

	if actualJsonStr != expectedJsonStr {
		t.Errorf("Expected the json string to be %s but got %s", expectedJsonStr, actualJsonStr)
	}
}
