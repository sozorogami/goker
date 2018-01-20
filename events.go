package goker

type DrawEvent struct {
	cards CardSet
}

type ShowEvent struct {
	Player
	Hand
	HandRank
}

type WinEvent struct {
	PotNumber, PotValue int
	Winners             []*Player
}
