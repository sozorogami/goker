package goker

import (
	"sort"
)

type HandRank interface {
	Value() []int
	Name() string
}

// Royal Straight Flush

type RoyalStraightFlush struct{}

func NewRoyalStraightFlush() *RoyalStraightFlush {
	rsf := RoyalStraightFlush{}
	return &rsf
}

func (rsf *RoyalStraightFlush) Value() []int {
	return []int{9}
}

func (rsf *RoyalStraightFlush) Name() string {
	return "RoyalStraightFlush"
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

func (sf *StraightFlush) Name() string {
	return "StraightFlush"
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

func (foak *FourOfAKind) Name() string {
	return "FourOfAKind"
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

func (fh *FullHouse) Name() string {
	return "FullHouse"
}

// Flush

type Flush struct {
	ranks []rank
}

func NewFlush(ranks []rank) *Flush {
	f := Flush{ranks}
	return &f
}

func (f *Flush) Value() []int {
	val := []int{5}
	ranks := rankSliceToSortedIntSlice(f.ranks)
	return append(val, ranks...)
}

func (f *Flush) Name() string {
	return "Flush"
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

func (s *Straight) Name() string {
	return "Straight"
}

// Three of a Kind

type ThreeOfAKind struct {
	trip    rank
	kickers []rank
}

func NewThreeOfAKind(trip rank, kickers []rank) *ThreeOfAKind {
	toak := ThreeOfAKind{trip, kickers}
	return &toak
}

func (toak *ThreeOfAKind) Value() []int {
	val := []int{3, int(toak.trip)}
	intRanks := rankSliceToSortedIntSlice(toak.kickers)
	return append(val, intRanks...)
}

func (toak *ThreeOfAKind) Name() string {
	return "ThreeOfAKind"
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
	intRanks := rankSliceToSortedIntSlice(pairs)
	intRanks = append([]int{2}, intRanks...)

	return append(intRanks, int(tp.kicker))
}

func (tp *TwoPair) Name() string {
	return "TwoPair"
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
	kickers := rankSliceToSortedIntSlice([]rank{p.kicker1, p.kicker2, p.kicker3})

	val := []int{1, int(p.pair)}
	return append(val, kickers...)
}

func (p *Pair) Name() string {
	return "Pair"
}

// High Card

type HighCard struct {
	ranks []rank
}

func NewHighCard(ranks []rank) *HighCard {
	hc := HighCard{ranks}
	return &hc
}

func (hc *HighCard) Value() []int {
	intRanks := rankSliceToSortedIntSlice(hc.ranks)

	return append([]int{0}, intRanks...)
}

func (hc *HighCard) Name() string {
	return "HighCard"
}

// Helper func to sort ranks high to low for use
// in a hand's value
func rankSliceToSortedIntSlice(s []rank) []int {
	ints := make([]int, len(s))
	for i := range s {
		ints[i] = int(s[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	return ints
}
