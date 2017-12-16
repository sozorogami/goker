package goker

import (
	"errors"
	"fmt"
)

type GameState struct {
	Dealer, Action, ActiveBigBlind *Player
	Players                        []*Player
	Pots                           []*Pot
	*Deck
	Board      *CardSet
	BetToMatch int
	TurnNumber int
	BettingRound
	GameRules
}

type GameRules struct {
	SmallBlind int
	BigBlind   int
}

func NewGame(players []*Player, rules GameRules) *GameState {
	if len(players) < 2 {
		panic("Need at least two players for poker!")
	}

	dealer := players[0] // TODO: Randomize

	var action *Player
	var bigBlind *Player
	if len(players) == 2 {
		dealer.Bet(rules.SmallBlind)
		bigBlind = dealer.NextPlayer
		bigBlind.Bet(rules.BigBlind)
		action = dealer
	} else {
		smallBlind := dealer.NextPlayer
		smallBlind.Bet(rules.SmallBlind)
		bigBlind = smallBlind.NextPlayer
		bigBlind.Bet(rules.BigBlind)
		action = bigBlind.NextPlayer
	}

	// TODO: What if big blind can't match and goes all-in?
	currentBet := rules.BigBlind
	deck := NewDeck()

	for _, player := range players {
		player.HoleCards = deck.Draw(2)
	}

	board := CardSet{}
	gs := GameState{dealer, action, bigBlind, players, []*Pot{}, deck, &board, currentBet, 0, Preflop, rules}
	return &gs
}

type BettingRound int8

const (
	Preflop BettingRound = iota
	Flop
	Turn
	River
)

type Action struct {
	Player     *Player
	ActionType ActionType
	Value      int
}

type ActionType int8

const (
	CheckCall ActionType = iota
	BetRaise
	Fold
)

func Transition(state GameState, action Action) (GameState, error) {
	if action.Player != state.Action {
		return state, errors.New(fmt.Sprintf("It's not %s's turn\n", action.Player.Name))
	}

	newState := state

	if action.ActionType == CheckCall {
		action.Player.Bet(state.BetToMatch - action.Player.CurrentBet)
	}

	if action.Player == state.ActiveBigBlind {
		newState.ActiveBigBlind = nil
	}

	if shouldAdvanceRound(newState) {
		newState.BettingRound = state.BettingRound + 1
		newState.Action = nextActivePlayer(state.Dealer)
	} else {
		newState.Action = state.Action.NextPlayer
	}
	return newState, nil
}

func nextActivePlayer(start *Player) *Player {
	if start.Status == Active {
		return start
	}
	for player := start.NextPlayer; player != start; player = player.NextPlayer {
		if player.Status == Active {
			return player
		}
	}
	panic("No active players!")
}

func shouldAdvanceRound(state GameState) bool {
	// Don't end the turn until big blind has had a chance to act
	if state.ActiveBigBlind != nil {
		return false
	}

	for _, player := range state.Players {
		if player.Status == Active && player.CurrentBet != state.BetToMatch {
			return false
		}
	}

	return true
}
