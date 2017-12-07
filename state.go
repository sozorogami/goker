package goker

type GameState struct {
	Dealer, Action *Player
	Players        []*Player
	Pots           []*Pot
	*Deck
	Board      *CardSet
	CurrentBet int
	TurnNumber int
	BettingRound
}

func NewGame(players []*Player) *GameState {
	dealer := players[0] // TODO: Randomize
	action := dealer.NextPlayer

	deck := NewDeck()

	for _, player := range players {
		player.HoleCards = deck.Draw(2)
	}

	board := CardSet{}
	gs := GameState{dealer, action, players, []*Pot{}, deck, &board, 0, 0, Preflop}
	return &gs
}

type BettingRound int8

const (
	Preflop BettingRound = iota
	Flop
	Turn
	River
)

type Action struct {
	player     *Player
	actionType ActionType
	value      int
}

type ActionType int8

const (
	Check ActionType = iota
	Call
	Bet
	Raise
	Fold
)

func Transition(state GameState, action Action) (GameState, error) {
	return GameState{}, nil
}
