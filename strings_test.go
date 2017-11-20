package goker_test

import (
	. "github.com/sozorogami/goker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Converting cards to strings", func() {
	It("prints the rank followed by the suit", func() {
		Expect(NewCard(Ace, Spade).String()).To(Equal("A♠"))
		Expect(NewCard(King, Heart).String()).To(Equal("K♥"))
		Expect(NewCard(Queen, Diamond).String()).To(Equal("Q♦"))
		Expect(NewCard(Jack, Club).String()).To(Equal("J♣"))
		Expect(NewCard(Ten, Diamond).String()).To(Equal("T♦"))
		Expect(NewCard(Nine, Club).String()).To(Equal("9♣"))
	})
})
