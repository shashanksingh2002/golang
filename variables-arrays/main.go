package main

import "fmt"

func main() {
	card := newCard()
	fmt.Println(card)

	// Call the Variables function from the cards package
	Variables()
	// declare an array of strings
	Arrays()
}

func newCard() string {
	return "Five of Diamonds"
}