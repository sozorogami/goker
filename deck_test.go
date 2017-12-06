package goker_test

import (
	. "github.com/sozorogami/goker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("A deck of cards", func() {
	It("starts with 52 cards", func() {
		deck := NewDeck()
		Expect(deck.Len()).To(Equal(52))
	})
	Context("when you draw cards from it", func() {
		deck := NewDeck()
		cards := deck.Draw(5)
		It("decreases in size", func() {
			Expect(deck.Len()).To(Equal(52 - 5))
		})
		It("gives you cards", func() {
			Expect(len(cards)).To(Equal(5))
		})
	})
})
