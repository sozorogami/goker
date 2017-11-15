package goker

import (
	"sort"
)

type HandRank interface {
	Value() []int
}

// Royal Straight Flush

type RoyalStraightFlush struct{}

func NewRoyalStraightFlush(highCard *card) *RoyalStraightFlush {
	rsf := RoyalStraightFlush{}
	return &rsf
}

func (rsf *RoyalStraightFlush) Value() []int {
	return []int{9}
}

// Straight Flush

type StraightFlush struct {
	highCard rank
}

func NewStraightFlush(highCard rank) *StraightFlush {
	sf := StraightFlush{highCard}
	return &sf
}

func (sf *StraightFlush) Value() []int {
	return []int{8, int(sf.highCard)}
}

// Four of a kind

type FourOfAKind struct {
	quad, kicker rank
}

func NewFourOfAKind(quadCard, kicker rank) *FourOfAKind {
	foak := FourOfAKind{quadCard, kicker}
	return &foak
}

func (foak *FourOfAKind) Value() []int {
	return []int{7, int(foak.quad), int(foak.kicker)}
}

// Full House

type FullHouse struct {
	trip, pair rank
}

func NewFullHouse(trip, pair rank) *FullHouse {
	fh := FullHouse{trip, pair}
	return &fh
}

func (fh *FullHouse) Value() []int {
	return []int{6, int(fh.trip), int(fh.pair)}
}

// Flush

type Flush struct {
	sortedRanks []rank
}

func NewFlush(h *hand) *Flush {
	f := Flush{}
	// Assumes hand is already sorted by rank!
	for _, card := range h {
		f.sortedRanks = append(f.sortedRanks, card.Rank())
	}
	return &f
}

func (f *Flush) Value() []int {
	val := []int{5}

	ranks := make([]int, 5)
	maxIdx := len(f.sortedRanks) - 1
	for i := maxIdx; i >= 0; i-- {
		ranks[i] = int(f.sortedRanks[maxIdx-i])
	}
	return append(val, ranks...)
}

// Straight

type Straight struct {
	highCard rank
}

func NewStraight(highCard rank) *Straight {
	s := Straight{highCard}
	return &s
}

func (s *Straight) Value() []int {
	return []int{4, int(s.highCard)}
}

// Three of a Kind

type ThreeOfAKind struct {
	trip, kicker1, kicker2 rank
}

func NewThreeOfAKind(trip, kicker1, kicker2 rank) *ThreeOfAKind {
	toak := ThreeOfAKind{trip, kicker1, kicker2}
	return &toak
}

func (toak *ThreeOfAKind) Value() []int {
	val := []int{3, int(toak.trip)}

	intRanks := rankSliceToIntSlice([]rank{toak.kicker1, toak.kicker2})
	sort.Sort(sort.Reverse(sort.IntSlice(intRanks)))

	return append(val, intRanks...)
}

// Two Pair

type TwoPair struct {
	pair1, pair2, kicker rank
}

func NewTwoPair(pair1, pair2, kicker rank) *TwoPair {
	tp := TwoPair{pair1, pair2, kicker}
	return &tp
}

func (tp *TwoPair) Value() []int {
	pairs := []rank{tp.pair1, tp.pair2}
	intRanks := rankSliceToIntSlice(pairs)
	sort.Sort(sort.Reverse(sort.IntSlice(intRanks)))
	intRanks = append([]int{2}, intRanks...)

	return append(intRanks, int(tp.kicker))
}

// Pair

type Pair struct {
	pair, kicker1, kicker2, kicker3 rank
}

func NewPair(pair, kicker1, kicker2, kicker3 rank) *Pair {
	p := Pair{pair, kicker1, kicker2, kicker3}
	return &p
}

func (p *Pair) Value() []int {
	kickers := rankSliceToIntSlice([]rank{p.kicker1, p.kicker2, p.kicker3})
	sort.Sort(sort.Reverse(sort.IntSlice(kickers)))

	val := []int{1, int(p.pair)}
	return append(val, kickers...)
}

// High Card

type HighCard struct {
	ranks []rank
}

func NewHighCard(h *hand) *HighCard {
	ranks := make([]rank, 5)
	for i, card := range h {
		ranks[i] = card.Rank()
	}
	hc := HighCard{ranks}
	return &hc
}

func (hc *HighCard) Value() []int {
	intRanks := rankSliceToIntSlice(hc.ranks)
	sort.Sort(sort.Reverse(sort.IntSlice(intRanks)))

	return append([]int{0}, intRanks...)
}
