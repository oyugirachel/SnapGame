package main

import (
	"fmt"
	"github.com/oyugirachel/deck"
)

func main() {
	cards := deck.New(deck.Deck(1), deck.Shuffle)
	fmt.Println(cards)
}
