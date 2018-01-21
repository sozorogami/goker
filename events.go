package goker

type DrawEvent struct {
	Cards CardSet
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
