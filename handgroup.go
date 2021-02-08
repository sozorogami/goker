package goker

// HandGroup represents a set of poker hands which can be sorted by value
type HandGroup []*Hand

func (hg HandGroup) Len() int {
	return len(hg)
}

func (hg HandGroup) Swap(i, j int) {
	hg[i], hg[j] = hg[j], hg[i]
}

func (hg HandGroup) Less(i, j int) bool {
	return hg[i].IsLessThan(hg[j])
}
