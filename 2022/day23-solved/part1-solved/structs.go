package day

const (
	n = 0
	e = 1
	s = 2
	w = 3
)

func arrowRunes() []rune {
	return []rune{'<', '>', '^', 'v'}
}

func oRunes() []rune {
	return []rune{'O'}
}

func xRunes() []rune {
	return []rune{'X'}
}

var initialDs = []int{n, s, w, e}

type crater struct {
	grid []string
	ds   []int
}
