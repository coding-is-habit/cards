package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected Deck length of 16, but got %v", len(d))
	}

	if d[0] != "One of Diamond" {
		t.Errorf("Expected first card of One of Daimond, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[0])
	}
}
