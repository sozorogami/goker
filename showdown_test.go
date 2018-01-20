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
	AfterEach(func() {
		charlie.MuckHand()
		dennis.MuckHand()
		dee.MuckHand()
		mac.MuckHand()
	})
	Describe("finding the winners", func() {
		Context("when there are not enough players", func() {
			It("panics", func() {
				charlie.GetHand(royalStraightFlush)
				Expect(func() {
					WinnerTiers([]*Player{charlie})
				}).To(Panic())
			})
		})
		Context("when there is a player without a hand", func() {
			It("panics", func() {
				dennis.GetHand(royalStraightFlush)
				Expect(func() {
					WinnerTiers([]*Player{dennis, charlie})
				}).To(Panic())
			})
		})
		Context("when a player has a winning hand", func() {
			winner := royalStraightFlush
			loser := highCard
			BeforeEach(func() {
				charlie.GetHand(winner)
				dennis.GetHand(loser)
			})
			It("finds one winner on two tiers", func() {
				Expect(WinnerTiers([]*Player{charlie, dennis})).To(Equal([][]*Player{[]*Player{charlie}, []*Player{dennis}}))
			})
		})
		Context("when there is a tie", func() {
			winner := royalStraightFlush
			otherWinner := otherRoyalStraightFlush
			BeforeEach(func() {
				charlie.GetHand(winner)
				dennis.GetHand(otherWinner)
			})
			It("finds two winners on one tier", func() {
				result := WinnerTiers([]*Player{charlie, dennis})[0]
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
				results, oddChips, _ := Showdown([]*Player{charlie, dennis, dee, mac}, []*Pot{onlyPot})
				Expect(results[charlie]).To(Equal(1000))
				Expect(results[dennis]).To(BeZero())
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
					results, oddChips, _ := Showdown([]*Player{charlie, dennis, dee, mac}, []*Pot{onlyPot})
					Expect(results[charlie]).To(Equal(500))
					Expect(results[dennis]).To(Equal(500))
					Expect(results[dee]).To(BeZero())
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
					results, oddChips, _ := Showdown([]*Player{charlie, dennis, dee, mac}, []*Pot{onlyPot})
					Expect(results[charlie]).To(Equal(333))
					Expect(results[dennis]).To(Equal(333))
					Expect(results[dee]).To(Equal(333))
					Expect(results[mac]).To(BeZero())
					Expect(oddChips[0]).To(Equal(NewPot(1, []*Player{charlie, dee, dennis})))
				})
			})
		})
	})

	Context("when there are side pots", func() {
		var mainPot *Pot
		var sidePot *Pot
		var pots []*Pot
		BeforeEach(func() {
			mainPot = NewPot(777, []*Player{mac, dennis, charlie})
			sidePot = NewPot(233, []*Player{mac, dennis, charlie, dee})
			pots = []*Pot{mainPot, sidePot}
		})

		var results map[*Player]int
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
				results, _, _ = Showdown([]*Player{dee, charlie, mac, dennis}, pots)
			})
			It("gives the side pot to the winnner ", func() {
				Expect(results[dee]).To(Equal(sidePot.Value))
			})
			It("gives the main pot to the runner up", func() {
				Expect(results[dennis]).To(Equal(mainPot.Value))
			})
		})
		Context("if someone in the main pot wins the hand", func() {
			BeforeEach(func() {
				dee.GetHand(loser3)
				charlie.GetHand(loser)
				dennis.GetHand(winner)
				mac.GetHand(loser2)
				results, _, _ = Showdown([]*Player{dee, charlie, mac, dennis}, pots)
			})
			It("gives both pots to the winner", func() {
				Expect(results[dennis]).To(Equal(sidePot.Value + mainPot.Value))
			})
		})
		Context("if there is a tie between someone in the main pot and the side pot", func() {
			var oddChips []*Pot
			BeforeEach(func() {
				dee.GetHand(winner)
				charlie.GetHand(loser)
				dennis.GetHand(otherRoyalStraightFlush)
				mac.GetHand(loser2)
				results, oddChips, _ = Showdown([]*Player{dee, charlie, mac, dennis}, pots)
			})
			It("gives the main pot to the person in the main pot and splits the side pot", func() {
				Expect(results[dennis]).To(Equal(mainPot.Value + sidePot.Value/2))
				Expect(results[dee]).To(Equal(sidePot.Value / 2))
			})
			It("only gives odd chips from the side pot to those who split", func() {
				Expect(oddChips[0]).To(Equal(NewPot(1, []*Player{dee, dennis})))
			})
		})
	})
})
