package goker_test

import (
	. "github.com/sozorogami/goker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Gameplay", func() {
	var state *GameState
	var rules GameRules
	var charlie, dee, mac, dennis *Player
	var players []*Player
	BeforeEach(func() {
		charlie = NewPlayer("Charlie")
		dee = NewPlayer("Dee")
		mac = NewPlayer("Mac")
		dennis = NewPlayer("Dennis")
		players = []*Player{charlie, dee, mac, dennis}
		initializePlayers(players)
	})
	Describe("inital game state", func() {
		BeforeEach(func() {
			rules = GameRules{SmallBlind: 25, BigBlind: 50}
			state = NewGame(players, rules)
		})
		It("has players", func() {
			Expect(state.Players).NotTo(BeEmpty())
		})
		It("sets a dealer", func() {
			Expect(state.Dealer).NotTo(BeNil())
		})
		It("posts small blind for the player left of the dealer", func() {
			Expect(state.Dealer.NextPlayer.CurrentBet).To(Equal(rules.SmallBlind))
		})
		It("posts big blind for the player left of the small blind", func() {
			Expect(state.Dealer.NextPlayer.NextPlayer.CurrentBet).To(Equal(rules.BigBlind))
		})
		It("sets action on the player right of the big blind", func() {
			Expect(state.Action).To(Equal(state.Dealer.NextPlayer.NextPlayer.NextPlayer))
		})
		It("sets big blind as the bet to match", func() {
			Expect(state.BetToMatch).To(Equal(rules.BigBlind))
		})
		It("does not have any pots", func() {
			Expect(state.Pots).To(BeEmpty())
		})
		It("starts at turn zero", func() {
			Expect(state.TurnNumber).To(Equal(0))
		})
		It("starts preflop", func() {
			Expect(state.BettingRound).To(Equal(Preflop))
		})
	})

	Describe("taking turns", func() {
		var action Action
		var newState GameState
		var err error

		BeforeEach(func() {
			rules = GameRules{SmallBlind: 25, BigBlind: 50}
			state = NewGame(players, rules)
		})
		Context("when someone besides action tries to take their turn", func() {
			BeforeEach(func() {
				action = Action{Player: dee, ActionType: Raise, Value: 1000}
				newState, err = Transition(*state, action)
			})
			It("rejects them", func() {
				Expect(err).NotTo(BeNil())
			})
			It("does not change the game state", func() {
				Expect(newState).To(Equal(*state))
			})
		})

		Context("when the person who is action takes their turn", func() {
			Context("when they call", func() {
				Context("and they have enough chips to cover the call", func() {
					BeforeEach(func() {
						action = Action{Player: dennis, ActionType: Call, Value: 0}
						newState, err = Transition(*state, action)
					})
					It("does not produce an error", func() {
						Expect(err).To(BeNil())
					})
					It("passes action to the next person", func() {
						Expect(newState.Action).To(Equal(charlie))
					})
					It("decreases the calling player's chips by the amount of the called bet", func() {
						Expect(dennis.Chips).To(Equal(950))
					})
					It("sets the calling player's current bet to the amount called", func() {
						Expect(dennis.CurrentBet).To(Equal(50))
					})
				})
				Context("and they do not have enough chips to cover the call", func() {
					BeforeEach(func() {
						dennis.Chips = 25
						action = Action{Player: dennis, ActionType: Call, Value: 0}
						newState, err = Transition(*state, action)
					})
					It("does not produce an error", func() {
						Expect(err).To(BeNil())
					})
					It("passes action to the next person", func() {
						Expect(newState.Action).To(Equal(charlie))
					})
					It("decreases the calling player's chips to 0", func() {
						Expect(dennis.Chips).To(Equal(0))
					})
					It("sets the player all-in", func() {
						Expect(dennis.Status).To(Equal(AllIn))
					})
					It("sets the calling player's current bet to the amount they were able to match", func() {
						Expect(dennis.CurrentBet).To(Equal(25))
					})
				})
			})
		})
	})
})

func initializePlayers(players []*Player) {
	SeatPlayers(players)
	for _, player := range players {
		player.Chips = 1000
	}
}
