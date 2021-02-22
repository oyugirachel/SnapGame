package main

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
	"github.com/oyugirachel/deck"
	"log"

	"time"
)

// Holds the last two cards that will be displayed to the user
var presentCards [2]deck.Card
var score = 0
var lastCard = 1
var input string

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

	presentCards = [2]deck.Card{cards[0], cards[1]}
	lastCard = 1
	// Showing the two initial  cards
	fmt.Printf("=============================[%2d/%2d]~ \n", lastCard+1, len(cards))
	fmt.Println(presentCards[0])
	fmt.Println(presentCards[1])
	fmt.Println("============================")
	fmt.Println("Your score is :", score)

	ticker := time.NewTicker(2 * time.Second)

	inputChannel := make(chan string)

	done := make(chan bool)
	go func() {

		var input string
		if _, err := fmt.Scanf("%s\n", &input); err != nil {

			log.Println(err)
		}
		inputChannel <- input
		return

	}()

	for {

		select {
		case input := <-inputChannel:

			if input != "" {
				fmt.Println("Snap")

			}
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
		if presentCards[0].Suit == presentCards[1].Suit {
			score++
			fmt.Println("\nYour score is:", score)
			return
		}
		score--
		fmt.Println("\nYour score is:", score)
		return
	}
	// this means they've not snapped
	if presentCards[0].Suit == presentCards[1].Suit {
		score--
		fmt.Println("\nYour score is:", score)

		return
	}
	fmt.Println("\nYour score is:", score)
	return
}
