package main

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
	"github.com/oyugirachel/deck"

	"math/rand"
	"time"
)

// Holds the last two cards that will be displayed to the user

var presentCards [2]deck.Card
var lastCard = 1
var score = 0
var cardsDrawn = 0

func checkLastTwoCards(snap bool) {
	if snap {
		if presentCards[0] == presentCards[1] {
			// increment the score for the user has snapped
			score++
		} else {
			// If the user snaps and the cards are not the same
			score--
		}
	} else {
		// Check if the last two cards are the same
		if presentCards[0] == presentCards[1] {
			// We are sure the user hasnt snapped so we deduct the score
			score--
		}

	}
	fmt.Println("\nYour score is:", score)

}

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
	// go timedShuffle(cards)
	presentCards := []deck.Card{cards[0], cards[1]}
	lastCard := 1

	ticker := time.NewTicker(2 * time.Second)
	inputChannel := make(chan string)
	done := make(chan bool)

	for {
		go func() {
			var input string
			fmt.Scanf("%s", &input)
			inputChannel <- input

		}()

		select {
		case <-ticker.C:
			fmt.Printf("=============================[%2d/%2d]~ \n", lastCard, len(cards))
			fmt.Println(presentCards[0])
			fmt.Println(presentCards[1])
			fmt.Println("============================")

			lastCard++
			if lastCard >= len(cards) {
				done <- true
				return
			}
			presentCards[0] = presentCards[1]
			presentCards[1] = cards[lastCard]
		case input := <-inputChannel:
			if input == "" {
				fmt.Println("SNAP")
				checkLastTwoCards(true)

			}
		case <-done:
			fmt.Println("Game over! your scored a total of ", score)
			return
		}
	}

}
