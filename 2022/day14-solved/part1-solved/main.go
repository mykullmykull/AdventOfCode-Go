package day

func main(input []string) int {
	g := parse(input)
	g = g.addSandUntilVoid()
	return g.sandCount
}
