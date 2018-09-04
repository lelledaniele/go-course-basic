package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Type
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := [4]string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := [4]string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardValues {
		for _, value := range cardSuits {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}

// Receiver
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ", ")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(bs), ", "))
}

func (d deck) shuffle() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	dl := len(d)

	for i := range d {
		ri := r.Intn(dl - 1)

		d[i], d[ri] = d[ri], d[i]
	}
}
