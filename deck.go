package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

// create new type of 'deck'
// which is slice of strings

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Diamond", "Spades", "Hearts", "Clubs"}
	cardValues := []string{"One", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 06666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")

	return deck(s)
}

func (d deck) shuffle() deck {
	size_of_deck := d.size()
	for i := range d {
		newPosition := rand.Intn(size_of_deck - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
	return d
}

func (d deck) size() int {
	return len(d)
}
