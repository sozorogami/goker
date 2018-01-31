package goker_test

import (
	. "github.com/sozorogami/goker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var otherRoyalStraightFlush = NewHand(
	NewCard(Ten, Spade),
	NewCard(Jack, Spade),
	NewCard(Queen, Spade),
	NewCard(King, Spade),
	NewCard(Ace, Spade))

var _ = Describe("Showdown", func() {
	charlie := NewPlayer("Charlie")
	dennis := NewPlayer("Dennis")
	dee := NewPlayer("Dee")
	mac := NewPlayer("Mac")
	BeforeEach(func() {
		SeatPlayers([]*Player{charlie, dee, mac, dennis})
	})
	AfterEach(func() {
		charlie.MuckHand()
		dennis.MuckHand()
		dee.MuckHand()
		mac.MuckHand()
	})

	Describe("dividing pots", func() {
		Context("", func() {})
	})
	Context("when there is only one pot", func() {
		var pots []*Pot
		BeforeEach(func() {
			onlyPot := NewPot(1000, []*Player{charlie, dennis, dee, mac})
			pots = []*Pot{onlyPot}
		})
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
				charlieChipsBefore := charlie.Chips
				dennisChipsBefore := dennis.Chips
				_ = Showdown(charlie, pots)
				oddChips, _ := PayOut(pots)
				Expect(charlie.Chips).To(Equal(charlieChipsBefore + 1000))
				Expect(dennis.Chips).To(Equal(dennisChipsBefore))
				Expect(oddChips).To(BeEmpty())
			})
		})
		Context("and there is a tie", func() {
			Context("where the pot can be evenly divided", func() {
				winner1 := royalStraightFlush
				winner2 := otherRoyalStraightFlush
				loser1 := flush
				loser2 := highCard
				BeforeEach(func() {
					charlie.GetHand(winner1)
					dennis.GetHand(winner2)
					dee.GetHand(loser1)
					mac.GetHand(loser2)
				})
				It("divides the pot evenly between the winners", func() {
					charlieChipsBefore := charlie.Chips
					dennisChipsBefore := dennis.Chips
					deeChipsBefore := dee.Chips
					_ = Showdown(charlie, pots)
					oddChips, _ := PayOut(pots)
					Expect(charlie.Chips).To(Equal(charlieChipsBefore + 500))
					Expect(dennis.Chips).To(Equal(dennisChipsBefore + 500))
					Expect(dee.Chips).To(Equal(deeChipsBefore))
					Expect(oddChips).To(BeEmpty())
				})
			})
			Context("where the pot cannot be divided evenly", func() {
				yetAnotherRoyalStraightFlush := NewHand(
					NewCard(Ten, Diamond),
					NewCard(Jack, Diamond),
					NewCard(Queen, Diamond),
					NewCard(King, Diamond),
					NewCard(Ace, Diamond))

				winner1 := royalStraightFlush
				winner2 := otherRoyalStraightFlush
				winner3 := yetAnotherRoyalStraightFlush
				loser := highCard
				BeforeEach(func() {
					charlie.GetHand(winner1)
					dennis.GetHand(winner2)
					dee.GetHand(winner3)
					mac.GetHand(loser)
				})
				It("divides the pot evenly, with odd chips returned separately", func() {
					charlieChipsBefore := charlie.Chips
					dennisChipsBefore := dennis.Chips
					deeChipsBefore := dee.Chips
					macChipsBefore := mac.Chips

					_ = Showdown(charlie, pots)
					oddChips, _ := PayOut(pots)

					Expect(charlie.Chips).To(Equal(charlieChipsBefore + 333))
					Expect(dennis.Chips).To(Equal(dennisChipsBefore + 333))
					Expect(dee.Chips).To(Equal(deeChipsBefore + 333))
					Expect(mac.Chips).To(Equal(macChipsBefore))
					Expect(oddChips[0]).To(Equal(NewPot(1, []*Player{charlie, dee, dennis})))
				})
			})
		})
	})

	Context("when there are side pots", func() {
		var mainPot *Pot
		var sidePot *Pot
		var pots []*Pot
		var deeChipsBefore int
		var dennisChipsBefore int
		BeforeEach(func() {
			mainPot = NewPot(777, []*Player{mac, dennis, charlie})
			sidePot = NewPot(233, []*Player{mac, dennis, charlie, dee})
			pots = []*Pot{mainPot, sidePot}
			deeChipsBefore = dee.Chips
			dennisChipsBefore = dennis.Chips
		})

		winner := royalStraightFlush
		loser := pair
		loser2 := flush
		loser3 := highCard
		Context("if the person in the side pot wins the hand", func() {
			BeforeEach(func() {
				dee.GetHand(winner)
				charlie.GetHand(loser)
				dennis.GetHand(loser2)
				mac.GetHand(loser3)
				_ = Showdown(charlie, pots)
				_, _ = PayOut(pots)
			})
			It("gives the side pot to the winnner ", func() {
				Expect(dee.Chips).To(Equal(deeChipsBefore + sidePot.Value))
			})
			It("gives the main pot to the runner up", func() {
				Expect(dennis.Chips).To(Equal(dennisChipsBefore + mainPot.Value))
			})
		})
		Context("if someone in the main pot wins the hand", func() {
			BeforeEach(func() {
				dee.GetHand(loser3)
				charlie.GetHand(loser)
				dennis.GetHand(winner)
				mac.GetHand(loser2)
				_ = Showdown(charlie, pots)
				_, _ = PayOut(pots)
			})
			It("gives both pots to the winner", func() {
				Expect(dennis.Chips).To(Equal(dennisChipsBefore + sidePot.Value + mainPot.Value))
			})
		})
		Context("if there is a tie between someone in the main pot and the side pot", func() {
			var oddChips []*Pot
			BeforeEach(func() {
				dee.GetHand(winner)
				charlie.GetHand(loser)
				dennis.GetHand(otherRoyalStraightFlush)
				mac.GetHand(loser2)
				_ = Showdown(charlie, pots)
				oddChips, _ = PayOut(pots)
			})
			It("gives the main pot to the person in the main pot and splits the side pot", func() {
				Expect(dennis.Chips).To(Equal(dennisChipsBefore + mainPot.Value + sidePot.Value/2))
				Expect(dee.Chips).To(Equal(deeChipsBefore + sidePot.Value/2))
			})
			It("only gives odd chips from the side pot to those who split", func() {
				Expect(oddChips[0]).To(Equal(NewPot(1, []*Player{dee, dennis})))
			})
		})
	})
})
