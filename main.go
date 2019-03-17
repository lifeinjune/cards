package main

func main() {
	//Go use Static types for variable so need to declare type of variable
	//and only contain that type of variable in
	//var card string = "Ace of Spades" (this is longer form)
	cards := newDeck() //slice declaration
	//cards = append(cards, "Six of spades") //add element in slice at the back of the slice
	//append returns new slice not modify slice
	//make the compiler to figure out what type is variable by value of it
	// := used only for new variable not reassign value of variable
	//var: short for variable that card is variable
	//card: name of variable
	//string: type of variable
	//"Acce of Spaces": value contain in variable
	/*for i, card := range cards {
	/*
		i: index of the slice
		card: value of the element in the slice
		range cards: loop through each element in data structure
	*/
	//fmt.Println(i, card) //print index and element
	// both index and card must be used
	//}
	cards.print()                         //calling function from deck.go
	hand, remainingDeck := deal(cards, 5) // because deal will returns two variables
	// first deck will be record on hand and another one to remaingdeck
	hand.print()
	remainingDeck.print()

}

/*func newCard() string { //after the () with space define return data type
	return "five of diamonds"
}*/

//functions can be called in different file but same package without
//import statement
