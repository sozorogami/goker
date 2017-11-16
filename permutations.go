package goker

import "fmt"

func Combinations(choose int, cards CardSet) []CardSet {
	if choose > len(cards) {
		msg := fmt.Sprintf("Can't choose %d from %d cards", choose, len(cards))
		panic(msg)
	}

	if choose == 0 {
		return []CardSet{CardSet{}}
	}

	lastRootIndex := len(cards) - choose
	roots := cards[0 : lastRootIndex+1]

	s := make([]CardSet, 0)
	for i, root := range roots {
		c := Concat(root, Combinations(choose-1, cards[i+1:]))
		s = append(s, c...)
	}

	return s
}

func Concat(root Card, slice []CardSet) []CardSet {
	if len(slice) == 0 {
		return []CardSet{CardSet{root}}
	}

	s := make([]CardSet, len(slice))
	for i, subslice := range slice {
		s[i] = append(CardSet{root}, subslice...)
	}

	return s
}
