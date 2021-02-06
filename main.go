package main

import (
	"fmt"
	"github.com/oyugirachel/deck"
	"math/rand"
	"time"
)

// Holds the last two cards that will be displayed to the user
var lastTwoCards [2]deck.Card
var score = 0

func main() {
	cards := deck.New(deck.Deck(1), deck.Shuffle)
	// calling in a goroutine to prevent blocking
	go timedShuffle(cards)
	// fmt.Println(cards)
	// fmt.Println(len(cards))
	// fmt.Println(drawRandomCard(cards))
	var input string

	for {
		fmt.Scanf("%s\n", &input)
		// fmt.Println(input)
		if input != "" {

			if lastTwoCards[0] == lastTwoCards[1] {
				// increment the score for the user has snapped
				score++
			} else {
				// If the user snaps and the cards are not the same
				score--
			}

		}
		// Check if the last two cards are the same
		if lastTwoCards[0] == lastTwoCards[1] {
			// We are sure the user hasnt snapped so we deduct the score
			score--
		}
		input = ""
		fmt.Println("Your score is:", score)

	}

}

// DrawRandomCard function
func drawRandomCard(cards []deck.Card) deck.Card {
	// Creating a random variable to draw one card
	rand.Seed(time.Now().UnixNano())
	// Generates a random card position between 0 and the length of the cards
	var cardPosition = rand.Intn(len(cards))
	// fmt.Println(cardPosition)
	// Returning the random chosen card
	return cards[cardPosition]

}
func timedShuffle(cards []deck.Card) {
	// creating our timer
	timer := time.NewTimer(2 * time.Second)

	for {
		select {
		// Waiting for the channel to emit a value
		case <-timer.C:
			// recursively call our shuffle
			go timedShuffle(cards)

			card := drawRandomCard(cards)
			// shift position to position one
			lastTwoCards[0] = lastTwoCards[1]
			// taken the random card to be the most recent one
			lastTwoCards[1] = card
			fmt.Println(lastTwoCards)

		}
	}
}

// scanInput function scans the input
func scanInput() {

}
