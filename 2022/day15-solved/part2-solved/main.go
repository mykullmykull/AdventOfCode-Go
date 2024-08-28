package day

func main(input []string, maxCoord int) int {
	sensors := parse(input)
	distressBeacon := point{col: -1}
	isCovered := false
	for row := 0; row <= maxCoord; row++ {
		if isRowEntirelyCovered(row, sensors, maxCoord) {
			continue
		}

		for col := 0; col <= maxCoord; col++ {
			p := point{col: col, row: row}
			isCovered = false
			for _, s := range sensors {
				if p.distanceFrom(s.p) <= s.distance {
					isCovered = true
					break
				}
			}
			if isCovered {
				continue
			}
			distressBeacon = p
			break
		}
		if !isCovered {
			break
		}
	}
	return distressBeacon.col*4000000 + distressBeacon.row
}

func isRowEntirelyCovered(row int, sensors []sensor, maxCoord int) bool {
	coverages := []rowCoverage{}
	for _, s := range sensors {
		min, max := s.rowMinMax(row, maxCoord)
		if max > min {
			coverages = append(coverages, rowCoverage{min: min, max: max})
		}
	}
	for {
		if len(coverages) == 1 {
			break
		}
		first := coverages[0]
		coverages = coverages[1:]
		isCombinable := false
		for x, c := range coverages {
			combined := first.combineIfOverlaps(c)
			if len(combined) == 1 {
				coverages[x] = combined[0]
				isCombinable = true
				break
			}
		}
		if isCombinable {
			continue
		}
		coverages = append(coverages, first)
		if len(coverages) == 2 {
			break
		}
	}
	return len(coverages) == 1
}
