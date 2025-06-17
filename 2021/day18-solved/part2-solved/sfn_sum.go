package day

func sum(input []string) string {
	sum := input[0]
	input = input[1:]

	for len(input) > 0 {
		sum = reduce([]string{sum, input[0]})
		input = input[1:]
	}
	return sum
}
