package goker

import (
	"errors"
	"fmt"
)

type GameState struct {
	Dealer, Action, ResolvingPlayer *Player
	Players                         []*Player
	Pots                            []*Pot
	*Deck
	Board                             CardSet
	BetToMatch, LastRaise, HandNumber int
	BettingRound
	GameRules
	Complete bool
	Events   []interface{}
}

type GameRules struct {
	SmallBlind int
	BigBlind   int
}

func NewGame(players []*Player, rules GameRules, deck *Deck) *GameState {
	dealer := players[0] // TODO: Randomize
	return NextHand(dealer, players, rules, deck, 0)
}

func NextHand(dealer *Player, players []*Player, rules GameRules, deck *Deck, handNumber int) *GameState {
	if len(players) < 2 {
		panic("Need at least two players for poker!")
	}

	var action *Player
	var bigBlind *Player
	if len(nonFoldedPlayers(players)) == 2 {
		dealer.Bet(rules.SmallBlind)
		bigBlind = nextActivePlayer(dealer.NextPlayer)
		bigBlind.Bet(rules.BigBlind)
		action = dealer
	} else {
		smallBlind := nextActivePlayer(dealer.NextPlayer)
		smallBlind.Bet(rules.SmallBlind)
		bigBlind = nextActivePlayer(smallBlind.NextPlayer)
		bigBlind.Bet(rules.BigBlind)
		action = nextActivePlayer(bigBlind.NextPlayer)
	}

	// TODO: What if big blind can't match and goes all-in?
	currentBet := rules.BigBlind

	for _, player := range players {
		if player.Status == Active {
			player.HoleCards = deck.Draw(2)
		}
	}

	board := CardSet{}
	gs := GameState{dealer, action, nil, players, []*Pot{}, deck, board, currentBet, 0, handNumber, Preflop, rules, false, []interface{}{}}
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

func handleCheckOrCall(state GameState) GameState {
	player := state.Action
	if state.ResolvingPlayer == nil {
		state.ResolvingPlayer = player
	}
	player.Bet(state.BetToMatch - player.CurrentBet)
	return state
}

func handleFold(state GameState) GameState {
	state.Pots = potsRemovingFoldedPlayer(state.Action, state.Pots)
	state.Action.Status = Folded
	return state
}

func handleShowdown(state GameState, pots []*Pot) {
	remainingPlayers := nonFoldedPlayers(state.Players)

	for _, player := range remainingPlayers {
		cards := append(state.Board, player.HoleCards...)
		player.GetHand(cards.BestPossibleHand())
	}

	events := Showdown(state.ResolvingPlayer, pots)
	state.Events = append(state.Events, events...)

	oddChipPots, events := PayOut(pots)

	state.Events = append(state.Events, events...)

	for _, oddChipPot := range oddChipPots {
		oddChipPayee := state.Dealer
		for oddChipPot.Value > 0 {
			_, inPot := oddChipPot.PotentialWinners[oddChipPayee]
			if inPot {
				oddChipPayee.Chips++
				oddChipPot.Value--
			}
			oddChipPayee = oddChipPayee.NextPlayer
		}
	}
}

func advanceHand(state GameState) GameState {
	refreshStatuses(state.Players)
	if onlyRemainingPlayer(state.Players) != nil {
		state.Complete = true
		return state
	}
	nextHand := NextHand(nextActivePlayer(state.Dealer.NextPlayer),
		state.Players,
		state.GameRules,
		NewDeck(),
		state.HandNumber+1)
	return *nextHand
}

func handleBetOrRaise(state GameState, increase int) (GameState, error) {
	betAmt := increase + state.BetToMatch - state.Action.CurrentBet

	if betAmt > state.Action.Chips {
		errMsg := fmt.Sprintf("%s does not have %d chips", state.Action.Name, betAmt)
		return state, errors.New(errMsg)
	}

	var minRaise int
	if state.LastRaise != 0 {
		minRaise = state.LastRaise * 2
	} else {
		minRaise = state.GameRules.BigBlind
	}
	if increase < minRaise && betAmt != state.Action.Chips {
		errMsg := fmt.Sprintf("%s must raise at least %d or go all in", state.Action.Name, minRaise)
		return state, errors.New(errMsg)
	}

	state.Action.Bet(betAmt)
	state.BetToMatch = betAmt
	state.ResolvingPlayer = state.Action
	state.LastRaise = increase
	return state, nil
}

func advanceRound(state GameState) GameState {
	if lastPlayer := onlyRemainingPlayer(state.Players); lastPlayer != nil {
		for i, pot := range state.Pots {
			state.Events = append(state.Events, WinEvent{i, pot.Value, []*Player{lastPlayer}})
		}
		pots := combinePots(gatherBets(state.Players), state.Pots)
		lastPlayer.Chips += totalPotValue(pots)
		return advanceHand(state)
	}

	pots, singleton := separateSingletonPot(gatherBets(state.Players))

	if singleton != nil {
		for player := range singleton.PotentialWinners {
			player.Chips += singleton.Value
		}
	}

	newPots := combinePots(pots, state.Pots)

	if state.BettingRound == River {
		handleShowdown(state, newPots)
		return advanceHand(state)
	}

	state.Pots = newPots
	state.BettingRound = state.BettingRound + 1
	state.BetToMatch = 0
	state.LastRaise = 0

	var cardsToDraw int
	switch state.BettingRound {
	case Flop:
		cardsToDraw = 3
	case Turn, River:
		cardsToDraw = 1
	}

	newCards := state.Deck.Draw(cardsToDraw)
	state.Events = append(state.Events, DrawEvent{newCards})
	state.Board = append(state.Board, newCards...)

	if np := nextActivePlayer(state.Dealer.NextPlayer); np != nil && nextActivePlayer(np.NextPlayer) != np {
		state.Action = np
		state.ResolvingPlayer = np
		return state
	}

	return advanceRound(state)
}

func Transition(state GameState, action Action) (GameState, error) {
	if state.Complete {
		return state, errors.New("This game is complete.")
	}

	if action.Player != state.Action {
		return state, errors.New(fmt.Sprintf("It's not %s's turn\n", action.Player.Name))
	}

	var err error
	newState := state
	newState.Events = []interface{}{}

	switch action.ActionType {
	case CheckCall:
		newState = handleCheckOrCall(newState)
	case BetRaise:
		newState, err = handleBetOrRaise(newState, action.Value)
	case Fold:
		newState = handleFold(newState)
	}

	if err != nil {
		return state, err
	}

	if shouldAdvanceRound(newState) {
		newState = advanceRound(newState)
	} else {
		newState.Action = nextActivePlayer(state.Action.NextPlayer)
	}
	return newState, nil
}

func refreshStatuses(players []*Player) {
	for _, player := range players {
		if player.Status != Eliminated {
			player.Status = Active
		}
		if player.Chips == 0 {
			player.Status = Eliminated
		}
		player.MuckHand()
		player.HoleCards = CardSet{}
	}
}

func totalPotValue(pots []*Pot) int {
	var total int
	for _, pot := range pots {
		total += pot.Value
	}
	return total
}

func potsRemovingFoldedPlayer(player *Player, pots []*Pot) []*Pot {
	cpy := make([]*Pot, len(pots))
	copy(cpy, pots)
	for _, pot := range cpy {
		delete(pot.PotentialWinners, player)
	}
	return cpy
}

func gatherBets(players []*Player) []*Pot {
	pots := []*Pot{}
	for mb := minBetter(players); mb != nil; mb = minBetter(players) {
		minBet := mb.CurrentBet
		value := 0
		potentialWinners := []*Player{}
		for _, player := range players {
			if player.CurrentBet > 0 {
				player.CurrentBet -= minBet
				value += minBet
				if player.Status != Folded {
					potentialWinners = append(potentialWinners, player)
				}
			}
		}
		pots = append(pots, NewPot(value, potentialWinners))
	}

	return pots
}

func shouldAdvanceRound(state GameState) bool {
	if onlyRemainingPlayer(state.Players) != nil {
		return true
	}
	next := nextActivePlayer(state.Action.NextPlayer)
	for between := state.Action.NextPlayer; between != next; between = between.NextPlayer {
		if between == state.ResolvingPlayer {
			return true
		}
	}
	return next == state.ResolvingPlayer || next == state.Action || next == nil
}
