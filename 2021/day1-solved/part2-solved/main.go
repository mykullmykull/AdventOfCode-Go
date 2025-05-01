package day

func main(input []int) int {
	increases := 0
	last := input[0]
	for x := 3; x < len(input); x++ {
		depth := input[x]
		if depth > last {
			increases++
		}
		last = input[x-2]
	}
	return increases
}
