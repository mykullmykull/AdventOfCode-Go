package day

func main(input []string) int {
	baseGrid := grid{rows: input}
	startSpots := baseGrid.getStartingSpots()
	maxPossibleSteps := len(baseGrid.rows) * len(baseGrid.rows[0])
	minSteps := maxPossibleSteps
	for _, s := range startSpots {
		g := grid{edgeSpots: []spot{s}}
		g = g.copyRows(baseGrid)
		for {
			g.printGrid()
			if len(g.edgeSpots) == 0 {
				g.steps = maxPossibleSteps
				break
			}
			if g.edgeSpots[0].elevation == 'E' {
				break
			}
			g = g.advanceOneStep()
		}
		if g.steps < minSteps {
			minSteps = g.steps
		}
	}
	return minSteps
}
