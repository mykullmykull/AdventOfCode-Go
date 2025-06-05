package day

func (t *tangle) printTangle() {
	for _, row := range t.octopi {
		println(row)
	}
	println()
}
