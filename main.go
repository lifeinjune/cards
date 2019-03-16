package main

import "fmt"

func main() {
	//Go use Static types for variable so need to declare type of variable
	//and only contain that type of variable in
	//var card string = "Ace of Spades" (this is longer form)
	card := newCard()
	//make the compiler to figure out what type is variable by value of it
	// := used only for new variable not reassign value of variable
	//var: short for variable that card is variable
	//card: name of variable
	//string: type of variable
	//"Acce of Spaces": value contain in variable
	fmt.Println(card)
}

func newCard() string { //after the () with space define return data type
	return "five of diamonds"
}
