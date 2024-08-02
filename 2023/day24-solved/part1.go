package day24

func runA(input []string, min float64, max float64) float64 {
	stones := parseStns(input)
	combos := getCombos(stones)
	willCross := float64(0)
	for _, c := range combos {
		if c.yxWillCrossPaths(min, max) {
			willCross++
		}
	}
	return willCross
}

func getCombos(stones []stn) []combo {
	combos := []combo{}
	for i, s1 := range stones {
		for j, s2 := range stones {
			if i >= j {
				continue
			}
			newCombo := combo{s1, s2}
			combos = append(combos, newCombo)
		}
	}
	return combos
}
