package main;

func main() {
	cardDeck := NewDeck();
	cardDeck.Print()
	deal1, deal2 := cardDeck.Deal(5);

	println("------- Deal 1 -------")
	deal1.Print();
	println("------- Deal 2 -------")
	deal2.Print();
}