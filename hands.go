package goker

import (
	"sort"
)

type hand [5]*card

func NewHand(card1, card2, card3, card4, card5 *card) *hand {
	h := hand([5]*card{card1, card2, card3, card4, card5})
	sort.Sort(&h)
	return &h
}

func (h *hand) IsFlush() bool {
	suit := h[0].Suit()
	for _, c := range h {
		if c.Suit() != suit {
			return false
		}
	}

	return true
}

func (h *hand) IsStraight() bool {
	aceLowStraight := NewHand(
		NewCard(Ace, Club),
		NewCard(Two, Heart),
		NewCard(Three, Spade),
		NewCard(Four, Diamond),
		NewCard(Five, Club))

	if h.equalRanks(aceLowStraight) {
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

func (h *hand) GroupsOf(n int) []rank {
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

func (h *hand) equalRanks(otherHand *hand) bool {
	m1, m2 := make(map[rank]int), make(map[rank]int)

	for i := range h {
		m1[h[i].Rank()]++
		m2[otherHand[i].Rank()]++
	}

	for rank, count := range m1 {
		if m2[rank] != count {
			return false
		}
	}
	return true
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
