package goker

import "sort"

// PossibleHands - Returns all possible 5 card hands that
// can be created with the cards in this set
func (c CardSet) PossibleHands() HandGroup {
	if len(c) < 5 {
		return HandGroup{}
	}

	combos := combinations(5, c)
	hands := HandGroup{}
	for _, combo := range combos {
		hands = append(hands, NewHandFromSet(combo))
	}

	return hands
}

// BestPossibleHand - Returns the best possible 5 card
// hand that can be created with the cards in this set
func (c CardSet) BestPossibleHand() *Hand {
	ph := c.PossibleHands()
	sort.Sort(ph)
	return ph[len(ph)-1]
}
