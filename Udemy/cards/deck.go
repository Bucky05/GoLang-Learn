package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// this means deck type has all the behaviour which []string has
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardValues {
		for _, value := range cardSuits {
			cards = append(cards, suit+" of "+value)
		}
	}
	return cards
}
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

// reason we are not adding receiver is we are returning two decks
// which are pointer of the same decks, you can use receiver but this is the
// reason that fits more currently
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {

	return strings.Join([]string(d), ", ")

}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	byteSlice, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	cardString := string(byteSlice)
	cardStringSlice := strings.Split(cardString, ", ")

	return deck(cardStringSlice)
}

func (d deck) shuffle() {

	// random uses a seed so to get different random number on different
	// run use different seed value each time
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
