package goker

// Pot represents an amount of chips with metadata about which players are
// allowed to win them
type Pot struct {
	Value            int
	PotentialWinners map[*Player]struct{}
}

// NewPot constructs a pot with a given amount of chips and possible winners
func NewPot(value int, potentialWinners []*Player) *Pot {
	potentialWinnersSet := make(map[*Player]struct{})
	for _, player := range potentialWinners {
		potentialWinnersSet[player] = struct{}{}
	}
	pot := Pot{value, potentialWinnersSet}
	return &pot
}

func (p Pot) removingWinner(player *Player) *Pot {
	delete(p.PotentialWinners, player)
	return &p
}
