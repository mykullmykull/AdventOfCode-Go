package day

func main(input []string) int {
	s := steam{}.parse(input)
	s = s.countExposedFaces()
	return s.count
}
