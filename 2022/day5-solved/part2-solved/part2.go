package day

func part2(stacksInput []string, movesInput []string) string {
	cargo := parseStacks(stacksInput)
	cargo = cargo.executeMoves(movesInput)
	return cargo.getTops()
}
