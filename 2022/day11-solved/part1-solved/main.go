package day

func main(input []string) int {
	ms := parse(input)
	for round := 1; round <= 20; round++ {
		for id := 0; id < len(ms.monkeyMap); id++ {
			ms = ms.takeTurn(id)
		}

		println()
		println()
		println("after round", round)
		ms.printItems()
	}
	monkeyBusiness := ms.getMonkeyBusiness()
	return monkeyBusiness
}
