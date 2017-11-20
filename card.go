package goker

import "fmt"

type suit int8
type rank int8

const (
	Spade suit = iota
	Heart
	Diamond
	Club
)

const (
	Two rank = iota + 2
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
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
