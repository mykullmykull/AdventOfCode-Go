package day

import (
	"fmt"
	"goAoc2021/utils"
)

func (c *cave) expandGrid() {
	oldGrid := c.grid
	expanded := make([]string, len(oldGrid)*5)
	for row := 0; row < len(oldGrid)*5; row++ {
		newLine := ""
		for col := 0; col < len(oldGrid[0])*5; col++ {
			if row < len(oldGrid) && col < len(oldGrid[0]) {
				newLine += oldGrid[row][col : col+1]
				continue
			}
			previousRisk := -1
			if col < len(oldGrid) {
				previousRisk = utils.Atoi(string(expanded[row-len(oldGrid)][col]))
			} else {
				previousRisk = utils.Atoi(string(newLine[col-len(oldGrid[0])]))
			}

			newRisk := previousRisk + 1
			if newRisk > 9 {
				newRisk = 1
			}
			newLine += fmt.Sprintf("%d", newRisk)
		}
		expanded[row] = newLine
	}
	c.grid = expanded
}
