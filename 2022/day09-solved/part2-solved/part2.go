package day

func part2(input []string) int {
	gridSize := 51
	// grid := generateGrid(gridSize)
	start := spot{r: gridSize / 2, c: gridSize / 2}
	t := thread{
		spots:     [10]spot{start, start, start, start, start, start, start, start, start, start},
		tailSpots: map[spot]bool{},
	}
	t.tailSpots[start] = true
	for _, line := range input {
		t = t.move(line)
		// println()
		// println(line)
		// t.printGrid(grid)
	}
	return len(t.tailSpots)
}

func generateGrid(count int) []string {
	emptyRow := ""
	for x := 0; x < count; x++ {
		emptyRow += "."
	}
	grid := make([]string, count)
	for x := 0; x < count; x++ {
		grid[x] = emptyRow
	}
	return grid
}
