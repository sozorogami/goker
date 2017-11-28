package goker

import (
	"sort"
)

type Player struct {
	name string
	hand *Hand
}

func (p *Player) GetHand(h *Hand) {
	p.hand = h
	h.owner = p
}

func (p *Player) MuckHand() {
	if p.hand != nil {
		p.hand.owner = nil
	}
	p.hand = nil
}

func NewPlayer(name string) *Player {
	p := Player{name, nil}
	return &p
}

type Pot struct {
	value   int
	players []*Player
}

func NewPot(value int, players []*Player) *Pot {
	pot := Pot{value, players}
	return &pot
}

func Showdown(players []*Player, pots []*Pot) (map[*Player][]*Pot, int) {
	bigWinner := Winners(players)[0][0]
	payouts := make(map[*Player][]*Pot)
	payouts[bigWinner] = []*Pot{pots[0]}
	return payouts, 0
}

func Winners(players []*Player) [][]*Player {
	hands := make(HandGroup, len(players))
	for i, player := range players {
		hands[i] = player.hand
	}
	sort.Sort(hands)

	winners := [][]*Player{}
	winningHandForTier := hands[len(hands)-1]
	winnersForTier := []*Player{winningHandForTier.owner}
	for i := len(hands) - 2; i >= 0; i-- {
		hand := hands[i]
		if winningHandForTier.IsEqual(hand) {
			winnersForTier = append(winnersForTier, hand.owner)
		} else {
			winners = append(winners, winnersForTier)
			winningHandForTier = hand
			winnersForTier = []*Player{hand.owner}
		}
	}
	winners = append(winners, winnersForTier)
	return winners
}
