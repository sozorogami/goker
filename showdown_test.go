package goker_test

import (
	. "github.com/sozorogami/goker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Showdown", func() {
	charlie := NewPlayer("Charlie")
	dennis := NewPlayer("Dennis")
	dee := NewPlayer("Dee")
	mac := NewPlayer("Mac")
	AfterEach(func() {
		charlie.MuckHand()
		dennis.MuckHand()
		dee.MuckHand()
		mac.MuckHand()
	})
	Describe("finding the winners", func() {
		Context("when a player has a winning hand", func() {
			winner := royalStraightFlush
			loser := highCard
			BeforeEach(func() {
				charlie.GetHand(winner)
				dennis.GetHand(loser)
			})
			It("finds one winner on two tiers", func() {
				Expect(Winners([]*Player{charlie, dennis})).To(Equal([][]*Player{[]*Player{charlie}, []*Player{dennis}}))
			})
		})
		Context("when there is a tie", func() {
			winner := royalStraightFlush
			otherWinner := NewHand(
				NewCard(Ten, Spade),
				NewCard(Jack, Spade),
				NewCard(Queen, Spade),
				NewCard(King, Spade),
				NewCard(Ace, Spade))
			BeforeEach(func() {
				charlie.GetHand(winner)
				dennis.GetHand(otherWinner)
			})
			It("finds two winners on one tier", func() {
				result := Winners([]*Player{charlie, dennis})[0]
				Expect(result).To(ConsistOf(charlie, dennis))
			})
		})
	})
	Describe("dividing pots", func() {
		Context("", func() {})
	})
	Context("when there is only one pot", func() {
		onlyPot := NewPot(1000, []*Player{charlie, dennis, dee, mac})
		Context("and a player has a winning hand", func() {
			winner := royalStraightFlush
			loser := pair
			loser2 := flush
			loser3 := highCard
			BeforeEach(func() {
				charlie.GetHand(winner)
				dennis.GetHand(loser)
				dee.GetHand(loser2)
				mac.GetHand(loser3)
			})
			It("gives the whole pot to the winning player", func() {
				results, oddChips := Showdown([]*Player{charlie, dennis, dee, mac}, []*Pot{onlyPot})
				Expect(results[charlie]).To(Equal(1000))
				Expect(results[dennis]).To(BeZero())
				Expect(oddChips).To(BeEmpty())
			})
		})
	})
})
