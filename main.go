package main

import (
	"fmt"

	"github.com/oyugirachel/deck"

	"math/rand"
	"time"

	"github.com/common-nighthawk/go-figure"
)

// Holds the last two cards that will be displayed to the user
var lastTwoCards [2]deck.Card
var score = 0
var cardsDrawn = 0

func main() {
	cards := deck.New(deck.Deck(1), deck.Shuffle)
	// Game instructions
	art := figure.NewColorFigure("SNAP GAME", "", "Red", true)
	art.Blink(3000, 500, -1)

	art.Print()

	message :=
		`
Press any key to say SNAP when the value of the last two cards displayed on the screen matches

               =====BONUS=====
  ** 1 point is gained if you SNAP correctly **
  ** 1 point is lost  if you SNAP when the cards dont match **
  ** 1 point is lost if you don't SNAP and the cards match **





BE ON THE LOOKOUT !


`
	fmt.Println(message)
	for k:=6; k>0; k--{
		fmt.Printf("%d ..",k)
		time.Sleep(time.Second)
	}
	fmt.Println("Gooooo!")

	// Creating a random variable to draw one card

	rand.Seed(time.Now().UTC().UnixNano())
	// calling in a goroutine to prevent blocking
	go timedShuffle(cards)

	// fmt.Println(cards)
	// fmt.Println(len(cards))
	// fmt.Println(drawRandomCard(cards))
	var input string

	for {
		if cardsDrawn == 52 {
			break
		}
		fmt.Scanf("%s\n", &input)
		// fmt.Println(input)
		if input != "" {
			fmt.Println("SNAP")
			checkLastTwoCards(true)

		}

		input = ""

	}
	fmt.Println("Players final score is :", score)

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
			// We are sure the user hasnt snapped so we deduct the score
			score--
		}

	}
	fmt.Println("\nYour score is:", score)

}

// DrawRandomCard function
func drawRandomCard(cards []deck.Card) deck.Card {

	// Generates a random card position between 0 and the length of the cards
	var cardPosition = rand.Intn(len(cards))
	// increment cards drawn
	cardsDrawn++
	// fmt.Println(cardPosition)
	// Returning the random chosen card

	return cards[cardPosition]

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

	for {
		select {
		// Waiting for the channel to emit a value
		case <-timer.C:
			// recursively call our shuffle

			// go timedShuffle(cards)

			card := drawRandomCard(cards)
			// shift position to position one
			lastTwoCards[0] = lastTwoCards[1]
			// taken the random card to be the most recent one
			lastTwoCards[1] = card
			checkLastTwoCards(false)

			for index, j := range lastTwoCards {
				if index == 0 { //If the value is first one
					fmt.Printf("[ '%v', ", j)
				} else if len(lastTwoCards) == index+1 { // If the value is the last one
					fmt.Printf("'%v' ]", j)
				} else {
					fmt.Printf(" '%v', ", j) // for all ( middle ) values
				}
			}

		}
	}
}

// scanInput function scans the input
func scanInput() {

}
