package goker_test

import (
	. "github.com/sozorogami/goker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Making hands from sets of cards", func() {
	Context("when there aren't enough cards in the set", func() {
		It("can't make any hands", func() {
			cards := CardSet{NewCard(Ace, Diamond), NewCard(King, Spade)}
			Expect(cards.PossibleHands()).To(BeEmpty())
		})
	})
	Context("when there are five cards in the set", func() {
		cards := CardSet{
			NewCard(Ace, Diamond),
			NewCard(King, Spade),
			NewCard(Queen, Spade),
			NewCard(Jack, Heart),
			NewCard(Ten, Spade)}
		It("can only make one hand", func() {
			Expect(len(cards.PossibleHands())).To(Equal(1))
		})
		It("finds the best possible hand", func() {
			Expect(cards.BestPossibleHand().Rank().Name()).To(Equal("Straight"))
		})
	})
	Context("when there are seven cards in the set", func() {
		cards := CardSet{
			NewCard(Ace, Diamond),
			NewCard(King, Spade),
			NewCard(Queen, Spade),
			NewCard(Jack, Heart),
			NewCard(Jack, Spade),
			NewCard(Ace, Spade),
			NewCard(Ten, Spade)}
		It("can make 21 hands", func() {
			Expect(len(cards.PossibleHands())).To(Equal(21))
		})
		It("finds the best possible hand", func() {
			Expect(cards.BestPossibleHand().Rank().Name()).To(Equal("RoyalStraightFlush"))
		})
	})
})
