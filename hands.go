package goker

import (
	"reflect"
	"sort"
)

type hand [5]*card

func NewHand(card1, card2, card3, card4, card5 *card) *hand {
	h := hand([5]*card{card1, card2, card3, card4, card5})
	sort.Sort(&h)
	return &h
}

func (h *hand) Rank() HandRank {
	if h.isFlush() && h.isStraight() {
		if h.highCard().Rank() == Ace {
			return newRoyalStraightFlush()
		}
		return newStraightFlush(h.highCard().Rank())
	}

	quads := h.groupsOf(4)
	if len(quads) != 0 {
		kicker := h.removeRanks(quads[0])[0]
		return newFourOfAKind(quads[0], kicker.Rank())
	}

	trips := h.groupsOf(3)
	pairs := h.groupsOf(2)

	if len(trips) == 1 && len(pairs) == 1 {
		return newFullHouse(trips[0], pairs[0])
	}

	if h.isFlush() {
		return newFlush(h.ranks())
	}

	if h.isStraight() {
		return newStraight(h.highCard().Rank())
	}

	if len(trips) == 1 {
		otherCards := h.removeRanks(trips[0])
		return newThreeOfAKind(trips[0], ranks(otherCards))
	}

	if len(pairs) == 2 {
		kicker := h.removeRanks(pairs...)[0]
		return newTwoPair(pairs[0], pairs[1], kicker.Rank())
	}

	if len(pairs) == 1 {
		kickers := h.removeRanks(pairs[0])
		return newPair(pairs[0], ranks(kickers))
	}

	return newHighCard(h.ranks())
}

func (h *hand) isFlush() bool {
	suit := h[0].Suit()
	for _, c := range h {
		if c.Suit() != suit {
			return false
		}
	}

	return true
}

func (h *hand) isStraight() bool {
	if h.isAceLowStraight() {
		return true
	}

	lastRank := h[0].Rank()
	for i, card := range h {
		if i == 0 {
			continue
		}

		if card.Rank() != lastRank+1 {
			return false
		}
		lastRank = card.Rank()
	}

	return true
}

// Returns all ranks such that the hand has n cards of that rank
// Examples: KKQQQ.groupsOf(2) -> [K], KKQQQ.groupsOf(3) -> [Q]
//           TTAA7.groupsOf(2) -> [T, A]
func (h *hand) groupsOf(n int) []rank {
	m := make(map[rank]int)
	for _, card := range h {
		m[card.rank]++
	}

	s := make([]rank, 0)
	for rank, count := range m {
		if count == n {
			s = append(s, rank)
		}
	}

	return s
}

// Returns the cards in a hand grouped by rank
func (h *hand) rankGroups() map[rank][]*card {
	m := make(map[rank][]*card)
	for _, card := range h {
		m[card.Rank()] = append(m[card.Rank()], card)
	}
	return m
}

func (h *hand) removeRanks(ranks ...rank) []*card {
	groups := h.rankGroups()
	for _, rank := range ranks {
		delete(groups, rank)
	}

	filtered := []*card{}
	for _, cards := range groups {
		filtered = append(filtered, cards...)
	}
	return filtered
}

// True if two hands are equal disregarding suit
func (h *hand) equalRanks(otherHand *hand) bool {
	m1, m2 := make(map[rank]int), make(map[rank]int)

	for i := range h {
		m1[h[i].Rank()]++
		m2[otherHand[i].Rank()]++
	}

	return reflect.DeepEqual(m1, m2)
}

func (h *hand) isAceLowStraight() bool {
	aceLowStraight := NewHand(
		NewCard(Ace, Club),
		NewCard(Two, Heart),
		NewCard(Three, Spade),
		NewCard(Four, Diamond),
		NewCard(Five, Club))

	return h.equalRanks(aceLowStraight)
}

// Highest card, assuming the hand is sorted and taking
// into account that A can be low card in straights
func (h *hand) highCard() *card {
	if h.isAceLowStraight() {
		return h[len(h)-2]
	}
	return h[len(h)-1]
}

func (h *hand) ranks() []rank {
	cards := h[0:]
	return ranks(cards)
}

func ranks(cards []*card) []rank {
	ranks := make([]rank, 5)
	for i, card := range cards {
		ranks[i] = card.Rank()
	}
	return ranks
}

// Sorting

func (h *hand) Len() int {
	return len(h)
}

func (h *hand) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hand) Less(i, j int) bool {
	return h[i].Rank() < h[j].Rank()
}
