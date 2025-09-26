package main

import "testing"

func TestNewDeck(t *testing.T) {
	cards := NewDeck()
	if len(cards) != 56 {
		t.Errorf("Expected deck length of 56, but got %d", len(cards))
	}
}

func TestDeal(t *testing.T) {
	cards := NewDeck()
	handSize := 5
	hand, remainingDeck := cards.Deal(handSize)

	if len(hand) != handSize {
		t.Errorf("Expected hand size of %d, but got %d", handSize, len(hand))
	}

	if len(remainingDeck) != len(cards)-handSize {
		t.Errorf("Expected remaining deck size of %d, but got %d", len(cards)-handSize, len(remainingDeck))
	}
}