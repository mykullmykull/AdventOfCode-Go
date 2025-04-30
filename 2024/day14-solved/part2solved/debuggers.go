package day

func printGrid(robots []robot, width int, height int) {
	grid := make([]string, height)
	for row := 0; row < height; row++ {
		grid[row] = ""
		for col := 0; col < width; col++ {
			grid[row] += "."
		}
	}
	for _, r := range robots {
		grid[r.pos.y] = grid[r.pos.y][:r.pos.x] + "#" + grid[r.pos.y][r.pos.x+1:]
	}
	for _, row := range grid {
		println(row)
	}
}
