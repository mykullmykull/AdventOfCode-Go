package day

func (v valley) getNewElfs() map[point]bool {
	newElfs := map[point]bool{}
	for current := range v.elfs {
		upPoint := point{r: current.r - 1, c: current.c}
		dnPoint := point{r: current.r + 1, c: current.c}
		ltPoint := point{r: current.r, c: current.c - 1}
		rtPoint := point{r: current.r, c: current.c + 1}
		newElfs = v.addIfValid(newElfs, upPoint)
		newElfs = v.addIfValid(newElfs, dnPoint)
		newElfs = v.addIfValid(newElfs, ltPoint)
		newElfs = v.addIfValid(newElfs, rtPoint)
		newElfs = v.addIfValid(newElfs, current)
	}
	return newElfs
}

func (v valley) addIfValid(newElfs map[point]bool, p point) map[point]bool {
	if p.r < 0 || p.r >= len(v.grid) || p.c < 1 || p.c >= len(v.grid[0])-1 {
		return newElfs
	}
	if v.grid[p.r][p.c] == '.' {
		newElfs[p] = true
	}
	return newElfs
}

func (v valley) haveReachedTheFinish() bool {
	for p := range v.elfs {
		if p.r == len(v.grid)-1 {
			return true
		}
	}
	return false
}

func (v valley) haveReachedTheStart() bool {
	for p := range v.elfs {
		if p.r == 0 {
			return true
		}
	}
	return false
}
