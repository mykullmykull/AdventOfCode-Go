package day21

func runA(grid []string, steps int) int {
	start := getStartingLocation(grid)
	locations := []location{start}
	for i := 0; i < steps; i++ {
		locations = getNextLocations(locations, grid)
	}
	return len(locations)
}

func getStartingLocation(grid []string) location {
	for r, row := range grid {
		for c, char := range row {
			if char == 'S' {
				return location{r, c}
			}
		}
	}
	return location{}
}

func getNextLocations(locations []location, grid []string) []location {
	locationsMap := map[location]bool{}
	for _, l := range locations {
		adjacentLocations := l.getNextLocations(grid)
		for _, nl := range adjacentLocations {
			locationsMap[nl] = true
		}
	}
	nextLocations := []location{}
	for k, _ := range locationsMap {
		nextLocations = append(nextLocations, k)
	}
	return nextLocations
}
