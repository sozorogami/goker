package goker

// Player represents a poker game participant
type Player struct {
	Name   string
	hand   *Hand
	chips  int
	status PlayerStatus
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
	p := Player{name, nil, 0, active}
	return &p
}

type PlayerStatus int8

const (
	active PlayerStatus = iota
	allIn
	folded
	eliminated
)
