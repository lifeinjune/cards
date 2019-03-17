package main

import "fmt"

//new type call 'deck' created
type deck []string //deck is slice of strings
//new deck type behave just like string

func newDeck() deck { // function for handout new deck
	cards := deck{} //Create variable with deck type
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	//Create slice of all card suits in card
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "J", "Q", "K"}
	//Create slice of all values in card

	for _, suit := range cardSuits { //loop through each suit in slice
		for _, value := range cardValues { // loop through each value in slice
			cards = append(cards, value+" of "+suit) //append value and suit in to card variable
		}
	}
	return cards //returns cards slice with new cards

}

func (d deck) print() { //function with receiver
	//(d deck) is receiver of this function
	//this setup that any type of deck can access to this function
	//(d deck): deck is type, d is variable that actual variable of deck
	// d is actual copy of deck
	// by go convension use single or two word of receiver as variable
	// this case receiver is deck so variable is d
	for i, card := range d { //same as previous that go through cards in the deck
		fmt.Println(i, card) //print out index and card
	}
}

func deal(d deck, handSize int) (deck, deck) { // deal function argument receive deck and int
	//(deck, deck) return two variables
	return d[:handSize], d[handSize:] //splite deck by using range
}
