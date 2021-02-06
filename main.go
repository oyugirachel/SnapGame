package main

import (
	"fmt"
	"github.com/oyugirachel/deck"
	"math/rand"
	"time"
)

// Holds the last two cards that will be displayed to the user
var lastTwoCards [2]deck.Card

func main() {
	cards := deck.New(deck.Deck(1), deck.Shuffle)
	// fmt.Println(cards)
	// fmt.Println(len(cards))
	fmt.Println(drawRandomCard(cards))

}

// DrawRandomCard function
func drawRandomCard(cards []deck.Card) deck.Card {
	// Creating a random variable to draw one card
	rand.Seed(time.Now().UnixNano())
	// Generates a random card position between 0 and the length of the cards
	var cardPosition = rand.Intn(len(cards))
	fmt.Println(cardPosition)
	// Returning the random chosen card
	return cards[cardPosition]

}
func timedShuffle() {
	// creating our timer
	timer := time.NewTimer(2 * time.Second)

	for {
		select {
		// Waiting for the channel to emit a value
		case <-timer.C:
			// recursively call our shuffle
			go timedShuffle()

		}
	}
}

