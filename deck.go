package main

import "fmt"

//new type call 'deck' created
type deck []string //deck is slice of strings
//new deck type behave just like string
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
