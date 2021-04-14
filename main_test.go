package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/oyugirachel/deck"
)
// Declaring a variables of customized values t be used during testing
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
		cards []deck.Card
		snap  bool
		score int
	}{
		{
			name: "differentCards-snap",

			cards: []deck.Card{six, tenSpades},
			snap:  true,
			score: -1,
		},
		{
			name:  "differentCards",
			cards: []deck.Card{six, tenSpades},
			snap:  false,
			score: 0,
		},

		{
			name:  "sameRank-SameSuits-snap",
			cards: []deck.Card{tenHearts, tenHearts},
			snap:  true,
			score: 1,
		},
		{
			name:  "sameRank-SameSuits",
			cards: []deck.Card{tenHearts, tenHearts},
			snap:  false,
			score: -1,
		},
		{
			name:  "sameRank-DiffSuits-snap",
			cards: []deck.Card{tenSpades, tenHearts},
			snap:  true,
			score: 1,
		},
		{
			name:  "sameRank-DiffSuits",
			cards: []deck.Card{tenSpades, tenHearts},
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
		name              string
		done              chan bool
		drawnCardposition int

		firstcards       []deck.Card
		drawCard         []deck.Card
		shouldSignalDone bool
		expected         []deck.Card
	}{
		{
			name:              "1st Card",
			done:              make(chan bool),
			firstcards:        []deck.Card{seven, six},
			drawnCardposition: 0,
			drawCard:          []deck.Card{ace},

			shouldSignalDone: false,
			expected:         []deck.Card{six, ace},
		},
		{
			name:              "26th Card",
			done:              make(chan bool),
			firstcards:        []deck.Card{eight, ace},
			drawCard:          []deck.Card{ace},
			drawnCardposition: 25,
			shouldSignalDone:  false,
			expected:          []deck.Card{ace, ace},
		},
		{
			name:              "52nd Card",
			done:              make(chan bool),
			firstcards:        []deck.Card{nine, queen},
			drawCard:          []deck.Card{three},
			drawnCardposition: 52,
			shouldSignalDone:  true,
			expected:          []deck.Card{queen, three},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			lastCard = tt.drawnCardposition
			presentCards = tt.firstcards
			done := make(chan bool)
			u := drawCard(done, tt.drawCard)

			if reflect.DeepEqual(u, tt.expected) {
				fmt.Printf("Tests passed")

			} else {
				fmt.Printf("Tests failed")
			}

		})
	}

}
