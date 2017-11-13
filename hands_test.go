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

			It("Returns the rank of both pairs", func() {
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

			It("Returns the rank", func() {
				Expect(hand.GroupsOf(3)[0]).To(Equal(Two))
			})

			It("Does not find pairs", func() {
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

			It("Returns the rank", func() {
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

			It("Returns an empty slice", func() {
				Expect(hand.GroupsOf(2)).To(BeEmpty())
			})
		})
	})
})
