package goker

import (
	"sort"
)

type Player struct {
	Name string
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

func (p Player) String() string {
	return p.Name
}

func NewPlayer(name string) *Player {
	p := Player{name, nil}
	return &p
}

type Pot struct {
	Value            int
	PotentialWinners map[*Player]struct{}
}

func NewPot(value int, potentialWinners []*Player) *Pot {
	potentialWinnersSet := make(map[*Player]struct{})
	for _, player := range potentialWinners {
		potentialWinnersSet[player] = struct{}{}
	}
	pot := Pot{value, potentialWinnersSet}
	return &pot
}

func Showdown(players []*Player, pots []*Pot) (map[*Player]int, []*Pot) {
	winnerTiers := Winners(players)
	payouts := make(map[*Player]int)
	oddChips := []*Pot{}

	for _, tier := range winnerTiers {
		for i, pot := range pots {
			if pot == nil {
				continue
			}

			numOfWinners := 0
			potWinners := []*Player{}
			for _, winner := range tier {
				_, exists := pot.PotentialWinners[winner]
				if exists {
					numOfWinners++
					potWinners = append(potWinners, winner)
				}
			}

			// If none of the winners on this tier can win the pot
			// (because they're only entitled to a side pot, for example)
			// it passes to a lower tier without being divided up
			if numOfWinners == 0 {
				continue
			}

			for _, winner := range potWinners {
				payouts[winner] += pot.Value / numOfWinners
			}

			if pot.Value%numOfWinners != 0 {
				oddChipPot := NewPot(pot.Value%numOfWinners, potWinners)
				oddChips = append(oddChips, oddChipPot)
			}

			// Nil out this pot so nobody can win it again!
			pots[i] = nil
		}
	}

	return payouts, oddChips
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
