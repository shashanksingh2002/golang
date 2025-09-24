package main

import (
	"fmt"
)

func main() {
	randomString := "Hello, World!"
	fmt.Println([]byte(randomString))
	cardDeck := NewDeck()
	cardDeck.Print()
	deal1, deal2 := cardDeck.Deal(5)

	println("------- Deal 1 -------")
	deal1.Print()
	println("------- Deal 2 -------")
	deal2.Print()
	cardDeck.Save("my_cards.txt", cardDeck.Convert2Byte(), 0666)
	byteSlice, err := ReadFile("my_cards.txt")

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("File contents:", string(byteSlice))
	}
}
