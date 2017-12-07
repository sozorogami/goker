package goker

// Player represents a poker game participant
type Player struct {
	Name       string
	HoleCards  CardSet
	hand       *Hand
	Chips      int
	Status     PlayerStatus
	NextPlayer *Player
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

// NewPlayer constructs a player with a name. All other values must be
// initialized separately
func NewPlayer(name string) *Player {
	p := Player{name, nil, nil, 0, active, nil}
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

type PlayerStatus int8

const (
	active PlayerStatus = iota
	allIn
	folded
	eliminated
)
