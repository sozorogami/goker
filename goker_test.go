package goker_test

import (
	. "github.com/sozorogami/goker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Goker", func() {
	It("has cards", func() {
		aceOfSpades := NewCard(Ace, Spade)
		Expect(aceOfSpades.Rank).To(Equal(Ace))
		Expect(aceOfSpades.Suit).To(Equal(Spade))

		tenOfDiamonds := NewCard(Ten, Diamond)
		Expect(tenOfDiamonds.Rank).To(Equal(Ten))
		Expect(tenOfDiamonds.Suit).To(Equal(Diamond))
	})
})
