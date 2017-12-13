package goker_test

import (
	. "github.com/sozorogami/goker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Gameplay", func() {
	Describe("inital game state", func() {
		var state *GameState
		var rules GameRules
		BeforeEach(func() {
			rules = GameRules{SmallBlind: 25, BigBlind: 50}
			state = initialState(rules)
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
})

func initialState(rules GameRules) *GameState {
	players := initializePlayers()
	return NewGame(players, rules)
}

func initializePlayers() []*Player {
	charlie := NewPlayer("Charlie")
	dee := NewPlayer("Dee")
	mac := NewPlayer("Mac")
	dennis := NewPlayer("Dennis")
	players := []*Player{charlie, dee, mac, dennis}
	SeatPlayers(players)
	for _, player := range players {
		player.Chips = 1000
	}
	return players
}
