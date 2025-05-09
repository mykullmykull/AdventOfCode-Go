package day

func main(input []string) int {
	sum := 0
	for _, line := range input {
		// create monitor with initial possibility maps
		m := monitor{}
		m.init(line)

		// loop until can't simplify any more
		m.eliminateDigits()
		m.eliminateSegs()

		for m.isSimpler() {
			m.eliminateMatchedSegs()
			m.filterByHypothetical()
		}

		// panic if we exited the loop without solving everything
		if !m.isSolvable(line) {
			panic("Not all wires solved")
		}

		// solve the output and add to sum
		sum += m.solve(line)
	}
	return sum
}
