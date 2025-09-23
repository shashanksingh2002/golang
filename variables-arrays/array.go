package main


func Arrays() {
	cards := [] string {"Ace of Diamonds", newCard(), "Six of Spades"}
	
	cards = append(cards, "Five of Hearts") // append function adds a new element to the end of the slice and returns a new slice

	for i := 0; i < len(cards); i++ {
		println(cards[i])
	}
}