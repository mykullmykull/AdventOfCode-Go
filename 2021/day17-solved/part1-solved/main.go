package day

func main(input string) int {
	t := newProbe(input)
	t.getMinVx()
	t.getMaxVy()
	return t.getMaxY()
}
