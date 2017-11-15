package goker_test

import (
	"sort"

	. "github.com/sozorogami/goker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Hands", func() {
	Describe("Creating a new hand", func() {
		hand := NewHand(
			NewCard(Two, Diamond),
			NewCard(Four, Diamond),
			NewCard(Eight, Diamond),
			NewCard(Queen, Diamond),
			NewCard(Three, Diamond))

		It("sorts the cards low to high", func() {
			sort.Sort(hand)
			Expect(hand[4]).To(Equal(NewCard(Queen, Diamond)))
		})
	})

	Describe("Ranking hands", func() {
		Context("When the hand contains a royal straight flush", func() {
			hand := NewHand(
				NewCard(Ten, Club),
				NewCard(Jack, Club),
				NewCard(Queen, Club),
				NewCard(King, Club),
				NewCard(Ace, Club))
			It("is ranked a royal straight flush", func() {
				Expect(hand.Rank().Name()).To(Equal("RoyalStraightFlush"))
			})
		})

		Context("When the hand contains a straight flush", func() {
			hand := NewHand(
				NewCard(Ten, Club),
				NewCard(Jack, Club),
				NewCard(Queen, Club),
				NewCard(King, Club),
				NewCard(Nine, Club))
			It("is ranked a straight flush", func() {
				Expect(hand.Rank().Name()).To(Equal("StraightFlush"))
			})
		})

		Context("When the hand contains a four of a kind", func() {
			hand := NewHand(
				NewCard(Ten, Club),
				NewCard(Ten, Heart),
				NewCard(Ten, Spade),
				NewCard(Ten, Diamond),
				NewCard(Nine, Club))
			It("is ranked a four of a kind", func() {
				Expect(hand.Rank().Name()).To(Equal("FourOfAKind"))
			})
		})

		Context("When the hand contains a full house", func() {
			hand := NewHand(
				NewCard(Ten, Club),
				NewCard(Ten, Heart),
				NewCard(Ten, Spade),
				NewCard(Nine, Diamond),
				NewCard(Nine, Club))
			It("is ranked a full house", func() {
				Expect(hand.Rank().Name()).To(Equal("FullHouse"))
			})
		})

		Context("When the hand contains a flush", func() {
			hand := NewHand(
				NewCard(Two, Club),
				NewCard(Five, Club),
				NewCard(Seven, Club),
				NewCard(Queen, Club),
				NewCard(King, Club))
			It("is ranked a flush", func() {
				Expect(hand.Rank().Name()).To(Equal("Flush"))
			})
		})

		Context("When the hand contains a straight", func() {
			hand := NewHand(
				NewCard(Two, Diamond),
				NewCard(Three, Club),
				NewCard(Four, Club),
				NewCard(Five, Club),
				NewCard(Six, Club))
			It("is ranked a straight", func() {
				Expect(hand.Rank().Name()).To(Equal("Straight"))
			})

			// We should test this by comparing value
			// Context("and the staight starts with an ace", func() {
			// 	al := NewHand(
			// 		NewCard(Two, Diamond),
			// 		NewCard(Three, Club),
			// 		NewCard(Four, Club),
			// 		NewCard(Five, Club),
			// 		NewCard(Ace, Club))
			// 	It("has high card of 5", func() {
			// 		Expect(al.Rank()).To(Equal(NewStraight(Five)))
			// 	})
			// })
		})

		Context("When the hand contains a three of a kind", func() {
			hand := NewHand(
				NewCard(Four, Club),
				NewCard(Four, Heart),
				NewCard(Four, Spade),
				NewCard(Seven, Club),
				NewCard(Nine, Club))
			It("is ranked a three of a kind", func() {
				Expect(hand.Rank().Name()).To(Equal("ThreeOfAKind"))
			})
		})

		Context("When the hand contains two pair", func() {
			hand := NewHand(
				NewCard(Four, Club),
				NewCard(Four, Heart),
				NewCard(Seven, Spade),
				NewCard(Seven, Club),
				NewCard(Nine, Club))
			It("is ranked two pair", func() {
				Expect(hand.Rank().Name()).To(Equal("TwoPair"))
			})
		})

		Context("When the hand contains one pair", func() {
			hand := NewHand(
				NewCard(Four, Club),
				NewCard(Four, Heart),
				NewCard(Eight, Spade),
				NewCard(Seven, Club),
				NewCard(Nine, Club))
			It("is ranked one pair", func() {
				Expect(hand.Rank().Name()).To(Equal("Pair"))
			})
		})

		Context("When the hand is trash", func() {
			hand := NewHand(
				NewCard(Two, Club),
				NewCard(Four, Heart),
				NewCard(Eight, Spade),
				NewCard(Seven, Club),
				NewCard(Nine, Club))
			It("is ranked by high card", func() {
				Expect(hand.Rank().Name()).To(Equal("HighCard"))
			})
		})
	})
})
