package goker

import "fmt"

type suit int8

const (
	// Spade - Constant representing the suit ♠
	Spade suit = iota
	// Heart - Constant representing the suit ♥
	Heart
	// Diamond - Constant representing the suit ♦
	Diamond
	// Club - Constant representing the suit ♣
	Club
)

type rank int8

const (
	// Two - Constant representing the rank 2
	Two rank = iota + 2
	// Three - Constant representing the rank 3
	Three
	// Four - Constant representing the rank 4
	Four
	// Five - Constant representing the rank 5
	Five
	// Six - Constant representing the rank 6
	Six
	// Seven - Constant representing the rank 7
	Seven
	// Eight - Constant representing the rank 8
	Eight
	// Nine - Constant representing the rank 9
	Nine
	// Ten - Constant representing the rank 10
	Ten
	// Jack - Constant representing the rank Jack
	Jack
	// Queen - Constant representing the rank Queen
	Queen
	// King - Constant representing the rank King
	King
	// Ace - Constant representing the rank Ace
	Ace
)

type Card interface {
	Rank() rank
	Suit() suit
	String() string
}

type card struct {
	rank rank
	suit suit
}

func (c *card) Rank() rank {
	return c.rank
}

func (c *card) Suit() suit {
	return c.suit
}

func (c *card) String() string {
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

func NewCard(r rank, s suit) *card {
	c := card{r, s}
	return &c
}
