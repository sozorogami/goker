package goker

import (
	"errors"
	"fmt"
	"math"
)

type GameState struct {
	Dealer, Action, ResolvingPlayer *Player
	Players                         []*Player
	Pots                            []*Pot
	*Deck
	Board      CardSet
	BetToMatch int
	LastRaise  int
	HandNumber int
	BettingRound
	GameRules
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

	for _, player := range players {
		player.HoleCards = deck.Draw(2)
	}

	board := CardSet{}
	gs := GameState{dealer, action, nil, players, []*Pot{}, deck, board, currentBet, 0, handNumber, Preflop, rules}
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
		state.ResolvingPlayer = state.Action
	}
	player.Bet(state.BetToMatch - player.CurrentBet)
	return state
}

func handleFold(state GameState) GameState {
	newPots := potsRemovingFoldedPlayer(state.Action, state.Pots)
	state.Action.Status = Folded

	// Award the pot immediately and advance hand if only one player remains
	if lastPlayer := onlyRemainingPlayer(state.Players); lastPlayer != nil {
		newPots = combinePots(gatherBets(state.Players), newPots)
		lastPlayer.Chips += totalPotValue(newPots)
		return advanceHand(state)
	}

	state.Pots = newPots
	return state
}

func handleShowdown(state GameState, pots []*Pot) {
	remainingPlayers := nonFoldedPlayers(state.Players)

	for _, player := range remainingPlayers {
		cards := append(state.Board, player.HoleCards...)
		player.GetHand(cards.BestPossibleHand())
	}

	payouts, oddChipPots := Showdown(remainingPlayers, pots)

	for player, winnings := range payouts {
		player.Chips += winnings
	}

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
	nextHand := NextHand(nextActivePlayer(state.Dealer),
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
	state.Action = nextActivePlayer(state.Dealer.NextPlayer)
	state.BetToMatch = 0
	state.LastRaise = 0
	state.ResolvingPlayer = state.Action

	var cardsToDraw int
	switch state.BettingRound {
	case Flop:
		cardsToDraw = 3
	case Turn, River:
		cardsToDraw = 1
	}

	state.Board = append(state.Board, state.Deck.Draw(cardsToDraw)...)
	return state
}

func Transition(state GameState, action Action) (GameState, error) {
	if action.Player != state.Action {
		return state, errors.New(fmt.Sprintf("It's not %s's turn\n", action.Player.Name))
	}

	var err error
	newState := state

	switch action.ActionType {
	case CheckCall:
		newState = handleCheckOrCall(state)
	case BetRaise:
		newState, err = handleBetOrRaise(state, action.Value)
	case Fold:
		newState = handleFold(state)
	}

	if err != nil {
		return state, err
	}

	if shouldAdvanceRound(newState) {
		newState = advanceRound(state)
	} else {
		newState.Action = nextActivePlayer(state.Action.NextPlayer)
	}
	return newState, nil
}

func nonFoldedPlayers(players []*Player) []*Player {
	nonFolded := []*Player{}
	for _, player := range players {
		if player.Status != Folded && player.Status != Eliminated {
			nonFolded = append(nonFolded, player)
		}
	}
	return nonFolded
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

func combinePots(newPots []*Pot, existingPots []*Pot) []*Pot {
	combined := make([]*Pot, len(existingPots))
	copy(combined, existingPots)
	for _, newPot := range newPots {
		potExists := false
		for _, oldPot := range combined {
			if sameWinners(oldPot, newPot) {
				potExists = true
				oldPot.Value += newPot.Value
				newPot.Value = 0
			}
		}
		if !potExists {
			combined = append(combined, newPot)
		}
	}
	return combined
}

// TODO: This could surely be cleaner...
func sameWinners(pot1, pot2 *Pot) bool {
	for player, _ := range pot1.PotentialWinners {
		_, ok := pot2.PotentialWinners[player]
		if !ok {
			return false
		}
	}
	for player, _ := range pot2.PotentialWinners {
		_, ok := pot1.PotentialWinners[player]
		if !ok {
			return false
		}
	}
	return true
}

func gatherBets(players []*Player) []*Pot {
	pots := []*Pot{}
	for minBetter := findMinBetter(players); minBetter != nil; minBetter = findMinBetter(players) {
		minBet := minBetter.CurrentBet
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

func separateSingletonPot(pots []*Pot) ([]*Pot, *Pot) {
	newPots := []*Pot{}
	var singleton *Pot
	for i := range pots {
		if len(pots[i].PotentialWinners) == 1 {
			singleton = pots[i]
		} else {
			newPots = append(newPots, pots[i])
		}
	}
	return newPots, singleton
}

func findMinBetter(players []*Player) *Player {
	minBet := math.MaxInt64
	var minBetter *Player
	for _, player := range players {
		if player.CurrentBet > 0 && player.CurrentBet < minBet {
			minBet = player.CurrentBet
			minBetter = player
		}
	}
	return minBetter
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

func onlyRemainingPlayer(players []*Player) *Player {
	var activePlayer *Player
	for _, player := range players {
		if player.Status == Active || player.Status == AllIn {
			if activePlayer == nil {
				activePlayer = player
			} else {
				return nil
			}
		}
	}
	return activePlayer
}

func shouldAdvanceRound(state GameState) bool {
	next := nextActivePlayer(state.Action.NextPlayer)
	return next == state.ResolvingPlayer || next == state.Action
}
