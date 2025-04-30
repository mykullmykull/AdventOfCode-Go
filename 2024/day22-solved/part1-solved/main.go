package day

func main(input []int, count int) int {
	sum := 0
	for _, monkey := range input {
		sum += evolve(monkey, count)
	}
	return sum
}

func mix(a int, b int) int {
	return a ^ b
}

func prune(n int) int {
	return n % 16777216
}

func evolve(start int, count int) int {
	for count > 0 {
		start = prune(mix(start, start*64))
		start = prune(mix(start, start/32))
		start = prune(mix(start, start*2048))
		count--
	}
	return start
}
