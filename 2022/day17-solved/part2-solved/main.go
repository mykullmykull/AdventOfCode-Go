package day

func main(winds string, maxDrops int) int {
	st := state{
		towerCols: columns{},
		rockIndex: -1, // start at -1 so we can increment at the start of getNextRock
		windIndex: -1, // start at -1 so we can increment at the start of getNextWind
	}

	sim := simulator{
		cache:  make(map[state]counters),
		counts: counters{},
		state:  st,
		winds:  winds,
	}
	sim = sim.getRocks()

	// drop rocks and cache states until we see a repeat
	sim = sim.dropRocksUntilRepeat()

	// calculate rocks to drop and height to add per repeat
	dropsPerRepeat := sim.counts.droppedCount - sim.cache[sim.state].droppedCount
	heightPerRepeat := sim.counts.height - sim.cache[sim.state].height

	// add to droppedCount and height until droppedCount is just less than maxDrops
	sim = sim.addDropsAndHeight(dropsPerRepeat, heightPerRepeat, maxDrops)

	// drop rocks until rockCount is maxDrops
	sim = sim.dropRocksUntilCount(maxDrops)
	return sim.counts.height
}
