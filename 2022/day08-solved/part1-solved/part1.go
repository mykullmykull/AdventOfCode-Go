package day

func part1(input []string) int {
	p := pointer{
		visibleTrees: map[int]bool{},
	}
	for p.c = 0; p.c < len(input[0]); p.c++ {
		// add trees visible from the top
		p = p.resetMax()
		for p.r = 0; p.r < len(input); p.r++ {
			p = p.addToVisibleTrees(input)
		}

		// add trees visible from the bottom
		p = p.resetMax()
		for p.r = len(input) - 1; p.r >= 0; p.r-- {
			p = p.addToVisibleTrees(input)
		}
	}

	for p.r = 0; p.r < len(input); p.r++ {
		// add trees visible from the left
		p = p.resetMax()
		for p.c = 0; p.c < len(input[0]); p.c++ {
			p = p.addToVisibleTrees(input)
		}

		// add trees visible from the right
		p = p.resetMax()
		for p.c = len(input[0]) - 1; p.c >= 0; p.c-- {
			p = p.addToVisibleTrees(input)
		}
	}

	return len(p.visibleTrees)
}
