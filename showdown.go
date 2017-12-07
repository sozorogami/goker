package goker

import (
	"sort"
)

// Player represents a poker game participant
type Player struct {
	Name string
	hand *Hand
}

// GetHand assigns ownership of the provided hand to the receiver
func (p *Player) GetHand(h *Hand) {
	p.hand = h
	h.owner = p
}

// MuckHand unassigns ownership of the player's hand
func (p *Player) MuckHand() {
	if p.hand != nil {
		p.hand.owner = nil
	}
	p.hand = nil
}

func (p Player) String() string {
	return p.Name
}

// NewPlayer constructs a player with a name and no hand
func NewPlayer(name string) *Player {
	p := Player{name, nil}
	return &p
}

// Pot represents an amount of chips with metadata about which players are
// allowed to win them
type Pot struct {
	Value            int
	PotentialWinners map[*Player]struct{}
}

// NewPot constructs a pot with a given amount of chips and possible winners
func NewPot(value int, potentialWinners []*Player) *Pot {
	potentialWinnersSet := make(map[*Player]struct{})
	for _, player := range potentialWinners {
		potentialWinnersSet[player] = struct{}{}
	}
	pot := Pot{value, potentialWinnersSet}
	return &pot
}

// Showdown takes a slice of at least two players, all of whom must have a hand,
// and a slice of all pots in play. It returns the payout in chips for each
// player when they reveal their hands and face off. It also returns a slice
// of any odd chips which could not be divided evenly during a tie.
func Showdown(players []*Player, pots []*Pot) (map[*Player]int, []*Pot) {
	winnerTiers := WinnerTiers(players)
	payouts := make(map[*Player]int)
	oddChips := []*Pot{}

	// For each winner tier, check each pot to see if any winners are entitled
	// to it, and if so, divide it among those winners
	for _, tier := range winnerTiers {
		for i, pot := range pots {

			// Skip pots that were already paid out
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
			// (because they're not entitled to a side pot, for example)
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

// WinnerTiers divides players into ranks ordered by winning poker hand,
// with all players who tied for the best hand at index 0, those who
// tied for 2nd best at index 1, and so on. There must be at least two
// players, and all players included must have a hand assigned.
func WinnerTiers(players []*Player) [][]*Player {
	if len(players) < 2 {
		panic("There must be at least two participants, or the winner is already decided.")
	}

	hands := make(HandGroup, len(players))
	for i, player := range players {
		if player.hand == nil {
			panic("All players involved in a showdown must have a hand!")
		}
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
