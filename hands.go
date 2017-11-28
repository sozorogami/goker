package goker

import (
	"reflect"
	"sort"
)

// Hand - Represents a 5 card poker hand
type Hand [5]Card

// NewHand - Returns a pointer to a new hand consisting of
// the required cards
func NewHand(card1, card2, card3, card4, card5 Card) *Hand {
	h := Hand([5]Card{card1, card2, card3, card4, card5})
	sort.Sort(&h)
	return &h
}

// NewHandFromSet - Returns a pointer to a new hand consisting of
// the first five cards in the card set provided
func NewHandFromSet(cards CardSet) *Hand {
	return NewHand(cards[0], cards[1], cards[2], cards[3], cards[4])
}

// Rank - The relative value of the hand according to the rules
// of poker
func (h *Hand) Rank() HandRank {
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
		otherCardSet := h.removeRanks(trips[0])
		return newThreeOfAKind(trips[0], ranks(otherCardSet))
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

func (h *Hand) isFlush() bool {
	suit := h[0].Suit()
	for _, c := range h {
		if c.Suit() != suit {
			return false
		}
	}

	return true
}

func (h *Hand) isStraight() bool {
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

// Returns all ranks such that the hand has n cards of that rank.
// The order of the ranks in the slice returned is undefined.
// Examples: QQQKK.groupsOf(2) -> [K], KKQQQ.groupsOf(3) -> [Q]
//           7TTAA.groupsOf(2) -> [T, A]
func (h *Hand) groupsOf(n int) []rank {
	m := make(map[rank]int)
	for _, card := range h {
		m[card.Rank()]++
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
func (h *Hand) rankGroups() map[rank][]Card {
	m := make(map[rank][]Card)
	for _, card := range h {
		m[card.Rank()] = append(m[card.Rank()], card)
	}
	return m
}

// Returns a slice of the cards that would remain if all the specified
// ranks were removed
// e.g. [9♠︎, 9♦︎, K♠︎, K♦︎, A♦︎].removeRanks(A, 9) -> [K♠︎, K♦︎]
func (h *Hand) removeRanks(ranks ...rank) []Card {
	groups := h.rankGroups()
	for _, rank := range ranks {
		delete(groups, rank)
	}

	filtered := []Card{}
	for _, cards := range groups {
		filtered = append(filtered, cards...)
	}
	return filtered
}

// True if two hands are equal disregarding suit
func (h *Hand) equalRanks(otherHand *Hand) bool {
	m1, m2 := make(map[rank]int), make(map[rank]int)

	for i := range h {
		m1[h[i].Rank()]++
		m2[otherHand[i].Rank()]++
	}

	return reflect.DeepEqual(m1, m2)
}

func (h *Hand) isAceLowStraight() bool {
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
func (h *Hand) highCard() Card {
	if h.isAceLowStraight() {
		return h[len(h)-2]
	}
	return h[len(h)-1]
}

func (h *Hand) ranks() []rank {
	cards := h[0:]
	return ranks(cards)
}

func ranks(cards []Card) []rank {
	ranks := make([]rank, 5)
	for i, card := range cards {
		ranks[i] = card.Rank()
	}
	return ranks
}

// Sorting

func (h *Hand) Len() int {
	return len(h)
}

func (h *Hand) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Hand) Less(i, j int) bool {
	return h[i].Rank() < h[j].Rank()
}
