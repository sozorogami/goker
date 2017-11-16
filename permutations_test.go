package goker_test

import (
	. "github.com/sozorogami/goker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Combinations", func() {
	Describe("Combining a root card with a slice of card slices", func() {
		root := NewCard(Ace, Diamond)
		Context("when the slice of card slices is empty", func() {
			s := []CardSet{}
			It("returns a slice whose only element is a slice of the root card", func() {
				Expect(Concat(root, s)).To(Equal([]CardSet{CardSet{root}}))
			})
		})
		Context("when the slice of card slices is not empty", func() {
			s := []CardSet{
				CardSet{NewCard(King, Heart), NewCard(Queen, Heart)},
				CardSet{NewCard(King, Spade), NewCard(Queen, Spade)},
			}
			It("appends the root card to the front of each card slice", func() {
				expected := []CardSet{
					CardSet{root, NewCard(King, Heart), NewCard(Queen, Heart)},
					CardSet{root, NewCard(King, Spade), NewCard(Queen, Spade)},
				}
				Expect(Concat(root, s)).To(Equal(expected))
			})
		})
	})

	Describe("Finding permutations of a set of cards", func() {
		Context("when you try to choose more cards than exist", func() {
			It("panics", func() {
				Expect(func() { Combinations(1, CardSet{}) }).To(Panic())
			})
		})
		Context("when the set of cards is empty", func() {
			It("returns a set of the empty set", func() {
				Expect(Combinations(0, CardSet{})[0]).To(BeEmpty())
			})
		})
		Context("when there is one card in the set", func() {
			It("returns a set of one permutation", func() {
				cards := CardSet{NewCard(Ace, Diamond)}
				Expect(Combinations(1, cards)).To(Equal([]CardSet{cards}))
			})
		})
		Context("when there are two cards in the set", func() {
			cards := CardSet{NewCard(Ace, Diamond), NewCard(Ace, Spade)}
			Context("and you choose one", func() {
				It("returns a set of two permutations", func() {
					Expect(len(Combinations(1, cards))).To(Equal(2))
				})
			})
			Context("and you choose two", func() {
				It("returns a set of one permutation", func() {
					Expect(len(Combinations(2, cards))).To(Equal(1))
				})
			})
		})
		Context("where there are three cards in the set", func() {
			cards := CardSet{NewCard(Ace, Diamond), NewCard(Ace, Spade), NewCard(Ace, Heart)}
			Context("and you choose two", func() {
				It("returns a set of three permutations", func() {
					Expect(len(Combinations(2, cards))).To(Equal(3))
				})
			})
			Context("and you choose one", func() {
				It("returns a set of three permutations", func() {
					Expect(len(Combinations(1, cards))).To(Equal(3))
				})
			})
		})
		Context("where there are seven cards in the set", func() {
			cards := CardSet{NewCard(Ace, Diamond),
				NewCard(Ace, Spade),
				NewCard(Ace, Heart),
				NewCard(Ace, Club),
				NewCard(King, Club),
				NewCard(Queen, Club),
				NewCard(Jack, Club)}
			Context("and you choose five", func() {
				It("returns a set of 21 permutations", func() {
					Expect(len(Combinations(5, cards))).To(Equal(21))
				})
			})
			Context("and you choose three", func() {
				It("returns a set of 35 permutations", func() {
					Expect(len(Combinations(3, cards))).To(Equal(35))
				})
			})
		})
	})
})
