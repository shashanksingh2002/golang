package main
import "fmt"

func Variables() {
	var card string = "Ace of Spades" // var declares 1 or more variables.
	// var is followed by the name of the variable (card), then the type (string), then the value (Ace of Spades).
	// Every statement in Go ends with a semicolon, but they are usually omitted. The Go compiler automatically inserts semicolons where they are needed.

	//Alternative way to declare and initialize a variable
	newCard := "Ace of Spades" // := is shorthand for declaring and initializing a variable. It can only be used inside a function.

	//Reassigning a variable
	card = "Five of Diamonds"
	newCard = "Five of Diamonds"

	//Printing to the console
	// fmt.Println(card) //Println is a function in the fmt package that prints to the console. It adds a newline at the end of the output.
	// fmt. is used to access functions and variables in the fmt package.
	println(card, newCard)

	fmt.Println(card, newCard)
}