package day10

import (
	"fmt"
	"strings"
)

type rowCol struct {
	row int
	col int
}

func runB(input []string) int {
	startingTiles := findStartTiles(input)
	starting := point{row: startingTiles[0].row, col: startingTiles[0].col}

	// This is a map so that it is quickly searchable. All values should be 1.
	loopPoints := map[point]int{}
	loopPoints[starting] = 1
	loopPoints = mapLoopIds(loopPoints, startingTiles, input)
	inputWithOnlyLoop := redrawInputWithOnlyLoop(startingTiles, loopPoints, input)
	fmt.Println()
	for _, line := range inputWithOnlyLoop {
		fmt.Println(line)
	}

	expanded := expandRows(inputWithOnlyLoop)
	expanded = expandCols(expanded)

	enclosedCount := 0
	for r, row := range input {
		for c, _ := range row {
			char := inputWithOnlyLoop[r][c]
			if char != '.' {
				continue
			}
			fmt.Printf("\nr: %d; c: %d\n", r, c)
			expandedP := rowCol{r * 2, c * 2}

			// need a copy that we can mark up for each point
			copyExpanded := make([]string, len(expanded))
			copy(copyExpanded, expanded)
			if isEnclosed(expandedP, copyExpanded) {
				enclosedCount++
			}
			fmt.Printf("enclosed: %d\n", enclosedCount)
		}
	}

	return enclosedCount
}

func mapLoopIds(loopIds map[point]int, currentTiles []tile, input []string) map[point]int {
	for {
		newCurrentTiles := []tile{}
		for _, tile := range currentTiles {
			newTile := getNewTile(tile, input)
			newPoint := point{newTile.row, newTile.col, 0, 0}

			if loopIds[newPoint] == 1 {
				return loopIds
			}
			loopIds[newPoint] = 1
			newCurrentTiles = append(newCurrentTiles, newTile)
		}
		currentTiles = newCurrentTiles
	}
}

func redrawInputWithOnlyLoop(startingTiles []tile, loopIds map[point]int, input []string) []string {
	newInput := make([]string, len(input))
	for r := 0; r < len(input); r++ {
		newRow := ""
		for c := 0; c < len(input[0]); c++ {
			p := point{r, c, 0, 0}
			newChar := "."

			if loopIds[p] == 1 {
				newChar = input[r][c : c+1]
			}

			startingPoint := point{startingTiles[0].row, startingTiles[0].col, 0, 0}
			if p == startingPoint {
				newChar = getNewStartingChar(startingTiles[0].next + startingTiles[1].next)
			}

			newRow += newChar
		}
		newInput[r] = newRow
	}
	return newInput
}

func expandRows(input []string) []string {
	newImage := []string{}
	for _, line := range input {
		newImage = append(newImage, line)
		newLine := ""
		for _, c := range line {
			if strings.Contains("|7F", string(c)) {
				newLine += "|"
			} else {
				newLine += "."
			}
		}
		newImage = append(newImage, newLine)
	}
	return newImage
}

func expandCols(input []string) []string {
	newImage := []string{}
	for _, line := range input {
		newLine := ""
		for _, c := range line {
			newLine += string(c)
			if strings.Contains("-LF", string(c)) {
				newLine += "-"
			} else {
				newLine += "."
			}
		}
		newImage = append(newImage, newLine)
	}
	return newImage
}

func isEnclosed(p rowCol, image []string) bool {
	currentPoints := []rowCol{p}
	for {
		if hasStraightShotOut(p, image) {
			fmt.Printf("has straight shot\n")
			return false
		}

		if hasEscaped(currentPoints, image) {
			return false
		}

		image = markPointsOnImage(currentPoints, image)

		currentPoints = expand(currentPoints, image)
		if len(currentPoints) == 0 {
			return true
		}
	}
}

func hasStraightShotOut(p rowCol, image []string) bool {
	// try north
	escapes := true
	for r := p.row; r >= 0; r-- {
		if image[p.row][p.col] != '.' {
			fmt.Println("not north")
			escapes = false
			break
		}
	}
	if escapes {
		return true
	}

	// try south
	escapes = true
	for r := p.row; r < len(image); r++ {
		if image[p.row][p.col] != '.' {
			fmt.Println("not south")
			escapes = false
			break
		}
	}
	if escapes {
		return true
	}

	// try east
	escapes = true
	for r := p.col; r < len(image[0]); r++ {
		if image[p.row][p.col] != '.' {
			fmt.Println("not east")
			escapes = false
			break
		}
	}
	if escapes {
		return true
	}

	// try west
	for r := p.col; r >= 0; r-- {
		if image[p.row][p.col] != '.' {
			fmt.Println("not west")
			escapes = false
			break
		}
	}
	return escapes
}

func hasEscaped(points []rowCol, image []string) bool {
	maxRow := len(image) - 1
	maxCol := len(image[0]) - 1
	for _, p := range points {
		if image[p.row][p.col] != '.' {
			continue
		}
		if p.row == 0 ||
			p.row == maxRow ||
			p.col == 0 ||
			p.col == maxCol {
			return true
		}
	}
	return false
}

func markPointsOnImage(points []rowCol, image []string) []string {
	for _, p := range points {
		row := image[p.row]
		image[p.row] = row[:p.col] + "+" + row[p.col+1:]
	}
	return image
}

func expand(points []rowCol, image []string) []rowCol {
	newPoints := []rowCol{}
	maxRow := len(image) - 1
	maxCol := len(image[0]) - 1
	for _, p := range points {
		nPoint := rowCol{p.row - 1, p.col}
		ePoint := rowCol{p.row, p.col + 1}
		wPoint := rowCol{p.row, p.col - 1}
		sPoint := rowCol{p.row + 1, p.col}

		if p.row > 0 && isNew(nPoint, image) {
			newPoints = append(newPoints, nPoint)
		}

		if p.col < maxCol && isNew(ePoint, image) {
			newPoints = append(newPoints, ePoint)
		}

		if p.col > 0 && isNew(wPoint, image) {
			newPoints = append(newPoints, wPoint)
		}

		if p.row < maxRow && isNew(sPoint, image) {
			newPoints = append(newPoints, sPoint)
		}
	}
	return newPoints
}

func isNew(p rowCol, image []string) bool {
	return image[p.row][p.col] == '.'
}
