package goker

import (
	"math/rand"
)

type CardSet []Card

type deck struct {
	cards CardSet
}

func NewDeck() *deck {
	cards := CardSet{}
	for s := Spade; s <= Club; s++ {
		for r := Two; r <= Ace; r++ {
			cards = append(cards, NewCard(r, s))
		}
	}
	d := deck{cards}
	d.Shuffle()
	return &d
}

func (d *deck) Len() int {
	return len(d.cards)
}

func (d *deck) Shuffle() {
	for i := len(d.cards) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
}

func (d *deck) Draw(n int) CardSet {
	cards := d.cards[0:n]
	d.cards = d.cards[n:]
	return cards
}
