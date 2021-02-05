package main

import (
	"fmt"
	"github.com/oyugirachel/deck"
	"strings"
	"time"
)

// Hand type
type Hand []deck.Card

func (h Hand) String() string {
	strs := make([]string, len(h))
	for i := range h {
		strs[i] = h[i].String()
	}
	return strings.Join(strs, ",")
}

// DealerString function
func (h Hand) DealerString() string {
	return h[0].String() + ",**HIDDEN**"

}

// Score Function
func (h Hand) Score() int {
	minScore := h.MinScore()
	if minScore > 11 {
		return minScore
	}
	for _, c := range h {
		if c.Rank == deck.Ace {
			// ace is currently worth 1, and we are changing it to be worth 11
			//11-1 =10
			return minScore + 10
		}

	}
	return minScore

}

// MinScore function
func (h Hand) MinScore() int {
	score := 0
	for _, c := range h {
		score += min(int(c.Rank), 10)
	}
	return score

}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Shuffle function implements the shuffling of decks
func Shuffle(gs GameState) GameState {
	ret := clone(gs)
	ret.Deck = deck.New(deck.Deck(1), deck.Shuffle)
	return ret
}

// Deal function
func Deal(gs GameState) GameState {
	ret := clone(gs)
	ret.Player = make(Hand, 0, 5)
	ret.Dealer = make(Hand, 0, 5)
	var card deck.Card
	// iterate to give each player two cards
	for i := 0; i < 2; i++ {
		card, ret.Deck = draw(ret.Deck)
		ret.Player = append(ret.Player, card)
		card, ret.Deck = draw(ret.Deck)
		ret.Dealer = append(ret.Dealer, card)

	}
	ret.State = StatePlayerTurn
	return ret
}

// Stand function
func Stand(gs GameState) GameState {
	ret := clone(gs)
	ret.State++
	return ret
}

// Snap function
func Snap(gs GameState) GameState {
	ret := clone(gs)
	hand := ret.CurrentPlayer()
	var card deck.Card
	card, ret.Deck = draw(ret.Deck)
	*hand = append(*hand, card)
	if hand.Score() > 21 {
		Stand(ret)
	}

	return ret
}

// EndHand function
func EndHand(gs GameState) GameState {
	ret := clone(gs)
	pScore, dScore := ret.Player.Score(), ret.Dealer.Score()
	// Printing out the final scores
	fmt.Println("**FINAL HANDS**")
	fmt.Println("Player", ret.Player, "\nScore:", pScore)
	fmt.Println("Dealer", ret.Dealer, "\n Score", dScore)
	switch {
	case pScore > 21:
		fmt.Println("You busted")
	case dScore > 21:
		fmt.Println("Dealer busted")
	case pScore > dScore:
		fmt.Println("You win!")
	case dScore > pScore:
		fmt.Println("You lose")
	case dScore == pScore:

		fmt.Println("Draw")

	}
	fmt.Println()
	ret.Player = nil
	ret.Dealer = nil
	return ret

}
func main() {
	// Running the timer in a goroutine so that it is non-blocking
	go timedShuffle()
	var gs GameState
	gs = Shuffle(gs)
	// for loop to iterate over after a set number of hands/Snap
	for i := 0; i < 4; i++ {
		gs = Deal(gs)
		var input string
		// s representing standing
		for gs.State == StatePlayerTurn {
			fmt.Println("Player:", gs.Player)
			fmt.Println("Dealer:", gs.Dealer.DealerString())
			fmt.Println("What will you do? (sn)ap,(s)tand")
			// Reading the users input
			fmt.Scanf("%s\n", &input)

			switch input {
			case "sn":
				gs = Snap(gs)
			case "s":
				gs = Stand(gs)
			default:
				fmt.Println("Invalid Option", input)
			}

		}

		for gs.State == StateDealerTurn {
			if gs.Dealer.Score() <= 16 || (gs.Dealer.Score() == 17 && gs.Dealer.MinScore() != 17) {
				gs = Snap(gs)

			} else {
				gs = Stand(gs)
			}

		}
		// if dealer score <= 16 we hit
		// if dealer has a soft 17, then we hit
		gs = EndHand(gs)

	}
	time.Sleep(2 * time.Second)

}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]

}

func timedShuffle() {
	// creating our timer
	timer := time.NewTimer(2 * time.Second)

	for {
		select {
		// Waiting for the channel to emit a value
		case <-timer.C:
			// recursively call our shuffle
			go timedShuffle()
			var gs GameState
			gs = Shuffle(gs)

		}
	}
}

// State type
type State int8

// Declaring constants
const (
	StatePlayerTurn State = iota
	StateDealerTurn
	StateHandOver
)

// GameState Struct
type GameState struct {
	Deck   []deck.Card
	State  State
	Player Hand
	Dealer Hand
}

// CurrentPlayer function
func (gs *GameState) CurrentPlayer() *Hand {
	switch gs.State {
	case StatePlayerTurn:
		return &gs.Player
	case StateDealerTurn:
		return &gs.Dealer
	default:
		panic("It isn't currently any players turn")

	}
}
func clone(gs GameState) GameState {
	ret := GameState{
		Deck:   make([]deck.Card, len(gs.Deck)),
		State:  gs.State,
		Player: make(Hand, len(gs.Player)),
		Dealer: make(Hand, len(gs.Dealer)),
	}
	copy(ret.Deck, gs.Deck)
	copy(ret.Player, gs.Dealer)
	copy(ret.Dealer, gs.Dealer)
	return ret

}
