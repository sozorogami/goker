package goker

import (
	"math/rand"
)

// CardSet is a slice of an arbitrary number of cards
type CardSet []*Card

// Deck represents a standard 52 card deck
type Deck struct {
	cards CardSet
}

// NewDeck constructs a CardSet representing a standard
// 52 card deck
func NewDeck() *Deck {
	cards := CardSet{}
	for s := Spade; s <= Club; s++ {
		for r := Two; r <= Ace; r++ {
			cards = append(cards, NewCard(r, s))
		}
	}
	d := Deck{cards}
	d.Shuffle()
	return &d
}

func NewStackedDeck(cards CardSet) *Deck {
	d := Deck{cards}
	return &d
}

// Len returns the number of cards remaining in the deck
func (d Deck) Len() int {
	return len(d.cards)
}

// Shuffle reorders the cards in the deck randomly, in place
func (d *Deck) Shuffle() {
	for i := len(d.cards) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
}

// Draw removes and returns the top n cards from the deck
func (d *Deck) Draw(n int) CardSet {
	splitIndex := d.Len() - n
	cards := d.cards[splitIndex:]
	d.cards = d.cards[0:splitIndex]
	return cards
}
