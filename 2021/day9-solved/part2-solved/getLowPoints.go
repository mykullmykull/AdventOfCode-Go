package day

import "goAoc2021/utils"

func getLowPoints(input []string) []point {
	lowPoints := []point{}
	for r, row := range input {
		for c := range row {
			if isLowPoint(input, r, c) {
				lowPoints = append(lowPoints, point{row: r, col: c})
			}
		}
	}
	return lowPoints
}

func isLowPoint(input []string, row int, col int) bool {
	current := utils.Btoi(input[row][col])
	up, dn, lt, rt := 9, 9, 9, 9

	// Check up
	if row > 0 {
		up = utils.Btoi(input[row-1][col])
	}

	// Check down
	if row < len(input)-1 {
		dn = utils.Btoi(input[row+1][col])
	}

	// Check left
	if col > 0 {
		lt = utils.Btoi(input[row][col-1])
	}

	// Check right
	if col < len(input[row])-1 {
		rt = utils.Btoi(input[row][col+1])
	}

	return current < up && current < dn && current < lt && current < rt
}
