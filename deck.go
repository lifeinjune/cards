package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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
	//_,suit on the loop, _ means that ignore the value return becuas it is not going to be use
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

func (d deck) toString() string { //function to convert deck to string
	return strings.Join([]string(d), ",")
	//[]string used to convert deck type to slice of string
	//string.Join function used to join up the slice of string to
	//one string seperate by ","
	// returns deck as one string
}

func (d deck) saveToFile(filename string) error { // receiver function to save the deck as string
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
	// Use standard package call WriteFile to save the deck as file
	// if error comes up it will return the error
	// 0666 standard permission that everyone can read and write the file
	// receive the filename argument as string which is name of the file
	//[]byte to convert string to the slice of bytes
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename) //ReadFile will return byte slice and error
	if err != nil {                      // if there is error err will be not nil
		// log the error and call newDeck() Op1
		// log error and quit program Op2
		fmt.Println("Error!", err) //print out the error message
		os.Exit(1)                 //program quit function when return other then 0
		//function from os package that can be work well on any other platforms
	}
	ss := strings.Split(string(bs), ",")
	//strings.Split function split the one string and convert into
	//slice of strings and put into ss slice
	return deck(ss) //slice of string convert into deck and return
}

func (d deck) shuffle() { //function to shuffle cards in the deck
	source := rand.NewSource(time.Now().UnixNano()) // create new source for random number generation
	//which using current time that keep changes
	r := rand.New(source) //create random type variable with new source
	for i := range d {    //loop through each element position
		newPosition := r.Intn(len(d) - 1) //generate random number
		d[i], d[newPosition] = d[newPosition], d[i]
		//replace current position card with random number position card
	}
}
