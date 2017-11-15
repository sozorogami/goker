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
				Expect(hand.Rank()).To(Equal(NewRoyalStraightFlush()))
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
				Expect(hand.Rank()).To(Equal(NewStraightFlush(King)))
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
				Expect(hand.Rank()).To(Equal(NewFourOfAKind(Ten, Nine)))
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
				Expect(hand.Rank()).To(Equal(NewFullHouse(Ten, Nine)))
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
				Expect(hand.Rank()).To(Equal(NewFlush(hand)))
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
				Expect(hand.Rank()).To(Equal(NewStraight(Six)))
			})

			Context("and the staight starts with an ace", func() {
				al := NewHand(
					NewCard(Two, Diamond),
					NewCard(Three, Club),
					NewCard(Four, Club),
					NewCard(Five, Club),
					NewCard(Ace, Club))
				It("has high card of 5", func() {
					Expect(al.Rank()).To(Equal(NewStraight(Five)))
				})
			})
		})

		Context("When the hand contains a three of a kind", func() {
			hand := NewHand(
				NewCard(Four, Club),
				NewCard(Four, Heart),
				NewCard(Four, Spade),
				NewCard(Seven, Club),
				NewCard(Nine, Club))
			It("is ranked a three of a kind", func() {
				Expect(hand.Rank()).To(Equal(NewThreeOfAKind(Four, Seven, Nine)))
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
				Expect(hand.Rank()).To(Equal(NewTwoPair(Four, Seven, Nine)))
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
				Expect(hand.Rank()).To(Equal(NewPair(Four, Eight, Seven, Nine)))
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
				Expect(hand.Rank()).To(Equal(NewHighCard(hand)))
			})
		})
	})

	Describe("Finding straights", func() {
		Context("When the hand contains a straight", func() {
			hand := NewHand(
				NewCard(Ten, Heart),
				NewCard(Jack, Spade),
				NewCard(Queen, Diamond),
				NewCard(King, Diamond),
				NewCard(Ace, Club))
			It("finds the straight", func() {
				Expect(hand.IsStraight()).To(BeTrue())
			})
		})

		Context("When the hand contains an ace-low straight", func() {
			hand := NewHand(
				NewCard(Ace, Diamond),
				NewCard(Two, Club),
				NewCard(Three, Heart),
				NewCard(Four, Spade),
				NewCard(Five, Diamond))
			It("finds the straight", func() {
				Expect(hand.IsStraight()).To(BeTrue())
			})
		})

		Context("When the hand contains no straight", func() {
			hand := NewHand(
				NewCard(Ace, Diamond),
				NewCard(Two, Club),
				NewCard(Three, Heart),
				NewCard(Four, Spade),
				NewCard(Six, Diamond))
			It("finds no straight", func() {
				Expect(hand.IsStraight()).To(BeFalse())
			})
		})
	})

	Describe("Finding flushes", func() {
		Context("When a hand contains a flush", func() {
			hand := NewHand(
				NewCard(Two, Diamond),
				NewCard(Four, Diamond),
				NewCard(Eight, Diamond),
				NewCard(Queen, Diamond),
				NewCard(Three, Diamond))

			It("finds the flush", func() {
				Expect(hand.IsFlush()).To(BeTrue())
			})
		})

		Context("When a hand does not contain a flush", func() {
			hand := NewHand(
				NewCard(Two, Diamond),
				NewCard(Four, Diamond),
				NewCard(Eight, Diamond),
				NewCard(Queen, Diamond),
				NewCard(Three, Heart))

			It("finds no flush", func() {
				Expect(hand.IsFlush()).To(BeFalse())
			})
		})
	})

	Describe("Finding pairs", func() {
		Context("When a hand contains a pair", func() {
			hand := NewHand(
				NewCard(Two, Diamond),
				NewCard(Two, Club),
				NewCard(King, Spade),
				NewCard(Queen, Club),
				NewCard(Three, Club))

			It("returns the rank of the pair", func() {
				Expect(hand.GroupsOf(2)[0]).To(Equal(Two))
			})
		})

		Context("When a hand contains two pair", func() {
			hand := NewHand(
				NewCard(Two, Diamond),
				NewCard(Two, Club),
				NewCard(King, Spade),
				NewCard(King, Club),
				NewCard(Three, Club))

			It("returns the rank of both pairs", func() {
				Expect(hand.GroupsOf(2)).To(ContainElement(King))
				Expect(hand.GroupsOf(2)).To(ContainElement(Two))
			})
		})

		Context("When a hand contains three of a kind", func() {
			hand := NewHand(
				NewCard(Two, Diamond),
				NewCard(Two, Club),
				NewCard(Two, Spade),
				NewCard(King, Club),
				NewCard(Three, Club))

			It("returns the rank", func() {
				Expect(hand.GroupsOf(3)[0]).To(Equal(Two))
			})

			It("does not find pairs", func() {
				Expect(hand.GroupsOf(2)).To(BeEmpty())
			})
		})

		Context("When a hand contains four of a kind", func() {
			hand := NewHand(
				NewCard(Two, Diamond),
				NewCard(Two, Club),
				NewCard(Two, Spade),
				NewCard(Two, Heart),
				NewCard(Three, Club))

			It("returns the rank", func() {
				Expect(hand.GroupsOf(4)[0]).To(Equal(Two))
			})
		})

		Context("When a hand contains no pair", func() {
			hand := NewHand(
				NewCard(Two, Diamond),
				NewCard(Four, Club),
				NewCard(King, Spade),
				NewCard(Ace, Club),
				NewCard(Three, Club))

			It("returns an empty slice", func() {
				Expect(hand.GroupsOf(2)).To(BeEmpty())
			})
		})
	})
})
