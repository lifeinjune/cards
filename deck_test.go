//test file end with _test.go
package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) { //test function start with uppercase
	d := newDeck()
	if len(d) != 52 { // if test fail
		t.Errorf("Expected deck length of 52 , but got %v", len(d))
		//tell test handler test is fail
	}

	if d[0] != "Ace of Spades" { // if test fail
		t.Errorf("Expected first card of Ace of Spaded, but got %v", d[0])
		//tell test handler test is fail
	}
	if d[len(d)-1] != "K of Clubs" { // if test fail
		t.Errorf("Expect last card K of Clubs, but got %v", d[len(d)-1])
		//tell test handler test is fail
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadDeck := newDeckFromFile("_decktesting")
	if len(loadDeck) != 52 { // if test fail
		t.Errorf("Expected deck length of 52 , but got %v", len(loadDeck))
		//tell test handler test is fail
	}
	os.Remove("_decktesting")
}
