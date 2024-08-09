package day

func part1(input string) int {
	for i := 3; i < len(input); i++ {
		last4 := []byte{input[i-3], input[i-2], input[i-1], input[i]}
		if areAllUnique(last4) {
			return i + 1
		}
	}
	panic("unable to find any start of packet marker")
}

func areAllUnique(bytes []byte) bool {
	for a := 0; a < len(bytes); a++ {
		for _, b := range bytes[a+1:] {
			if b == bytes[a] {
				return false
			}
		}
	}
	return true
}
