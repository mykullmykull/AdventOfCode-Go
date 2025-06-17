package day

func main(input string) int {
	t := newProbe(input)
	t.getMinVx()
	t.getMinVy()
	return t.countSuccessfulLaunches()
}
