package main

type decks []string // creating a new type called deck which is a slice of strings
type cards string

// What is a receiver function?

// A receiver function is a function that is associated with a specific type. It allows you to define methods on that type, enabling you to call the function using the dot notation on instances of that type. In Go, receiver functions are defined by specifying a receiver argument before the function name. The receiver argument is typically a value or a pointer of the type you want to associate the method with.

// In this case, the receiver function printDeck is associated with the deck type. This means that you can call the printDeck function on any variable of type deck using the dot notation, like cards.printDeck() in the main function.

func (d decks) print(){
	for _, card := range d {
		println(card)
	}
}

// this or self is represented by d here

func (d decks) append(c cards, count int) decks {
	println("Appending card:", c, "count:", count)
	return append(d, string(c))
}

// Return the len of the array

func (d decks) len() int {
	return len(d)
}