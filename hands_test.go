package goker_test

import (
	. "github.com/sozorogami/goker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var royalStraightFlush = NewHand(
	NewCard(Ten, Club),
	NewCard(Jack, Club),
	NewCard(Queen, Club),
	NewCard(King, Club),
	NewCard(Ace, Club))

var straightFlush = NewHand(
	NewCard(Ten, Club),
	NewCard(Jack, Club),
	NewCard(Queen, Club),
	NewCard(King, Club),
	NewCard(Nine, Club))

var fourOfAKind = NewHand(
	NewCard(Ten, Club),
	NewCard(Ten, Heart),
	NewCard(Ten, Spade),
	NewCard(Ten, Diamond),
	NewCard(Nine, Club))

var fullHouse = NewHand(
	NewCard(Ten, Club),
	NewCard(Ten, Heart),
	NewCard(Ten, Spade),
	NewCard(Nine, Diamond),
	NewCard(Nine, Club))

var flush = NewHand(
	NewCard(Two, Club),
	NewCard(Five, Club),
	NewCard(Seven, Club),
	NewCard(Queen, Club),
	NewCard(King, Club))

var straight = NewHand(
	NewCard(Two, Diamond),
	NewCard(Three, Club),
	NewCard(Four, Club),
	NewCard(Five, Club),
	NewCard(Six, Club))

var aceLowStraight = NewHand(
	NewCard(Two, Diamond),
	NewCard(Three, Club),
	NewCard(Four, Club),
	NewCard(Five, Club),
	NewCard(Ace, Club))

var threeOfAKind = NewHand(
	NewCard(Four, Club),
	NewCard(Four, Heart),
	NewCard(Four, Spade),
	NewCard(Seven, Club),
	NewCard(Nine, Club))

var twoPair = NewHand(
	NewCard(Four, Club),
	NewCard(Four, Heart),
	NewCard(Seven, Spade),
	NewCard(Seven, Club),
	NewCard(Nine, Club))

var pair = NewHand(
	NewCard(Four, Club),
	NewCard(Four, Heart),
	NewCard(Eight, Spade),
	NewCard(Seven, Club),
	NewCard(Nine, Club))

var highCard = NewHand(
	NewCard(Two, Club),
	NewCard(Four, Heart),
	NewCard(Eight, Spade),
	NewCard(Seven, Club),
	NewCard(Nine, Club))

var _ = Describe("Hands", func() {
	Describe("Creating a new hand", func() {
		hand := NewHand(
			NewCard(Two, Diamond),
			NewCard(Four, Diamond),
			NewCard(Eight, Diamond),
			NewCard(Queen, Diamond),
			NewCard(Three, Diamond))

		It("sorts the cards low to high", func() {
			Expect(&hand.Cards[4]).To(Equal(NewCard(Queen, Diamond)))
		})
	})

	Describe("Identifying hands", func() {
		Context("When the hand consists of AKQJT of the same suit", func() {
			It("is a royal straight flush", func() {
				Expect(royalStraightFlush.Rank().Name()).To(Equal("RoyalStraightFlush"))
			})
		})

		Context("When the hand consists of five cards of the same suit with consecutive ranks", func() {
			It("is a straight flush", func() {
				Expect(straightFlush.Rank().Name()).To(Equal("StraightFlush"))
			})
		})

		Context("When the hand contains four cards of the same rank", func() {
			It("is a four of a kind", func() {
				Expect(fourOfAKind.Rank().Name()).To(Equal("FourOfAKind"))
			})
		})

		Context("When the hand contains three cards of one rank and two cards of another", func() {
			It("is a full house", func() {
				Expect(fullHouse.Rank().Name()).To(Equal("FullHouse"))
			})
		})

		Context("When the hand consists of five cards of the same suit with non-consecutive ranks", func() {
			It("is a flush", func() {
				Expect(flush.Rank().Name()).To(Equal("Flush"))
			})
		})

		Context("When the hand consists of five cards of consecutive rank of different suits", func() {
			It("is a straight", func() {
				Expect(straight.Rank().Name()).To(Equal("Straight"))
			})
		})

		Context("When the hand contains three cards of the same rank and no other pairs", func() {
			It("is a three of a kind", func() {
				Expect(threeOfAKind.Rank().Name()).To(Equal("ThreeOfAKind"))
			})
		})

		Context("When the hand contains two sets of two cards of the same rank", func() {
			It("is a two pair", func() {
				Expect(twoPair.Rank().Name()).To(Equal("TwoPair"))
			})
		})

		Context("When the hand contains only one set of two cards of the same rank", func() {
			It("is one pair", func() {
				Expect(pair.Rank().Name()).To(Equal("Pair"))
			})
		})

		Context("When none of the above are true", func() {
			It("is ranked by high card", func() {
				Expect(highCard.Rank().Name()).To(Equal("HighCard"))
			})
		})
	})
})
