package day

func main(input []string) int {
	g := parse(input)
	g = g.addSandUntilFull()
	return g.sandCount
}
