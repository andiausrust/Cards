package main

func main () {

	//var card string = "Ace of Spades"
	//card := "Ace of spades"

	//cards := newDeckFromFile("my_cards")
	//cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()


}