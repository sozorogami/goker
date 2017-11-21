package goker

type HandGroup []*Hand

func (hg HandGroup) Len() int {
	return len(hg)
}

func (hg HandGroup) Swap(i, j int) {
	hg[i], hg[j] = hg[j], hg[i]
}

func (hg HandGroup) Less(i, j int) bool {
	for idx := range hg[i].Rank().Value() {
		// Compare the value of each hand rank, with decreasing
		// significance, until one is higher
		leftVal := hg[i].Rank().Value()[idx]
		rightVal := hg[j].Rank().Value()[idx]
		if leftVal < rightVal {
			return true
		}
		if leftVal > rightVal {
			return false
		}
	}
	return false
}
