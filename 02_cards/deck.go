package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//Create a new type of deck which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handsize int) (deck, deck) {
	//Return a deck from index 0 to handsize, non inclusive and from index handsize to the end of the array (slice)
	return d[:handsize], d[handsize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		//Option #1 - Log the error and return a call to newDeck()
		//Option #2 - Log the error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	cardString := string(bs)            //Ace of Spades,Two of Spades...
	s := strings.Split(cardString, ",") //return a slice of strings
	return deck(s)
}

func (d deck) shuffle() {
	//Use current time as seed to generate random number generator
	currentTime := time.Now().UnixNano()

	source := rand.NewSource(currentTime)
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
