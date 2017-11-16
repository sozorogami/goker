package goker

import (
	"sort"
)

type HandRank interface {
	Value() []int
	Name() string
}

// Royal Straight Flush

type royalStraightFlush struct{}

func newRoyalStraightFlush() *royalStraightFlush {
	rsf := royalStraightFlush{}
	return &rsf
}

func (rsf *royalStraightFlush) Value() []int {
	return []int{9}
}

func (rsf *royalStraightFlush) Name() string {
	return "RoyalStraightFlush"
}

// Straight Flush

type straightFlush struct {
	highCard rank
}

func newStraightFlush(highCard rank) *straightFlush {
	sf := straightFlush{highCard}
	return &sf
}

func (sf *straightFlush) Value() []int {
	return []int{8, int(sf.highCard)}
}

func (sf *straightFlush) Name() string {
	return "StraightFlush"
}

// Four of a kind

type fourOfAKind struct {
	quad, kicker rank
}

func newFourOfAKind(quadCard, kicker rank) *fourOfAKind {
	foak := fourOfAKind{quadCard, kicker}
	return &foak
}

func (foak *fourOfAKind) Value() []int {
	return []int{7, int(foak.quad), int(foak.kicker)}
}

func (foak *fourOfAKind) Name() string {
	return "FourOfAKind"
}

// Full House

type fullHouse struct {
	trip, pair rank
}

func newFullHouse(trip, pair rank) *fullHouse {
	fh := fullHouse{trip, pair}
	return &fh
}

func (fh *fullHouse) Value() []int {
	return []int{6, int(fh.trip), int(fh.pair)}
}

func (fh *fullHouse) Name() string {
	return "FullHouse"
}

// Flush

type flush struct {
	ranks []rank
}

func newFlush(ranks []rank) *flush {
	f := flush{ranks}
	return &f
}

func (f *flush) Value() []int {
	val := []int{5}
	ranks := rankSliceToSortedIntSlice(f.ranks)
	return append(val, ranks...)
}

func (f *flush) Name() string {
	return "Flush"
}

// Straight

type straight struct {
	highCard rank
}

func newStraight(highCard rank) *straight {
	s := straight{highCard}
	return &s
}

func (s *straight) Value() []int {
	return []int{4, int(s.highCard)}
}

func (s *straight) Name() string {
	return "Straight"
}

// Three of a Kind

type threeOfAKind struct {
	trip    rank
	kickers []rank
}

func newThreeOfAKind(trip rank, kickers []rank) *threeOfAKind {
	toak := threeOfAKind{trip, kickers}
	return &toak
}

func (toak *threeOfAKind) Value() []int {
	val := []int{3, int(toak.trip)}
	intRanks := rankSliceToSortedIntSlice(toak.kickers)
	return append(val, intRanks...)
}

func (toak *threeOfAKind) Name() string {
	return "ThreeOfAKind"
}

// Two Pair

type twoPair struct {
	pair1, pair2, kicker rank
}

func newTwoPair(pair1, pair2, kicker rank) *twoPair {
	tp := twoPair{pair1, pair2, kicker}
	return &tp
}

func (tp *twoPair) Value() []int {
	pairs := []rank{tp.pair1, tp.pair2}
	intRanks := rankSliceToSortedIntSlice(pairs)
	intRanks = append([]int{2}, intRanks...)

	return append(intRanks, int(tp.kicker))
}

func (tp *twoPair) Name() string {
	return "TwoPair"
}

// Pair

type onePair struct {
	pair    rank
	kickers []rank
}

func newPair(pair rank, kickers []rank) *onePair {
	p := onePair{pair, kickers}
	return &p
}

func (p *onePair) Value() []int {
	kickers := rankSliceToSortedIntSlice(p.kickers)

	val := []int{1, int(p.pair)}
	return append(val, kickers...)
}

func (p *onePair) Name() string {
	return "Pair"
}

// High Card

type highCard struct {
	ranks []rank
}

func newHighCard(ranks []rank) *highCard {
	hc := highCard{ranks}
	return &hc
}

func (hc *highCard) Value() []int {
	intRanks := rankSliceToSortedIntSlice(hc.ranks)

	return append([]int{0}, intRanks...)
}

func (hc *highCard) Name() string {
	return "HighCard"
}

// Helper func to sort ranks high to low for use in a rank's Value()
func rankSliceToSortedIntSlice(s []rank) []int {
	ints := make([]int, len(s))
	for i := range s {
		ints[i] = int(s[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	return ints
}
