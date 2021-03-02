package main

import (
	"github.com/oyugirachel/deck"
	"testing"
	
)

var (
	ace   = deck.Card{Suit: deck.Spade, Rank: deck.Ace}
	two   = deck.Card{Suit: deck.Spade, Rank: deck.Two}
	three = deck.Card{Suit: deck.Spade, Rank: deck.Three}
	four  = deck.Card{Suit: deck.Spade, Rank: deck.Four}
	five  = deck.Card{Suit: deck.Spade, Rank: deck.Five}
	six   = deck.Card{Suit: deck.Spade, Rank: deck.Six}
	seven = deck.Card{Suit: deck.Spade, Rank: deck.Seven}
	eight = deck.Card{Suit: deck.Spade, Rank: deck.Eight}
	nine  = deck.Card{Suit: deck.Spade, Rank: deck.Nine}
	ten   = deck.Card{Suit: deck.Spade, Rank: deck.Ten}
	jack  = deck.Card{Suit: deck.Spade, Rank: deck.Jack}
	queen = deck.Card{Suit: deck.Spade, Rank: deck.Queen}
	king  = deck.Card{Suit: deck.Spade, Rank: deck.King}
)

func Test_scoring(t *testing.T) {

	type args struct {
		snap bool
	}
	tests := []struct {
		name  string
		cards [2]deck.Card
		args  args
		score int
	}{
		{
			name: "differentCards",

			cards: [2]deck.Card{six, ten},
			args:  args{snap: true},
			score : -1,
		},
		{
			name:  "differentCards",
			cards: [2]deck.Card{six, ten},
			args:  args{snap: false},
			score: 0,
			
		},

		{
			name:  "sameCards",
			cards: [2]deck.Card{ace, ace},
			args:  args{snap: true},
			score: 1,
		},
		{
			name:  "sameCards",
			cards: [2]deck.Card{ten, ten},
			args:  args{snap: false},
			score: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cards := deck.New(deck.Deck(1), deck.Shuffle)
			lastCard = 1
			presentCards[0] = presentCards[1]
			presentCards[1] = cards[lastCard]
			presentCards = tt.cards
			
			assertEquals(t,scoring(tt.args.snap),tt.score)
			// scoring(tt.args.snap)
		})
	}
}
func assertEquals(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
