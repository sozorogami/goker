package goker_test

import (
	. "github.com/sozorogami/goker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Gameplay", func() {
	var charlie, dee, mac, dennis *Player
	Describe("inital game state", func() {
		var state *GameState
		BeforeEach(func() {
			charlie = NewPlayer("Charlie")
			dee = NewPlayer("Dee")
			mac = NewPlayer("Mac")
			dennis = NewPlayer("Dennis")
			players := []*Player{charlie, dee, mac, dennis}
			initializePlayers(players)
			state = NewGame(players)
		})
		It("has players", func() {
			Expect(state.Players).NotTo(BeEmpty())
		})
		It("sets a dealer", func() {
			Expect(state.Dealer).NotTo(BeNil())
		})
		It("sets the player to the left of dealer as active player", func() {
			Expect(state.Action).To(Equal(state.Dealer.NextPlayer))
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

func initializePlayers(players []*Player) {
	SeatPlayers(players)
	for _, player := range players {
		player.Chips = 1000
	}
}
