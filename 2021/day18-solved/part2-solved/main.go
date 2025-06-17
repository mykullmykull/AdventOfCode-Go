package day

func main(input []string) int {
	max := 0
	for x := range input {
		for y := range input {
			if x == y {
				continue
			}

			sum := reduce([]string{input[x], input[y]})
			magnitude := getMagnitude(sum)
			if magnitude > max {
				max = magnitude
			}
		}
	}
	return max
}
