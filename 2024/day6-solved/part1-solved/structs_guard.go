package day

type guard struct {
	row int
	col int
	dir int
}

func (g guard) turnRight() guard {
	g.dir = (g.dir + 1) % 4
	return g
}
