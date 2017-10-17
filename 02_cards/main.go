package main

func main() {
	//var card string = "Ace of Spades"
	//var card = "Ace of Spades"
	//card := "Ace of Spades"

	//card = newCard()

	//cards := deck{"Five of Diamonds", "Ace of Spades", newCard()}
	//cards = append(cards, "Six of Spades")

	//Print line
	//fmt.Println(cards)

	//For loops in go
	//cards.print()

	cards := newDeck()
	//cards.print()

	//hand, remainingDeck := deal(cards, 5)

	//hand.print()
	//remainingDeck.print()

	//I am returning the error from saveToFile
	//cardsTwo := newDeckFromFile("my")
	//cardsTwo.print()

	cards.shuffle()
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
