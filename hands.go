package goker

type hand [5]card

func NewHand(card1, card2, card3, card4, card5 card) hand {
  return [5]card{card1, card2, card3, card4, card5}
}

func (h hand) GroupsOf(n int) []rank {
  m := make(map[rank]int)
  for _, card := range h {
    m[card.rank]++
  }

  s := make([]rank, 0)
  for rank, count := range m {
    if count == n {
      s = append(s, rank)
    }
  }

  return s
}
