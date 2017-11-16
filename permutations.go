package goker

import "fmt"

func Combinations(choose int, cards []PlayingCard) [][]PlayingCard {
	if choose > len(cards) {
		msg := fmt.Sprintf("Can't choose %d from %d cards", choose, len(cards))
		panic(msg)
	}

	if choose == 0 {
		return [][]PlayingCard{[]PlayingCard{}}
	}

	lastRootIndex := len(cards) - choose
	roots := cards[0 : lastRootIndex+1]

	s := make([][]PlayingCard, 0)
	for i, root := range roots {
		c := Concat(root, Combinations(choose-1, cards[i+1:]))
		s = append(s, c...)
	}

	return s
}

func Concat(root PlayingCard, slice [][]PlayingCard) [][]PlayingCard {
	if len(slice) == 0 {
		return [][]PlayingCard{[]PlayingCard{root}}
	}

	s := make([][]PlayingCard, len(slice))
	for i, subslice := range slice {
		s[i] = append([]PlayingCard{root}, subslice...)
	}

	return s
}
