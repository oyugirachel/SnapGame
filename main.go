package main

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
	"github.com/oyugirachel/deck"
	"context"
	"log"

	"time"
)

// Holds the last two cards that will be displayed to the user
var lastTwoCards [2]deck.Card
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

	presentCards := []deck.Card{cards[0], cards[1]}
	lastCard := 1

	ticker := time.NewTicker(2 * time.Second)
	inputChannel := make(chan string)
	ctx := context.Background()

	go func() {
		for ctx.Err() != nil {
			var input string
			if _, err := fmt.Scanf("%s\n", &input); err != nil {
				log.Println(err)
			}
			inputChannel <- input
			return
		}
	}()

	for {
		select {
		case <-ticker.C:
			fmt.Printf("=============================[%2d/%2d]~ \n", lastCard+1, len(cards))
			fmt.Println(presentCards[0])
			fmt.Println(presentCards[1])
			fmt.Println("============================")
			lastCard++
			if lastCard >= len(cards) {
				ctx.Done()
				return

			}
			presentCards[0] = presentCards[1]
			presentCards[1] = cards[lastCard]
			Scoring(false)

		case input := <-inputChannel:

			if input != "" {
				fmt.Println("Snap")

			}
			Scoring(true)

		case <-ctx.Done():
			fmt.Println("Game over! you scored a total of ", score)
			return

		}

	}

	
}
// Scoring function
func Scoring(Snap bool) {
	score = 0
	
	if (input == "") && (lastTwoCards[0] != lastTwoCards[1]) {
		score = 0
	} else if (input == "") && (lastTwoCards[0] == lastTwoCards[1]) {
		score--
	}
	if Snap {
		if (input == "") && (lastTwoCards[0] == lastTwoCards[1]) {
			score++
		} else if (input != "") && (lastTwoCards[0] != lastTwoCards[1]) {
			score--

		}

	}
	fmt.Println("\nYour score is:", score)

}





