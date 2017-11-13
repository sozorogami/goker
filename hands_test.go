package goker_test

import (
	. "github.com/sozorogami/goker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Hands", func() {
	Describe("Finding pairs", func() {
		Context("When a hand contains a pair", func() {
			hand := NewHand(
				NewCard(Two, Diamond),
				NewCard(Two, Club),
				NewCard(King, Spade),
				NewCard(Queen, Club),
				NewCard(Three, Club))

			It("Returns the rank of the pair", func() {
				Expect(hand.Pairs()[0]).To(Equal(Two))
			})
		})

		Context("When a hand contains two pair", func() {
			hand := NewHand(
				NewCard(Two, Diamond),
				NewCard(Two, Club),
				NewCard(King, Spade),
				NewCard(King, Club),
				NewCard(Three, Club))

			It("Returns the rank of both pairs", func() {
				Expect(hand.Pairs()).To(ContainElement(King))
				Expect(hand.Pairs()).To(ContainElement(Two))
			})
		})

		Context("When a hand contains no pair", func() {
			hand := NewHand(
				NewCard(Two, Diamond),
				NewCard(Four, Club),
				NewCard(King, Spade),
				NewCard(Ace, Club),
				NewCard(Three, Club))

			It("Returns an empty slice", func() {
				Expect(hand.Pairs()).To(BeEmpty())
			})
		})
	})
})
