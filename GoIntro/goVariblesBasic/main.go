package main

var card2 string // variables can be initialized outside of a function, but cannot be assigned a varible.
func basicOfVariables() string { // string is the return type
	//var card string = "Ace of Apades"
	card := "Ace of Apades" // it's the same thing than the line above
	card2 = "Five of Gold"  // := only is used in the declaration
	return card
}
