package day

func main(grid []string) int {
	c := crater{grid: grid, ds: initialDs}
	for x := 0; x < 10; x++ {
		c.grid = expand(c.grid)
		c = c.proposeMoves()
		c.grid = applyMoves(c.grid)
		c.ds = append(c.ds[1:], c.ds[0])
	}
	c.grid = shrink(c.grid)
	printGrid(c.grid)
	return countEmpty(c.grid)
}
