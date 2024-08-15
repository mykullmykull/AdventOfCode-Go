package day

func part1(input []string) int {
	gridSize := 11
	// grid := generateGrid(gridSize)
	t := thread{
		head:      spot{r: gridSize / 2, c: gridSize / 2},
		tail:      spot{r: gridSize / 2, c: gridSize / 2},
		tailSpots: map[spot]bool{},
	}
	t.tailSpots[t.tail] = true
	for _, line := range input {
		t = t.move(line)
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
