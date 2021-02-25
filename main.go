package main

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"

	"github.com/eiannone/keyboard"
	"github.com/oyugirachel/deck"

	"log"

	"time"
)

// Holds the last two cards that will be displayed to the user
var presentCards [2]deck.Card
var score = 0
var lastCard = 1

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
  ** 1 point is lost  if you SNAP when the value of the cards dont match **
  ** 1 point is lost if you don't SNAP and the value of cards match **
  ** Press esc key to exit the game **





BE ON THE LOOKOUT !


`
	fmt.Println(message)
	for k := 6; k > 0; k-- {
		fmt.Printf("%d ..", k)
		time.Sleep(time.Second)
	}
	fmt.Println("Gooooo!")

	presentCards = [2]deck.Card{cards[0], cards[1]}
	lastCard = 1
	// Showing the two initial  cards
	fmt.Printf("=============================[%2d/%2d]~ \n", lastCard+1, len(cards))
	fmt.Println(presentCards[0])
	fmt.Println(presentCards[1])
	fmt.Println("============================")
	fmt.Println("Your score is :", score)

	ticker := time.NewTicker(2 * time.Second)

	inputChannel := make(chan rune)

	done := make(chan bool)
	go func() {
		for {

			char, key, err := keyboard.GetSingleKey()
			if err != nil {
				log.Println(err)
			}

			if key == keyboard.KeyEsc {
				done <- true
			}
			inputChannel <- char

		}

	}()

	for {

		select {
		case <-inputChannel:

			fmt.Println("snap")

			scoring(true)

			drawCard(done, cards, ticker)

		case <-ticker.C:

			drawCard(done, cards, ticker)
			scoring(false)

		case <-done:
			fmt.Println("Game over! you scored a total of ", score)
			return

		}

	}

}

// drawCard function that gets the next card from the deck and adds it to the list of the present cards
func drawCard(done chan bool, cards []deck.Card, ticker *time.Ticker) {
	lastCard++

	if lastCard >= len(cards) {
		// incase a channel doesnot have a ready receiver, it doesnt block code execution
		go func() {
			done <- true
		}()
		return
	}
	presentCards[0] = presentCards[1]
	presentCards[1] = cards[lastCard]
	fmt.Printf("=============================[%2d/%2d]~ \n", lastCard+1, len(cards))
	fmt.Println(presentCards[0])
	fmt.Println(presentCards[1])
	fmt.Println("============================")
	ticker = time.NewTicker(2 * time.Second)
}

// scoring function
func scoring(snap bool) {

	if snap {
		if presentCards[0].Rank == presentCards[1].Rank {
			score++
			fmt.Println("\nYour score is:", score)
			return
		}
		score--
		fmt.Println("\nYour score is:", score)
		return
	}
	// This means they've not snapped
	if presentCards[0].Rank == presentCards[1].Rank {
		score--
		fmt.Println("\nYour score is:", score)

		return
	}
	fmt.Println("\nYour score is:", score)
	return
}
