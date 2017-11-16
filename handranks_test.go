package goker_test

import (
	"sort"

	. "github.com/sozorogami/goker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("When comparing hands", func() {
	Context("if they are of different ranks", func() {
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

	Context("if they are both", func() {
		Context("straight flushes", func() {
			otherStraightFlush := NewHand(
				NewCard(Ace, Club),
				NewCard(Two, Club),
				NewCard(Three, Club),
				NewCard(Four, Club),
				NewCard(Five, Club))
			hands := HandGroup{straightFlush, otherStraightFlush}
			It("ranks the hand with higher high card higher", func() {
				sort.Sort(hands)
				Expect(hands[1]).To(Equal(straightFlush))
			})
		})

		Context("four of a kind", func() {
			Context("of different ranks", func() {
				otherFourOfAKind := NewHand(
					NewCard(Eight, Club),
					NewCard(Eight, Heart),
					NewCard(Eight, Spade),
					NewCard(Eight, Diamond),
					NewCard(Nine, Club))

				hands := HandGroup{fourOfAKind, otherFourOfAKind}
				It("ranks the hand with the higher quad higher", func() {
					sort.Sort(hands)
					Expect(hands[1]).To(Equal(fourOfAKind))
				})
			})
			Context("of the same rank", func() {
				otherFourOfAKind := NewHand(
					NewCard(Ten, Club),
					NewCard(Ten, Heart),
					NewCard(Ten, Spade),
					NewCard(Ten, Diamond),
					NewCard(Eight, Club))

				hands := HandGroup{fourOfAKind, otherFourOfAKind}
				It("ranks the hand with the higher quad higher", func() {
					sort.Sort(hands)
					Expect(hands[1]).To(Equal(fourOfAKind))
				})
			})
		})

		Context("full houses", func() {
			Context("with different trips", func() {
				otherFullHouse := NewHand(
					NewCard(Eight, Club),
					NewCard(Eight, Heart),
					NewCard(Eight, Spade),
					NewCard(Nine, Heart),
					NewCard(Nine, Spade))
				hands := HandGroup{fullHouse, otherFullHouse}
				It("ranks the hand with higher trips higher", func() {
					sort.Sort(hands)
					Expect(hands[1]).To(Equal(fullHouse))
				})
			})
			Context("with the same trips", func() {
				otherFullHouse := NewHand(
					NewCard(Ten, Club),
					NewCard(Ten, Heart),
					NewCard(Ten, Spade),
					NewCard(Eight, Diamond),
					NewCard(Eight, Club))
				hands := HandGroup{fullHouse, otherFullHouse}
				It("ranks the hand with higher pair higher", func() {
					sort.Sort(hands)
					Expect(hands[1]).To(Equal(fullHouse))
				})
			})
		})

		Context("flushes", func() {
			otherFlush := NewHand(
				NewCard(Two, Heart),
				NewCard(Five, Heart),
				NewCard(Seven, Heart),
				NewCard(Queen, Heart),
				NewCard(Jack, Heart))
			hands := HandGroup{flush, otherFlush}
			It("ranks the hand with high card higher", func() {
				sort.Sort(hands)
				Expect(hands[1]).To(Equal(flush))
			})
		})

		Context("straights", func() {
			hands := HandGroup{straight, aceLowStraight}
			It("ranks the hand with high card higher", func() {
				sort.Sort(hands)
				Expect(hands[1]).To(Equal(straight))
			})
		})

		Context("three of a kind", func() {
			Context("of different rank", func() {
				otherThreeOfAKind := NewHand(
					NewCard(Two, Club),
					NewCard(Two, Heart),
					NewCard(Two, Spade),
					NewCard(Seven, Heart),
					NewCard(Nine, Heart))
				hands := HandGroup{threeOfAKind, otherThreeOfAKind}
				It("ranks the hand with higher trips higher", func() {
					sort.Sort(hands)
					Expect(hands[1]).To(Equal(threeOfAKind))
				})
			})
			Context("of the same rank", func() {
				otherThreeOfAKind := NewHand(
					NewCard(Four, Club),
					NewCard(Four, Heart),
					NewCard(Four, Spade),
					NewCard(Seven, Heart),
					NewCard(Two, Heart))
				hands := HandGroup{threeOfAKind, otherThreeOfAKind}
				It("ranks the hand with higher kickers higher", func() {
					sort.Sort(hands)
					Expect(hands[1]).To(Equal(threeOfAKind))
				})
			})
		})

		Context("two pair", func() {
			Context("with different high pair", func() {
				otherTwoPair := NewHand(
					NewCard(Four, Club),
					NewCard(Four, Heart),
					NewCard(Six, Spade),
					NewCard(Six, Club),
					NewCard(Nine, Club))
				hands := HandGroup{twoPair, otherTwoPair}
				It("ranks the hand with higher high pair higher", func() {
					sort.Sort(hands)
					Expect(hands[1]).To(Equal(twoPair))
				})
			})
			Context("with the same high pair", func() {
				Context("but different low pair", func() {
					otherTwoPair := NewHand(
						NewCard(Three, Club),
						NewCard(Three, Heart),
						NewCard(Seven, Spade),
						NewCard(Seven, Club),
						NewCard(Nine, Club))
					hands := HandGroup{twoPair, otherTwoPair}
					It("ranks the hand with higher low pair higher", func() {
						sort.Sort(hands)
						Expect(hands[1]).To(Equal(twoPair))
					})
				})
				Context("and the same low pair", func() {
					otherTwoPair := NewHand(
						NewCard(Four, Club),
						NewCard(Four, Heart),
						NewCard(Seven, Spade),
						NewCard(Seven, Diamond),
						NewCard(Three, Club))
					hands := HandGroup{twoPair, otherTwoPair}
					It("ranks the hand with higher low pair higher", func() {
						sort.Sort(hands)
						Expect(hands[1]).To(Equal(twoPair))
					})
				})
			})
		})

		Context("pairs", func() {
			Context("of different ranks", func() {
				otherPair := NewHand(
					NewCard(Three, Club),
					NewCard(Three, Heart),
					NewCard(Eight, Spade),
					NewCard(Seven, Club),
					NewCard(Nine, Club))
				hands := HandGroup{pair, otherPair}
				It("ranks the pair with higher rank higher", func() {
					sort.Sort(hands)
					Expect(hands[1]).To(Equal(pair))
				})
			})
			Context("of the same rank", func() {
				otherPair := NewHand(
					NewCard(Four, Club),
					NewCard(Four, Heart),
					NewCard(Two, Spade),
					NewCard(Seven, Club),
					NewCard(Nine, Club))
				hands := HandGroup{pair, otherPair}
				It("ranks the hand with higher kickers higher", func() {
					sort.Sort(hands)
					Expect(hands[1]).To(Equal(pair))
				})
			})
		})

		Context("high card", func() {
			otherHighCard := NewHand(
				NewCard(Two, Club),
				NewCard(Four, Heart),
				NewCard(Eight, Spade),
				NewCard(Six, Club),
				NewCard(Nine, Spade))
			hands := HandGroup{highCard, otherHighCard}
			It("ranks the hand with higest high card higher", func() {
				sort.Sort(hands)
				Expect(hands[1]).To(Equal(highCard))
			})
		})
	})
})
