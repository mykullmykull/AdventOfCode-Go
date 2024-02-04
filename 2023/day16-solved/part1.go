package day16

import "fmt"

func runA(grid []string) int {
	gridMap := parseGridMap(grid)
	beams := beam{point{0, -1}, rt}.advance(gridMap)
	gridMap[point{0, 0}] = gridMap[point{0, 0}].addBeams(beams)
	for {
		gridMap, beams = update(gridMap, beams)
		if len(beams) == 0 {
			break
		}
	}

	return countEnergizedGridPoints(gridMap)
}

func parseGridMap(grid []string) map[point]gridPoint {
	gridMap := make(map[point]gridPoint)
	for r, row := range grid {
		for c, char := range row {
			p := point{r, c}
			if char == '.' {
				gridMap[p] = gridPoint{false, false, false, false, empty}
			} else if char == '|' {
				gridMap[p] = gridPoint{false, false, false, false, vSplitter}
			} else if char == '-' {
				gridMap[p] = gridPoint{false, false, false, false, hSplitter}
			} else if char == '/' {
				gridMap[p] = gridPoint{false, false, false, false, aMirror}
			} else if char == '\\' {
				gridMap[p] = gridPoint{false, false, false, false, bMirror}
			}
		}
	}
	return gridMap
}

func update(gridMap map[point]gridPoint, beams []beam) (map[point]gridPoint, []beam) {
	newBeams := []beam{}
	for _, b := range beams {
		newBeams = append(newBeams, b.advance(gridMap)...)
	}
	for _, b := range newBeams {
		gridMap[b.location] = gridMap[b.location].addBeams([]beam{b})
	}
	return gridMap, newBeams
}

func printGridMap(gridMap map[point]gridPoint) {
	fmt.Println()
	row := 0
	col := -1
	for {
		col++
		p := point{row, col}
		gp, ok := gridMap[p]
		if !ok && col == 0 {
			break
		}

		if !ok {
			col = -1
			row++
			fmt.Println()
			continue
		}

		if gp.reflector == empty {
			if gp.uBeam || gp.lBeam || gp.dBeam || gp.rBeam {
				fmt.Print("#")
				continue
			}

			fmt.Print(".")
			continue
		}

		fmt.Print(string(gridMap[p].reflector))
	}
}

func countEnergizedGridPoints(gridMap map[point]gridPoint) int {
	count := 0
	for _, gp := range gridMap {
		if gp.uBeam || gp.lBeam || gp.dBeam || gp.rBeam {
			count++
		}
	}
	return count
}

func log(msg string, shouldPrint bool) {
	if shouldPrint {
		fmt.Println(msg)
	}
}
