package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

// create a new type of deck
// which is a slice of strings

type deck[]string

func (d deck) print () {
	for i, card:= range d {
		fmt.Println(i, card)
	}
}

func newDeck () deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs "}
	cardValues := []string{"Ace", "Two", "Three"}

	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardSuit + " of " + cardValue)
		}
	}
	return cards
}

func deal (d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString () string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {

	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)

}

func newDeckFromFile(filename string) deck {

	bs, err  := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle ( ) {
	for i := range d {
		newPosition := rand.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}