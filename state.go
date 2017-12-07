package goker

type GameState struct {
	action, dealer *Player
	players        []*Player
	pots           []*Pot
	currentBet     int
	turnNumber     int
	round          BettingRound
}

type BettingRound int8

const (
	preflop BettingRound = iota
	flop
	turn
	river
)
