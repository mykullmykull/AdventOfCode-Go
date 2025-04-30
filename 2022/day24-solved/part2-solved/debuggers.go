package day

func (v valley) print() {
	copy := make([]string, len(v.grid))
	for x, row := range v.grid {
		copy[x] = row
	}
	for p := range v.elfs {
		copy[p.r] = copy[p.r][:p.c] + "E" + copy[p.r][p.c+1:]
	}
	for _, row := range copy {
		println(row)
	}
	println()
}
