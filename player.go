package goker

import "math"

// Player represents a poker game participant
type Player struct {
	Name       string
	HoleCards  CardSet
	hand       *Hand
	Chips      int
	Status     PlayerStatus
	NextPlayer *Player
	CurrentBet int
}

// GetHand assigns ownership of the provided hand to the receiver
func (p *Player) GetHand(h *Hand) {
	p.hand = h
	h.owner = p
}

// MuckHand unassigns ownership of the player's hand
func (p *Player) MuckHand() {
	if p.hand != nil {
		p.hand.owner = nil
	}
	p.hand = nil
}

func (p Player) String() string {
	return p.Name
}

func (p *Player) Bet(value int) int {
	if p.Chips <= 0 {
		panic("Player with no chips can't bet!")
	}

	var chipsBet int
	if p.Chips <= value {
		p.Status = AllIn
		chipsBet = p.Chips
		p.Chips = 0
	} else {
		chipsBet = value
		p.Chips -= value
	}

	p.CurrentBet += chipsBet
	return chipsBet
}

// NewPlayer constructs a player with a name. All other values must be
// initialized separately
func NewPlayer(name string) *Player {
	p := Player{name, nil, nil, 0, Active, nil, 0}
	return &p
}

func SeatPlayers(players []*Player) {
	if len(players) < 2 {
		panic("Not enough players to start game!")
	}
	for i := 0; i < len(players); i++ {
		if i == len(players)-1 {
			players[i].NextPlayer = players[0]
		} else {
			players[i].NextPlayer = players[i+1]
		}
	}
}

func nextActivePlayer(start *Player) *Player {
	if start.Status == Active {
		return start
	}
	for player := start.NextPlayer; player != start; player = player.NextPlayer {
		if player.Status == Active {
			return player
		}
	}
	panic("No active players!")
}

func onlyRemainingPlayer(players []*Player) *Player {
	var activePlayer *Player
	for _, player := range players {
		if player.Status == Active || player.Status == AllIn {
			if activePlayer == nil {
				activePlayer = player
			} else {
				return nil
			}
		}
	}
	return activePlayer
}

func nonFoldedPlayers(players []*Player) []*Player {
	nonFolded := []*Player{}
	for _, player := range players {
		if player.Status != Folded && player.Status != Eliminated {
			nonFolded = append(nonFolded, player)
		}
	}
	return nonFolded
}

func minBetter(players []*Player) *Player {
	minBet := math.MaxInt64
	var minBetter *Player
	for _, player := range players {
		if player.CurrentBet > 0 && player.CurrentBet < minBet {
			minBet = player.CurrentBet
			minBetter = player
		}
	}
	return minBetter
}

type PlayerStatus int8

const (
	Active PlayerStatus = iota
	AllIn
	Folded
	Eliminated
)
