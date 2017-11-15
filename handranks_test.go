package goker_test

import (
	"sort"

	. "github.com/sozorogami/goker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Comparing hands", func() {
	Context("when they are of different ranks", func() {
		allRanks := HandGroup{
			threeOfAKind,
			twoPair,
			pair,
			highCard,
			royalStraightFlush,
			straightFlush,
			fourOfAKind,
			fullHouse,
			flush,
			straight}

		It("orders them correctly", func() {
			correctOrder := HandGroup{
				highCard,
				pair,
				twoPair,
				threeOfAKind,
				straight,
				flush,
				fullHouse,
				fourOfAKind,
				straightFlush,
				royalStraightFlush,
			}

			sort.Sort(allRanks)
			Expect(allRanks).To(Equal(correctOrder))
		})
	})

	Describe("Royal Straight Flush", func() {
		It("has a value of 9", func() {
			rsf := RoyalStraightFlush{}
			Expect(rsf.Value()[0]).To(Equal(9))
		})
	})

	Describe("Straight Flush", func() {
		It("has a value of 8.x, where x is the high card's rank", func() {
			sf := NewStraightFlush(NewCard(Jack, Spade).Rank())
			Expect(sf.Value()[0]).To(Equal(8))
			Expect(sf.Value()[1]).To(Equal(11))
		})
	})

	Describe("Four of a Kind", func() {
		It("has a value of 7.xy, where x is the rank of the four cards and y is the rank of the kicker", func() {
			foak := NewFourOfAKind(NewCard(Jack, Spade).Rank(), NewCard(Ace, Heart).Rank())
			Expect(foak.Value()[0]).To(Equal(7))
			Expect(foak.Value()[1]).To(Equal(11))
			Expect(foak.Value()[2]).To(Equal(14))
		})
	})

	Describe("Full House", func() {
		It("has a value of 6.xy, where x is the rank of the three of a kind and y is the rank of the pair", func() {
			fh := NewFullHouse(NewCard(Two, Spade).Rank(), NewCard(Four, Diamond).Rank())
			Expect(fh.Value()[0]).To(Equal(6))
			Expect(fh.Value()[1]).To(Equal(2))
			Expect(fh.Value()[2]).To(Equal(4))
		})
	})

	Describe("Flush", func() {
		It("has a value of 5.abcde, where abcde is the ranks of the cards in decending order", func() {
			hand := NewHand(
				NewCard(Two, Spade),
				NewCard(Four, Spade),
				NewCard(King, Spade),
				NewCard(Seven, Spade),
				NewCard(Nine, Spade))

			f := NewFlush(hand)
			Expect(f.Value()[0]).To(Equal(5))
			Expect(f.Value()[1]).To(Equal(13))
			Expect(f.Value()[2]).To(Equal(9))
			Expect(f.Value()[3]).To(Equal(7))
			Expect(f.Value()[4]).To(Equal(4))
			Expect(f.Value()[5]).To(Equal(2))
		})
	})

	Describe("Straight", func() {
		It("has a value of 4.x, where x is the high card's rank", func() {
			hand := NewHand(
				NewCard(Two, Spade),
				NewCard(Four, Spade),
				NewCard(King, Spade),
				NewCard(Seven, Spade),
				NewCard(Nine, Spade))

			f := NewFlush(hand)
			Expect(f.Value()[0]).To(Equal(5))
			Expect(f.Value()[1]).To(Equal(13))
			Expect(f.Value()[2]).To(Equal(9))
			Expect(f.Value()[3]).To(Equal(7))
			Expect(f.Value()[4]).To(Equal(4))
			Expect(f.Value()[5]).To(Equal(2))
		})
	})

	Describe("Three of a Kind", func() {
		It("Has a value of 3.xy, where x and y are kickers", func() {
			toak := NewThreeOfAKind(NewCard(Ace, Diamond).Rank(),
				NewCard(Three, Spade).Rank(),
				NewCard(Two, Heart).Rank())
			Expect(toak.Value()[0]).To(Equal(3))
			Expect(toak.Value()[1]).To(Equal(14))
			Expect(toak.Value()[2]).To(Equal(3))
			Expect(toak.Value()[3]).To(Equal(2))
		})
	})

	Describe("Two Pair", func() {
		It("Has a value of 2.xyz, where x is the rank of the high pair, y the low pair, and z the kicker", func() {
			tp := NewTwoPair(NewCard(Ace, Diamond).Rank(),
				NewCard(Three, Spade).Rank(),
				NewCard(Two, Heart).Rank())
			Expect(tp.Value()[0]).To(Equal(2))
			Expect(tp.Value()[1]).To(Equal(14))
			Expect(tp.Value()[2]).To(Equal(3))
			Expect(tp.Value()[3]).To(Equal(2))
		})
	})

	Describe("One Pair", func() {
		It("Has a value of 1.abcd, where a is the rank of the pair, and bcd the kickers", func() {
			p := NewPair(NewCard(Ace, Diamond).Rank(),
				NewCard(Three, Spade).Rank(),
				NewCard(Four, Spade).Rank(),
				NewCard(Two, Heart).Rank())
			Expect(p.Value()[0]).To(Equal(1))
			Expect(p.Value()[1]).To(Equal(14))
			Expect(p.Value()[2]).To(Equal(4))
			Expect(p.Value()[3]).To(Equal(3))
			Expect(p.Value()[4]).To(Equal(2))
		})
	})

	Describe("High Card", func() {
		It("Has a value of 0.abcde, where abcde is the rank of cards high to low", func() {
			hand := NewHand(
				NewCard(Two, Spade),
				NewCard(Four, Heart),
				NewCard(King, Spade),
				NewCard(Seven, Club),
				NewCard(Nine, Spade))

			hc := NewHighCard(hand)

			Expect(hc.Value()[0]).To(Equal(0))
			Expect(hc.Value()[1]).To(Equal(13))
			Expect(hc.Value()[2]).To(Equal(9))
			Expect(hc.Value()[3]).To(Equal(7))
			Expect(hc.Value()[4]).To(Equal(4))
			Expect(hc.Value()[5]).To(Equal(2))
		})
	})
})
