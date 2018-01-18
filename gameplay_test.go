package goker_test

import (
	"strconv"
	"strings"

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

		rules = GameRules{SmallBlind: 25, BigBlind: 50}
	})
	Describe("inital game state", func() {
		BeforeEach(func() {
			state = NewGame(players, rules, NewDeck())
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
			Expect(state.HandNumber).To(Equal(0))
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
			state = NewGame(players, rules, NewDeck())
		})
		Context("when someone besides action tries to take their turn", func() {
			BeforeEach(func() {
				action = Action{Player: dee, ActionType: BetRaise, Value: 1000}
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
						action = Action{Player: dennis, ActionType: CheckCall, Value: 0}
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
						action = Action{Player: dennis, ActionType: CheckCall, Value: 0}
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

	Describe("state after a simple preflop", func() {
		BeforeEach(func() {
			// Charlie deals and Dee and Mac post blinds
			state = NewGame(players, rules, NewDeck())
			// Dennis, Charlie and Dee call, Mac checks
			newState := advance(*state, "C,C,C,C")
			state = &newState
		})
		It("has reduced all players' chips by the big blind", func() {
			for _, player := range state.Players {
				Expect(player.Chips).To(Equal(950))
			}
		})
		It("is now the flop", func() {
			Expect(state.BettingRound).To(Equal(Flop))
		})
		It("puts action to the left the dealer", func() {
			Expect(state.Action).To(Equal(state.Dealer.NextPlayer))
		})
		It("has a pot with the blinds that includes all players", func() {
			pots := state.Pots
			Expect(len(pots)).To(Equal(1))
			Expect(pots[0].Value).To(Equal(200))
			Expect(len(pots[0].PotentialWinners)).To(Equal(4))
		})
		It("sets all players' current bets to zero", func() {
			for _, player := range state.Players {
				Expect(player.CurrentBet).To(Equal(0))
			}
		})
	})

	Describe("state after everyone folds to a player", func() {
		BeforeEach(func() {
			// Charlie deals and Dee and Mac post blinds
			state = NewGame(players, rules, NewDeck())
			// Dennis calls, Charlie, Dee and Mac fold
			newState := advance(*state, "C,F,F,F")
			state = &newState
		})
		It("is the next hand", func() {
			Expect(state.HandNumber).To(Equal(1))
		})
		It("sets all players to active", func() {
			for _, player := range state.Players {
				Expect(player.Status).To(Equal(Active))
			}
		})
		It("gives the pot to the remaining player", func() {
			Expect(dennis.Chips + dennis.CurrentBet).To(Equal(1075))
		})
	})

	Describe("state when a player bets more than is matched", func() {
		BeforeEach(func() {
			dee.Chips = 200
			// Charlie deals and Dee and Mac post blinds
			state = NewGame(players, rules, NewDeck())
			// Dennis folds, Charlie raises to 500, Dee calls and Mac folds
			newState := advance(*state, "F,B450,C,F")
			state = &newState
		})
		It("is the flop", func() {
			Expect(state.BettingRound).To(Equal(Flop))
		})
		It("puts the matcher all-in", func() {
			Expect(dee.Status).To(Equal(AllIn))
		})
		It("only makes the better pay what was matched", func() {
			Expect(charlie.Chips).To(Equal(800))
		})
	})

	Describe("a hand that goes to showdown", func() {
		BeforeEach(func() {
			// Charlie deals and Dee and Mac post blinds
			state = NewGame(players, rules, NewDeck())
			charlie.HoleCards = CardSet{NewCard(Ace, Spade), NewCard(Ace, Diamond)}
			dee.HoleCards = CardSet{NewCard(King, Heart), NewCard(Queen, Heart)}
			mac.HoleCards = CardSet{NewCard(Seven, Club), NewCard(Two, Club)}
			dennis.HoleCards = CardSet{NewCard(Seven, Diamond), NewCard(Two, Diamond)}
			deck := NewStackedDeck(CardSet{
				NewCard(Ace, Heart),
				NewCard(Two, Heart),
				NewCard(Two, Spade),
				NewCard(Seven, Heart),
				NewCard(Nine, Club),
			})
			state.Deck = deck
			preflop := advance(*state, "F,B50,C,F") // Mac and Dennis fold, Dee calls Charlie's raise
			flop := advance(preflop, "C,B100,C")    // Dee checks, then calls Charlie's bet of 100
			turn := advance(flop, "C,B100,C")
			river := advance(turn, "C,B100,C")
			state = &river
		})
		It("gives the winner the loser's money", func() {
			Expect(charlie.Chips + charlie.CurrentBet).To(Equal(1450))
			Expect(dee.Chips + dee.CurrentBet).To(Equal(600))
		})
		It("advances the dealer", func() {
			Expect(state.Dealer).To(Equal(dee))
		})
	})
})

func advance(state GameState, transitions string) GameState {
	split := strings.Split(transitions, ",")

	var err error
	for _, t := range split {
		player := state.Action
		value := 0
		var actionType ActionType

		switch {
		case t == "C":
			actionType = CheckCall
		case t == "F":
			actionType = Fold
		case strings.HasPrefix(t, "B"):
			actionType = BetRaise

			amtStr := strings.TrimPrefix(t, "B")
			value, err = strconv.Atoi(amtStr)
			if err != nil {
				panic("Bad numeric value")
			}
		default:
			panic("Unable to parse string")
		}

		state, err = Transition(state, Action{Player: player, ActionType: actionType, Value: value})

		if err != nil {
			panic("Advance failed: " + err.Error())
		}
	}

	return state
}

func initializePlayers(players []*Player) {
	SeatPlayers(players)
	for _, player := range players {
		player.Chips = 1000
	}
}
