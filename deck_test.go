package main

import (
	"os"
	"testing"
)

func TestNewDeck (t *testing.T ) {
	d := newDeck()

	if len(d) != 12 {
		t.Errorf("Expected deck length of 13, but got %v", len(d))
	}

	if d[0] != "Spades of Ace" {
		t.Errorf("expected first card Ace of Spades, but got %v", d[0])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 12 {
		t.Errorf("expected 16 cards in deck, got %v", len(loadedDeck))
	}
}
