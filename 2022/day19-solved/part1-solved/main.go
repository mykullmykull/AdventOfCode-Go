package day

func main(input []string) int {
	factories := parseFactories(input)
	totalQuality := 0
	for _, f := range factories {
		maxGeos := f.getMaxGeos()
		quality := f.id * maxGeos
		println("f", f.id, "maxGeos:", maxGeos, "quality:", quality)
		totalQuality += quality
	}
	return totalQuality
}
