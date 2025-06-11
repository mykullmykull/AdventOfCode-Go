package day

func main(grid []string) int {
	c := cave{
		riskMap: map[point]int{},
		grid:    grid,
	}

	start := point{0, 0}
	end := point{len(c.grid) - 1, len(c.grid[0]) - 1}

	c.riskMap[start] = c.pointRisk(start)

	madeChange := true
	for madeChange {
		madeChange = c.minimizeRiskMap()
	}

	return c.riskMap[end] - c.riskMap[start]
}
