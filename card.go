package goker

import "fmt"

type suit int8

const (
	// Spade represents the suit ♠
	Spade suit = iota
	// Heart represents the suit ♥
	Heart
	// Diamond represents the suit ♦
	Diamond
	// Club represents the suit ♣
	Club
)

func (s suit) String() string {
	switch s {
	case Spade:
		return "♠"
	case Heart:
		return "♥"
	case Diamond:
		return "♦"
	case Club:
		return "♣"
	default:
		return "?"
	}
}

type rank int8

const (
	// Two represents the rank 2
	Two rank = iota + 2
	// Three represents the rank 3
	Three
	// Four represents the rank 4
	Four
	// Five represents the rank 5
	Five
	// Six represents the rank 6
	Six
	// Seven represents the rank 7
	Seven
	// Eight represents the rank 8
	Eight
	// Nine represents the rank 9
	Nine
	// Ten represents the rank 10
	Ten
	// Jack represents the rank Jack
	Jack
	// Queen represents the rank Queen
	Queen
	// King represents the rank King
	King
	// Ace represents the rank Ace
	Ace
)

func (r rank) String() string {
	switch r {
	case Ace:
		return "A"
	case King:
		return "K"
	case Queen:
		return "Q"
	case Jack:
		return "J"
	case Ten:
		return "T"
	default:
		return fmt.Sprintf("%d", int(r))
	}
}

// Card represents a playing card, with a suit and rank
type Card struct {
	Rank rank
	Suit suit
}

func (c Card) String() string {
	return c.Rank.String() + c.Suit.String()
}

// NewCard constructs a new card of the given suit and rank
func NewCard(r rank, s suit) *Card {
	c := Card{r, s}
	return &c
}
