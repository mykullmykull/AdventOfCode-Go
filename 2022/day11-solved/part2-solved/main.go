package day

func main(input []string) int {
	ms := parse(input)
	for round := 1; round <= 10000; round++ {
		for id := 0; id < len(ms.monkeyMap); id++ {
			ms = ms.takeTurn(id)
		}

		if round%1000 == 0 {
			println("round", round)
		}
	}
	monkeyBusiness := ms.getMonkeyBusiness()
	return monkeyBusiness
}
