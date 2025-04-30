package day

type simulator struct {
	winds  string
	wind   string
	rocks  []rock
	rock   rock
	counts counters
	state  state
	cache  map[state]counters
}

func (s simulator) getRocks() simulator {
	s.rocks = []rock{
		{r1},
		{r2},
		{r3},
		{r4},
		{r5},
	}
	return s
}

func (s simulator) getNextRock() simulator {
	s.state.rockIndex = (s.state.rockIndex + 1) % len(s.rocks)
	s.rock = s.rocks[s.state.rockIndex]
	return s
}

func (s simulator) getNextWind() simulator {
	s.state.windIndex = (s.state.windIndex + 1) % len(s.winds)
	s.wind = string(s.winds[s.state.windIndex])
	return s
}

func (s simulator) addRockToTop() simulator {
	rockHeight := s.rock.getHeight()
	s.state.towerCols = s.state.towerCols.addHeight(rockHeight)
	s.counts.height += rockHeight
	return s
}

func (s simulator) addDropsAndHeight(dropsPerRepeat int, heightPerRepeat int, dropsMax int) simulator {
	repeats := (dropsMax - s.counts.droppedCount) / dropsPerRepeat
	s.counts.droppedCount += dropsPerRepeat * repeats
	s.counts.height += heightPerRepeat * repeats
	return s
}

func (s simulator) rockCanStillDrop() bool {
	for _, point := range s.rock.points {
		nextRow := point[0] + 1
		col := point[1]
		if s.state.towerCols.isBlocked(col, nextRow) {
			return false
		}
	}
	return true
}

func (s simulator) addSettledRock() simulator {
	// add rock to spaces
	for _, point := range s.rock.points {
		s.state.towerCols = s.state.towerCols.markRock(point)
	}
	s.counts.droppedCount++
	return s
}

func (s simulator) dropRock() simulator {
	// move rock down if spaces are minimized
	if s.state.towerCols.getMinSpace() == 0 {
		s.rock = s.rock.drop()
		return s
	}
	// else move spaces up
	s.state.towerCols = s.state.towerCols.moveUp()
	s.counts.height--
	return s
}

func (s simulator) dropNextRock() simulator {
	s = s.getNextRock()

	// adjust rock horizontally with the next 4 winds
	for x := 0; x < 4; x++ {
		s = s.getNextWind()
		s.rock = s.rock.blowRock(s.wind)
	}

	s = s.addRockToTop()
	for s.rockCanStillDrop() {
		s = s.dropRock()
		s = s.getNextWind()
		s.rock = s.rock.blowFallenRock(s.wind, s.state.towerCols)
	}
	s = s.addSettledRock()
	s.state.towerCols = s.state.towerCols.shrinkToMaxSpace()

	return s
}

func (s simulator) dropRocksUntilRepeat() simulator {
	for {
		s = s.dropNextRock()
		if _, ok := s.cache[s.state]; ok {
			return s
		}
		s.cache[s.state] = counters{s.counts.droppedCount, s.counts.height}
	}
}

func (s simulator) dropRocksUntilCount(dropsMax int) simulator {
	for s.counts.droppedCount < dropsMax {
		s = s.dropNextRock()
	}
	return s
}
