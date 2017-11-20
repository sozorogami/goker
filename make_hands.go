package goker

import "sort"

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

func (c CardSet) BestPossibleHand() *hand {
	ph := c.PossibleHands()
	sort.Sort(ph)
	return ph[len(ph)-1]
}
