package main

import (
	"os"
	"testing"
)

// t is test handler
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	} else if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but found %v", d[0])
	} else if d[15] != "Four of Clubs" {
		t.Errorf("Expected first card to be Four of Clubs, but found %v", d[15])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 160 {
		t.Errorf("Expected 16 cards, got %v", len(loadedDeck))
	}

	os.Remove("_decttesting")
}
