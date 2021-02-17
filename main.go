package main

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
	"github.com/oyugirachel/deck"

	"math/rand"
	"time"
)

// Holds the last two cards that will be displayed to the user
var lastTwoCards [2]deck.Card
var score = 0
var cardsDrawn = 0
var lastCard = 1

func main() {
	cards := deck.New(deck.Deck(1), deck.Shuffle)
	// Game instructions
	art := figure.NewColorFigure("SNAP GAME", "", "Red", true)
	art.Blink(3000, 500, -1)

	art.Print()

	message :=
		`
Press any key and enter to say SNAP when the value of the last two cards displayed on the screen matches

               =====BONUS=====
  ** 1 point is gained if you SNAP correctly **
  ** 1 point is lost  if you SNAP when the cards dont match **
  ** 1 point is lost if you don't SNAP and the cards match **





BE ON THE LOOKOUT !


`
	fmt.Println(message)
	for k := 6; k > 0; k-- {
		fmt.Printf("%d ..", k)
		time.Sleep(time.Second)
	}
	fmt.Println("Gooooo!")

	// Creating a random variable to draw one card

	rand.Seed(time.Now().UTC().UnixNano())

	// calling in a goroutine to prevent blocking
	go timedShuffle(cards)

	var input string

	for {

		fmt.Scanf("%s\n", &input)
		// fmt.Println(input)
		if input != "" {
			fmt.Println("SNAP")
			checkLastTwoCards(true)

		}

		input = ""
		fmt.Println("Players final score is :", score)

	}

}

func checkLastTwoCards(snap bool) {
	if snap {
		if lastTwoCards[0] == lastTwoCards[1] {
			// increment the score for the user has snapped
			score++
		} else {
			// If the user snaps and the cards are not the same
			score--
		}
	} else {
		// Check if the last two cards are the same
		if lastTwoCards[0] == lastTwoCards[1] {

			score++
		}

	}
	fmt.Println("\nYour score is:", score)

}

// randInt function to randomize time between max and min
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

// timedShuffle function
func timedShuffle(cards []deck.Card) {

	// t := randInt(1, 5)
	// x := time.Duration(t)

	// // creating our timer and randomizing it

	// timer := time.NewTimer(x * time.Second)

	timer := time.NewTicker(time.Second * 2)
	lastTwoCards := []deck.Card{cards[0], cards[1]}
	// done := make(chan bool)

	for {
		select {
		// Waiting for the channel to emit a value
		case <-timer.C:

			fmt.Printf("=============================[%2d/%2d]~ \n", lastCard, len(cards))
			fmt.Println(lastTwoCards[0])
			fmt.Println(lastTwoCards[1])
			fmt.Println("============================")

			lastCard++

			if lastCard >= len(cards) {
				return
			}

			lastTwoCards[0] = lastTwoCards[1]

			lastTwoCards[1] = cards[lastCard]

			checkLastTwoCards(false)

		}

	}
}

// scanInput function scans the input
func scanInput() {

}
