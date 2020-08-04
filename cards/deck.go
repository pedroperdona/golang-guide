package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (deck deck) print() {
	for index, card := range deck {
		fmt.Println(index, card)
	}
}

func deal(deck deck, handSize int) (deck, deck) {
	return deck[:handSize], deck[handSize:]
}

func (deck deck) toString() string {
	return strings.Join([]string(deck), ",")
}

func (deck deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(deck.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return strings.Split(string(bytes), ",")
}

func (deck deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	for index := range deck {
		newPosition := random.Intn(len(deck) - 1)

		deck[index], deck[newPosition] = deck[newPosition], deck[index]
	}
}
