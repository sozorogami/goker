package goker

import (
	"math/rand"
)

// CardSet - A slice of an arbitrary number of cards
type CardSet []Card

// Deck - Represents a standard 52 card deck
type Deck struct {
	cards CardSet
}

// NewDeck - Constructs a CardSet representing a standard
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

// Len - The number of cards remaining in the deck
func (d *Deck) Len() int {
	return len(d.cards)
}

// Shuffle - Reorders the cards in the deck randomly, in place
func (d *Deck) Shuffle() {
	for i := len(d.cards) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
}

// Draw - Removes and returns the top n cards from the deck
func (d *Deck) Draw(n int) CardSet {
	cards := d.cards[0:n]
	d.cards = d.cards[n:]
	return cards
}
