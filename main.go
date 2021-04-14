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
var presentCards []deck.Card
//Score holds the current score
var score = 0
// lastCard variable which simply is an integer which counts each time a card is removed from the deck inclusive of the position
var lastCard = 1

func main() {
	// Shuffling one deck of cards
	cards := deck.New(deck.Deck(1), deck.Shuffle)
	// Game instructions
	// Creating Newfigure which takes in text, font, strict mode(either tre or false) and a"color" constructor takes a color as an additional arg
	// An empty string is passed for the font name, which provides a default font
	// Strict mode dictates how to handle characters outside of standard ASCII. When set to true, a non-ASCII character (outside character codes 32-127) will cause the program to panic. When set to false, these characters are replaced with a question mark ('?')
	art := figure.NewColorFigure("SNAP GAME", "", "Red", true)
	// A figure responds to the func Blink, taking three arguments. duration is the total time the banner will display, in milliseconds. timeOn is the length of time the text will blink on (also in ms). timeOff is the length of time the text will blink off (ms). For an even blink, set timeOff to -1 (same as setting timeOff to the value of timeOn). There is no return value.
	art.Blink(3000, 500, -1)

	art.Print()
	
	// Message definition

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
    // printing the message
	fmt.Println(message)
	// represents a counter 0-6
	for k := 6; k > 0; k-- {
		fmt.Printf("%d ..", k)
		time.Sleep(time.Second)
	}
	fmt.Println("Gooooo!")

	// assigning presentCards to the proper index
	presentCards = []deck.Card{cards[0], cards[1]}
	lastCard = 1
	// displaying the two initial  cards
	fmt.Printf("=============================[%2d/%2d]~ \n", lastCard+1, len(cards))
	fmt.Println(presentCards[0])
	fmt.Println(presentCards[1])
	fmt.Println("============================")
	fmt.Println("Your score is :", score)

	// Creating a time ticker
	ticker := time.NewTicker(2 * time.Second)
	// Creating an inputChannel that receive an input ones there is a key pressed
	inputChannel := make(chan rune)
	// Creating a done channel
	done := make(chan bool)
	// a goroutine implementing(listening) the key press
	go func() {
		for {

            // getting a single keystroke
			char, key, err := keyboard.GetSingleKey()
			// Checking for any error
			if err != nil {
				log.Println(err)
			}
            // A condition that calls the done channel when the escape key is called
			if key == keyboard.KeyEsc {
				done <- true
			}
			// Sending info to the input channel
			inputChannel <- char

		}

	}()
	// Calling the goroutine function that calls each channel according to the predetermined condition
	Goroutine(done, inputChannel, ticker, cards)

}

// Goroutine channels which instatiates the done channel, input channel and the ticker channel
func Goroutine(done chan bool, inputChannel chan rune, ticker *time.Ticker, cards []deck.Card) {
    // starting an endless loop of executing the goroutine when snap is false
	for {
		// Anotating snap to false
		snap := false
		// a select case which blocks an unblocks a channel depending on the one which case is free
		select {
			// When the goroutine runs and the card is the last one at the deck, a done message will be printed to the screen
		case <-done:
			fmt.Println("Game over! you scored a total of ", score)
			return
		// When the goroutine runs and there is a keypress, it will print "snap "
		// To the screen and annotate snap to true	
		case <-inputChannel:
			fmt.Println("snap")
			snap = true
		// Each time the goroutine runs the ticker goes off and the case ticker is run
		case <-ticker.C:
		}
		// updating the score and returning the change in the score
		points := scoring(snap)
		score += points
		// Printing the score to the screen
		fmt.Println("Your score is ", score)

        // Calling the draw function
		drawCard(done, cards)

		fmt.Printf("=============================[%2d/%2d]~ \n", lastCard+1, len(cards))
		fmt.Println(presentCards[0])
		fmt.Println(presentCards[1])
		fmt.Println("============================")
	}

}

// drawCard function that gets the next card from the deck and adds it to the list of the present cards
func drawCard(done chan bool, cards []deck.Card) []deck.Card {
	// incrementing the lastcard
	lastCard++
	// a condition implementing when the last card is at the end of the deck, then the game exits
	if lastCard >= len(cards) {
		// incase a channel doesnot have a ready receiver, it doesnt block code execution
		go func() {
			done <- true
		}()
		return presentCards

	}

	// Re-asigning of the cards
	presentCards[0] = presentCards[1]
	presentCards[1] = cards[lastCard]

	return presentCards

}

// scoring function compares the drawn two cards and if snapped, returns the change in the score
func scoring(snap bool) int {
	// if the player has snapped
	if snap {
		if presentCards[0].Rank == presentCards[1].Rank {

			return +1

		}

		return -1

	}
	// This means they've not snapped
	if presentCards[0].Rank == presentCards[1].Rank {

		return -1

	}

	return 0

}
