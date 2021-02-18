package main

import (
	"fmt"
	"log"
	"time"

	"github.com/oyugirachel/deck"
)

/**
Snap Game Mechanics
1. A deck of cards is shuffled
2. Game draws a card
3. Game displays last 2 drawn cards
4. Game listens for user input used for scoring
5. Game scores user based on input? and game output
6. Game ticks for 2 seconds
7. Loop until all 52 cards are displayed

*/

type SnapGame struct {
	CardDeck      []deck.Card // an array of 52 cards make a deck
	LastDrawnCard *deck.Card  // a holder for the last drawn card
	Score         int         // a value with the total score of the player
}

// Tick/wait for 2 seconds
func (snap *SnapGame) Tick() {
	// tick for 2 seconds
	timer := time.NewTicker(time.Second * 2)
	<-timer.C // block until we get response from ticker
}

func (snap *SnapGame) CardDrawer() deck.Card {
	// 	draw a card from the top of the stack,
	drawnCard := snap.CardDeck[0]

	// 	adjust/shift cards on deck accounting for 1 removed card
	snap.CardDeck = snap.CardDeck[1:]

	return drawnCard
}

// PrintCards is a function that formats the output of the displayed cards
func (snap *SnapGame) PrintCards(previousCard, currentCard *deck.Card) {
	fmt.Println("=====================================")
	fmt.Printf("[ '%v', '%v' ]\n", previousCard, currentCard)
	fmt.Println("=====================================")
	fmt.Println()
}

func (snap *SnapGame) Scorer(drawnCard deck.Card) {
	// run in a separate goroutine
	go func() {
		var input string
		if _, err := fmt.Scanf("%s\n", &input); err != nil {
			log.Println(err)
			return
		}

		if input == "" && &drawnCard == snap.LastDrawnCard {
			// scoring rule if input is nil and last two cards match, deduct a point
			snap.Score -= 1
		} else if input != "" && &drawnCard == snap.LastDrawnCard {
			// scoring rule if input is given and last two cards match, add a point
			snap.Score += 1
		} else if input != "" && &drawnCard != snap.LastDrawnCard {
			// scoring rule if input is given and last two cards dont match, deduct a point
			snap.Score -= 1
		}
	}()
}

func (snap *SnapGame) Run() {
	for len(snap.CardDeck) > 0 {
		// draw a card
		drawnCard := snap.CardDrawer()

		// display drawn card and the previously drawn card
		snap.PrintCards(snap.LastDrawnCard, &drawnCard)

		// score the player after drawing card
		snap.Scorer(drawnCard)

		// tick for 2 seconds
		snap.Tick()

		log.Printf("Deck cards: %v: Score: %v", len(snap.CardDeck), snap.Score)

		// update last drawn card
		snap.LastDrawnCard = &drawnCard
	}
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	cards := deck.New(deck.Deck(1), deck.Shuffle)
	snapGame := &SnapGame{CardDeck: cards[:52]}

	// Game instructions
	// art := figure.NewColorFigure("SNAP GAME", "", "Red", true)
	// art.Blink(3000, 500, -1)
	//
	// art.Print()
	// 	message :=
	// 		`
	// Press any key and enter to say SNAP when the value of the last two cards displayed on the screen matches
	//
	//                =====BONUS=====
	//   ** 1 point is gained if you SNAP correctly **
	//   ** 1 point is lost  if you SNAP when the cards dont match **
	//   ** 1 point is lost if you don't SNAP and the cards match **
	//
	// BE ON THE LOOKOUT !
	//
	// `

	snapGame.Run()

}

