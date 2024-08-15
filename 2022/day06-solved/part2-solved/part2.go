package day

func part2(input string) int {
	for x := 13; x < len(input); x++ {
		signals := make([]string, 14)
		for y := 0; y < 14; y++ {
			signals[y] = input[x-y : x-y+1]
		}
		if areAllUnique(signals) {
			return x + 1
		}
	}
	panic("unable to find any start of packet marker")
}

func areAllUnique(signals []string) bool {
	for x := 0; x < len(signals); x++ {
		this := signals[x]
		for _, that := range signals[x+1:] {
			if that == this {
				return false
			}
		}
	}
	return true
}
