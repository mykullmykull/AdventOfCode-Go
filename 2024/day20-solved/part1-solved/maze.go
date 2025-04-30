package day

func mapBaseRoute(baseMaze []string) map[point]int {
	maze := make([]string, len(baseMaze))
	copy(maze, baseMaze)
	start := findPoint(maze, 'S')
	maze = updateGrid(maze, start, "O", true)
	baseRoute := make(map[point]int)
	baseRoute[start] = 0
	steps := 0
	r := 0
	c := 0
	for {
		if c >= len(maze[r]) {
			r++
			c = 0
		}
		if r >= len(maze) {
			r = 0
			c = 0
		}
		cell := maze[r][c]
		p := point{row: r, col: c}
		c++
		if cell == '#' || cell == 'O' {
			continue
		}
		if isNextStep(maze, p) {
			steps++
			baseRoute[p] = steps
			if cell == 'E' {
				return baseRoute
			}
			maze = updateGrid(maze, p, "O", true)
		}
	}
}

func isNextStep(maze []string, p point) bool {
	neighbors := getNeighbors(maze, p)
	for _, n := range neighbors {
		if maze[n.row][n.col] == 'O' {
			return true
		}
	}
	return false
}
