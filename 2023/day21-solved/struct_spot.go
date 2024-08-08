package day21

type spot struct {
	r int
	c int
}

func (s spot) equals(s2 spot) bool {
	return s.r == s2.r && s.c == s2.c
}

func (s spot) spotsAfterOneStep(g grid, p prev) []spot {
	nth := spot{s.r - 1, s.c}
	sth := spot{s.r + 1, s.c}
	est := spot{s.r, s.c + 1}
	wst := spot{s.r, s.c - 1}
	possible := []spot{nth, sth, est, wst}
	newLocations := []spot{}
	for _, s2 := range possible {
		isValid := s2.isValid(g, p)
		if isValid {
			newLocations = append(newLocations, s2)
		}
	}
	return newLocations
}

func (s spot) isValid(g grid, p prev) bool {
	if p.isAlreadyCounted(s) {
		return false
	}

	if s.isOutside(g) {
		return false
	}

	return g.getSpot(s) != "#"
}

func (s spot) isOutside(g grid) bool {
	return s.r < 0 || s.c < 0 || s.r >= len(g.rows) || s.c >= len(g.rows[0])
}
