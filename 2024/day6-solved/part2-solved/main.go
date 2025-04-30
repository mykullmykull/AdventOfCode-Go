package day

func main(grid []string) int {
	countOfLoopCausingObstacles := 0
	l := lab{grid: grid}
	l = l.findGuard()
	for r, row := range l.grid {
		for c, cell := range row {
			if cell == '^' || cell == '#' {
				continue
			}
			gCopy := make([]string, len(l.grid))
			copy(gCopy, l.grid)
			gCopy[l.guard.row] = gCopy[l.guard.row][:l.guard.col] + "." + l.grid[l.guard.row][l.guard.col+1:]
			gCopy[r] = gCopy[r][:c] + "#" + gCopy[r][c+1:]
			lCopy := lab{grid: gCopy, guard: l.guard}

			println("trying row", r, "col", c)
			causesLoop := lCopy.doesGuardLoop()
			if causesLoop {
				countOfLoopCausingObstacles++
			}
		}
	}
	return countOfLoopCausingObstacles
}
