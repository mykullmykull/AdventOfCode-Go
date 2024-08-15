package day

func part1(stacksInput []string, movesInput []string) string {
	cargo := parseStacks(stacksInput)
	cargo = cargo.executeMoves(movesInput)
	return cargo.getTops()
}
