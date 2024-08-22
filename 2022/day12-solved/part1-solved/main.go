package day

func main(input []string) int {
	g := grid{rows: input}
	start := g.getSpotByRune('S')
	g.edgeSpots = []spot{start}
	for {
		if g.edgeSpots[0].elevation == 'E' {
			break
		}
		g = g.advanceOneStep()
	}
	return g.steps
}
