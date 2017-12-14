package goker

import "errors"

type GameState struct {
	Dealer, Action *Player
	Players        []*Player
	Pots           []*Pot
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
	if len(players) == 2 {
		dealer.Bet(rules.SmallBlind)
		dealer.NextPlayer.Bet(rules.BigBlind)
		action = dealer
	} else {
		small := dealer.NextPlayer
		small.Bet(rules.SmallBlind)
		big := small.NextPlayer
		big.Bet(rules.BigBlind)
		action = big.NextPlayer
	}

	// TODO: What if big blind can't match and goes all-in?
	currentBet := rules.BigBlind
	deck := NewDeck()

	for _, player := range players {
		player.HoleCards = deck.Draw(2)
	}

	board := CardSet{}
	gs := GameState{dealer, action, players, []*Pot{}, deck, &board, currentBet, 0, Preflop, rules}
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
	Check ActionType = iota
	Call
	Bet
	Raise
	Fold
)

func Transition(state GameState, action Action) (GameState, error) {
	if action.Player != state.Action {
		return state, errors.New("It's not your turn")
	}

	newState := state

	if action.ActionType == Call {
		action.Player.Bet(state.BetToMatch)
	}

	newState.Action = state.Action.NextPlayer
	return newState, nil
}
