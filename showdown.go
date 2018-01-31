package goker

func Showdown(resolvingPlayer *Player, pots []*Pot) []interface{} {
	p := resolvingPlayer
	events := []interface{}{}
	for {
		mucked := true
		for _, pot := range pots {
			if _, ok := pot.PotentialWinners[p]; ok {
				if len(pot.Winners) == 0 || p.hand.IsEqual(pot.Winners[0].hand) {
					pot.Winners = append(pot.Winners, p)
					mucked = false
				} else if !p.hand.IsLessThan(pot.Winners[0].hand) {
					pot.Winners = []*Player{p}
					mucked = false
				}
			}
		}

		if mucked {
			events = append(events, MuckEvent{p})
		} else {
			events = append(events, ShowEvent{p, p.hand, p.hand.Rank()})
		}

		p = nextActivePlayer(p.NextPlayer)
		if p == resolvingPlayer || p == nil {
			break
		}
	}
	return events
}

func PayOut(pots []*Pot) ([]*Pot, []interface{}) {
	events := []interface{}{}
	oddChipPots := []*Pot{}
	for i, pot := range pots {
		events = append(events, WinEvent{i, pot.Value, pot.Winners})
		payout := pot.Value / len(pot.Winners)
		for _, winner := range pot.Winners {
			winner.Chips += payout
		}
		if oddChips := pot.Value % len(pot.Winners); oddChips != 0 {
			potential := make(map[*Player]struct{})
			for _, player := range pot.Winners {
				potential[player] = struct{}{}
			}
			oddChipPot := Pot{oddChips, potential, []*Player{}}
			oddChipPots = append(oddChipPots, &oddChipPot)
		}
	}
	return oddChipPots, events
}
