package day

type machine struct {
	a     button
	b     button
	prize point
}

type button struct {
	xChange int
	yChange int
	tokens  int
}

type point struct {
	x int
	y int
}
