package day21

func parse(input []string, steps int) grid {
	g := grid{rows: input, maxSteps: steps}
	g.start = g.getStart()
	g.centerOddity = g.getCenterOddity()
	g.subFarmSize = g.getSubFarmSize()
	g.maxFullSubFarmDistance = g.getMaxFullSubFarmDistance()
	g.edgesOddity = g.getEdgesOddity()
	g.evenSpotsCount, g.oddSpotsCount = g.countEvenOdd()
	return g
}
