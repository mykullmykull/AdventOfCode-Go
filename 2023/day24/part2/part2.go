package day24Part2

import "fmt"

func runB(input []string) float64 {
	stones := parseStones(input)
	t := thrown{}.resetRanges()
	hasValidRanges := false
	dDiff := 0.0
	for {
		if t.dDiff != dDiff {
			dDiff = t.dDiff
			fmt.Printf("dDiff: %v\n", dDiff)
		}
		for _, s := range stones {
			t = t.narrowThrownRanges(s)
			if !t.rangeX.isValid() || !t.rangeY.isValid() || !t.rangeZ.isValid() {
				t = t.getNextDs()
				t = t.resetRanges()
				break
			}
			hasValidRanges = true
		}
		if hasValidRanges && t.rangeX.isFinished() && t.rangeY.isFinished() && t.rangeZ.isFinished() {
			break
		}
	}
	fmt.Printf("found valid ranges: %v\n", t)
	return t.rangeX.min + t.rangeY.min + t.rangeZ.min
}
