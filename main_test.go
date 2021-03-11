package main

import (
	"fmt"
	"testing"

	"github.com/oyugirachel/deck"
)

var (
	ace       = deck.Card{Suit: deck.Spade, Rank: deck.Ace}
	two       = deck.Card{Suit: deck.Spade, Rank: deck.Two}
	three     = deck.Card{Suit: deck.Spade, Rank: deck.Three}
	four      = deck.Card{Suit: deck.Spade, Rank: deck.Four}
	five      = deck.Card{Suit: deck.Spade, Rank: deck.Five}
	six       = deck.Card{Suit: deck.Spade, Rank: deck.Six}
	seven     = deck.Card{Suit: deck.Spade, Rank: deck.Seven}
	eight     = deck.Card{Suit: deck.Spade, Rank: deck.Eight}
	nine      = deck.Card{Suit: deck.Spade, Rank: deck.Nine}
	tenSpades = deck.Card{Suit: deck.Spade, Rank: deck.Ten}
	tenHearts = deck.Card{Suit: deck.Heart, Rank: deck.Ten}
	jack      = deck.Card{Suit: deck.Spade, Rank: deck.Jack}
	queen     = deck.Card{Suit: deck.Spade, Rank: deck.Queen}
	king      = deck.Card{Suit: deck.Spade, Rank: deck.King}
)

func Test_scoring(t *testing.T) {

	tests := []struct {
		name  string
		cards [2]deck.Card
		snap  bool
		score int
	}{
		{
			name: "differentCards-snap",

			cards: [2]deck.Card{six, tenSpades},
			snap:  true,
			score: -1,
		},
		{
			name:  "differentCards",
			cards: [2]deck.Card{six, tenSpades},
			snap:  false,
			score: 0,
		},

		{
			name:  "sameRank-SameSuits-snap",
			cards: [2]deck.Card{tenHearts, tenHearts},
			snap:  true,
			score: 1,
		},
		{
			name:  "sameRank-SameSuits",
			cards: [2]deck.Card{tenHearts, tenHearts},
			snap:  false,
			score: -1,
		},
		{
			name:  "sameRank-DiffSuits-snap",
			cards: [2]deck.Card{tenSpades, tenHearts},
			snap:  true,
			score: 1,
		},
		{
			name:  "sameRank-DiffSuits",
			cards: [2]deck.Card{tenSpades, tenHearts},
			snap:  false,
			score: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			presentCards = tt.cards
			score = 0

			assertEquals(t, scoring(tt.snap), tt.score)

		})
	}
}
func assertEquals(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func Test_drawCard(t *testing.T) {

	tests := []struct {
		name       string
		done       chan bool
		firstcards [2]deck.Card
		drawCard   []deck.Card
		expected   [2]deck.Card
	}{
		{
			name:       "Same",
			done:       make(chan bool),
			firstcards: [2]deck.Card{seven, six},
			drawCard:   []deck.Card{ace},
			expected:   [2]deck.Card{six, ace},
		},
		{
			name:       "different-presentCards",
			done:       make(chan bool),
			firstcards: [2]deck.Card{eight, nine},
			drawCard:   []deck.Card{tenHearts},
			expected:   [2]deck.Card{nine, tenHearts},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			lastCard = len(tt.drawCard)
			presentCards = tt.firstcards
			drawCard(tt.done, tt.drawCard)

			if presentCards == tt.expected {
				fmt.Println("tests are successful")

			} else {
				fmt.Println("tests failed")
			}

		})
	}

}
