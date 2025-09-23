package main;

type deck []string;

func NewDeck() deck {
	cardDeck := deck{}

	cardNames := []string {"Heart", "Diamond", "Club", "Spade"};
	cardValues := []string {"Ace", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "King", "Queen", "Joker"}

	for _,cName := range cardNames {
		for _,cValue := range cardValues {
			cardDeck = append(cardDeck, cValue + " of " + cName);
		}
	}

	return cardDeck;
}

func (d deck) Print () {
	for _, card := range d {
		println(card)
	}
}

func (d deck) Shuffle() {
	
}

func (d deck) Deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}



