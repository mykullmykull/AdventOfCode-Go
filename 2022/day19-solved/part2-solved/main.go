package day

func main(input []string) int {
	factories := parseFactories(input)
	product := 1
	for _, f := range factories[:3] {
		maxGeos := f.getMaxGeos()
		println("f", f.id, "maxGeos:", maxGeos)
		println()
		product = product * maxGeos
	}
	return product
}
