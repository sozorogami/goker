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

// Card represents a playing card, with a suit and rank
type Card interface {
	Rank() rank
	Suit() suit
	String() string
}

type card struct {
	rank rank
	suit suit
}

func (c card) Rank() rank {
	return c.rank
}

func (c card) Suit() suit {
	return c.suit
}

func (c card) String() string {
	var suitStr string
	switch c.Suit() {
	case Spade:
		suitStr = "♠"
	case Heart:
		suitStr = "♥"
	case Diamond:
		suitStr = "♦"
	case Club:
		suitStr = "♣"
	}

	var rankStr string
	switch c.Rank() {
	case Ace:
		rankStr = "A"
	case King:
		rankStr = "K"
	case Queen:
		rankStr = "Q"
	case Jack:
		rankStr = "J"
	case Ten:
		rankStr = "T"
	default:
		rankStr = fmt.Sprintf("%d", int(c.Rank()))
	}

	return rankStr + suitStr
}

// NewCard constructs a new card of the given suit and rank
func NewCard(r rank, s suit) Card {
	c := card{r, s}
	return &c
}
