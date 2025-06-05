package day

import "goAoc2021/utils"

func getBasinSize(input []string, lowPoint point) int {
	floor := make([]string, len(input))
	for i := range floor {
		floor[i] = input[i]
	}
	floor[lowPoint.row] = utils.ReplaceAtIndex(floor[lowPoint.row], lowPoint.col, "+")
	floor = markBasin(floor)
	basinSize := measureBasin(floor)
	return basinSize
}

func markBasin(floor []string) []string {
	hasChanged := true
	for hasChanged {
		hasChanged = false
		for r, row := range floor {
			for c, cell := range row {
				if cell == '+' || cell == '9' {
					continue
				}
				up, dn, lt, rt := '9', '9', '9', '9'
				if r > 0 {
					up = rune(floor[r-1][c])
				}
				if r < len(floor)-1 {
					dn = rune(floor[r+1][c])
				}
				if c > 0 {
					lt = rune(floor[r][c-1])
				}
				if c < len(floor[r])-1 {
					rt = rune(floor[r][c+1])
				}
				if up == '+' || dn == '+' || lt == '+' || rt == '+' {
					floor[r] = utils.ReplaceAtIndex(floor[r], c, "+")
					hasChanged = true
				}
			}
		}
	}
	return floor
}

func measureBasin(floor []string) int {
	basinSize := 0
	for _, row := range floor {
		for _, cell := range row {
			if cell == '+' {
				basinSize++
			}
		}
	}
	return basinSize
}
