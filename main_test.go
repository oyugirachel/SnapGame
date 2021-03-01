package main

import (
	// "github.com/oyugirachel/deck"
	"testing"
)

func Test_scoring(t *testing.T) {
	type args struct {
		snap bool
	}
	tests := []struct {
		name         string
		presentCards string
		args         args
	}{
		{
			name:         "differentCards",
			presentCards: "Ace of Hearts, Eight of Spades",
			args:         args{snap: true},
		},
		{
			name:         "differentCards",
			presentCards: "Ace of Hearts, Eight of Spades",
			args:         args{snap: false},
		},

		{
			name:         "sameCards",
			presentCards: "King of Spades, King of Diamonds",
			args:         args{snap: true},
		},
		{
			name:         "sameCards",
			presentCards:"King of Spades, King of Diamonds",
			args:         args{snap: false},
		},
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scoring(tt.args.snap)
		})
	}
}
