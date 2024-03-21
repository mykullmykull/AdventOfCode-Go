package day21

import "fmt"

func runB(grid []string, steps int) int {
	// runTest()
	start := getStartingLocation(grid)
	locations := []location{start}
	for i := 0; i < steps; i++ {
		// debug := i%100 == 0
		// log(fmt.Sprintf("Step %d", i), debug)
		// log(fmt.Sprintf("Locations: %d", len(locations)), debug)
		locations = getNextLocations2(locations, grid)
		println("\nStep: ", i)
		printGrid(grid, locations)
	}
	return len(locations)
}

func runTest() {
	row := -113
	println("what is -113 % 5")
	println(row % 5)
}

func getNextLocations2(locations []location, grid []string) []location {
	debug := false
	locationDeDuplicator := map[int]bool{}
	nextLocationsMap := map[int][]location{}
	nextLocations := []location{}
	for _, l := range locations {
		base := l.base(grid)
		log(fmt.Sprintf("l: %s, base: %s", l.toString(), base.toString()), debug)

		baseAdjacents := nextLocationsMap[base.address()]
		if len(baseAdjacents) == 0 {
			baseAdjacents = base.getNextLocations2(grid)
			nextLocationsMap[base.address()] = baseAdjacents
			log(fmt.Sprintf("base adj: %v, nextLocationsMap: %v", baseAdjacents, nextLocationsMap), debug)
		}
		for _, newL := range baseAdjacents {
			newL.row += l.row - base.row
			newL.col += l.col - base.col
			if _, ok := locationDeDuplicator[newL.address()]; !ok {
				locationDeDuplicator[newL.address()] = true
				nextLocations = append(nextLocations, newL)
			}
		}
	}
	return nextLocations
}
