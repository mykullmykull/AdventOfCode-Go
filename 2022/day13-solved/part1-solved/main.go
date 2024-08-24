package day

func main(input []string) int {
	pairs := parse(input)
	sumOfOrdered := 0
	for x, p := range pairs {
		if p.isInOrder() {
			sumOfOrdered += x + 1
			continue
		}
	}
	return sumOfOrdered
}
