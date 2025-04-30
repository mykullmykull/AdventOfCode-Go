package day

func main(input []string) int {
	network, all, _ := parseNetwork(input)
	count := 0
	for x, c1 := range all {
		for y := x + 1; y < len(all); y++ {
			c2 := all[y]
			for z := y + 1; z < len(all); z++ {
				c3 := all[z]

				if c1[0] != 't' && c2[0] != 't' && c3[0] != 't' {
					continue
				}

				ps := []pair{
					{c1, c2},
					{c1, c3},
					{c2, c3},
				}
				if network[ps[0]] && network[ps[1]] && network[ps[2]] {
					count++
				}
			}
		}
	}
	return count
}
