package day

func main(grid []string) int {
	c := crater{grid: grid, ds: initialDs}
	count := 0
	for {
		println(count)
		count++
		c.grid = expand(c.grid)
		moves := 0
		c, moves = c.proposeMoves()
		if moves == 0 {
			break
		}

		c.grid = applyMoves(c.grid)
		c.ds = append(c.ds[1:], c.ds[0])
	}
	c.grid = shrink(c.grid)
	printGrid(c.grid)
	return count
}
