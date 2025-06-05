package day

func main(input []string) int {
	c := caves{}
	c.init(input)

	for !c.isFinished {
		lastIndex := -1

		if c.reachedEnd() {
			c.countPaths++
			lastIndex = c.rewind()
		}
		c.addNextCave(lastIndex)
	}
	return c.countPaths
}
