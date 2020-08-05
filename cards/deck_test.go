package main

import (
	"os"
	"testing"
)

const deckTestingFilename = "_decktesting"

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	if len(deck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(deck))
	}

	if deck[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", deck[0])
	}

	if deck[len(deck)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", deck[len(deck)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove(deckTestingFilename)

	deck := newDeck()
	deck.saveToFile(deckTestingFilename)

	loadedDeck := newDeckFromFile(deckTestingFilename)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, but got %v", len(loadedDeck))
	}

	os.Remove(deckTestingFilename)
}
