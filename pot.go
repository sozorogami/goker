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

// TODO: This could surely be cleaner...
func (p Pot) sameWinners(otherPot *Pot) bool {
	for player, _ := range p.PotentialWinners {
		_, ok := otherPot.PotentialWinners[player]
		if !ok {
			return false
		}
	}
	for player, _ := range otherPot.PotentialWinners {
		_, ok := p.PotentialWinners[player]
		if !ok {
			return false
		}
	}
	return true
}

func combinePots(newPots []*Pot, existingPots []*Pot) []*Pot {
	combined := make([]*Pot, len(existingPots))
	copy(combined, existingPots)
	for _, newPot := range newPots {
		potExists := false
		for _, oldPot := range combined {
			if oldPot.sameWinners(newPot) {
				potExists = true
				oldPot.Value += newPot.Value
				newPot.Value = 0
			}
		}
		if !potExists {
			combined = append(combined, newPot)
		}
	}
	return combined
}

func separateSingletonPot(pots []*Pot) ([]*Pot, *Pot) {
	newPots := []*Pot{}
	var singleton *Pot
	for i := range pots {
		if len(pots[i].PotentialWinners) == 1 {
			singleton = pots[i]
		} else {
			newPots = append(newPots, pots[i])
		}
	}
	return newPots, singleton
}
