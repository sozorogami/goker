package goker

import (
	"sort"
)

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

			potWinners := []*Player{}
			for _, winner := range tier {
				_, exists := pot.PotentialWinners[winner]
				if exists {
					potWinners = append(potWinners, winner)
				}
			}
			numOfWinners := len(potWinners)

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

	// Sanity check
	for i := range pots {
		if pots[i] != nil {
			panic("All pots should be paid out at end of showdown!")
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
