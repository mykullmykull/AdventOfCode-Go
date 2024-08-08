package day21

type prev struct {
	spots map[int]map[int]bool
}

func (p prev) addToPrev(toAdd []spot) prev {
	for _, s := range toAdd {
		if p.spots == nil {
			p.spots = map[int]map[int]bool{}
		}

		if p.spots[s.r] == nil {
			p.spots[s.r] = map[int]bool{}
		}

		p.spots[s.r][s.c] = true
	}
	return p
}

func (p prev) isAlreadyCounted(s spot) bool {
	if p.spots[s.r] != nil && p.spots[s.r][s.c] {
		return true
	}
	return false
}

func (p prev) prevSpots() []spot {
	spots := []spot{}
	for r, row := range p.spots {
		for c := range row {
			spots = append(spots, spot{r, c})
		}
	}
	return spots
}
