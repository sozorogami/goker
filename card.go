package goker

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

func rankSliceToIntSlice(s []rank) []int {
	ints := make([]int, len(s))
	for i := range s {
		ints[i] = int(s[i])
	}
	return ints
}

type PlayingCard interface {
	Rank() rank
	Suit() suit
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

func NewCard(r rank, s suit) *card {
	c := card{r, s}
	return &c
}
