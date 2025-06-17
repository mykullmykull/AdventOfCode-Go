package day

func sum(input []string) string {
	sum := input[0]
	input = input[1:]

	for len(input) > 0 {
		// println()
		// println(" ", sum)
		// println("+", input[0])
		sum = reduce([]string{sum, input[0]})
		// println("=", sum)
		input = input[1:]
	}
	return sum
}
