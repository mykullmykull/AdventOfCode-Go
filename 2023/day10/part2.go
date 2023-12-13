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

	for r, row := range input {
		for c, _ := range row {
			char := inputWithOnlyLoop[r][c]
			if char != '.' {
				continue
			}
			expandedP := rowCol{r * 2, c * 2}
			expandedChar := expanded[expandedP.row][expandedP.col]
			if expandedChar != '.' {
				continue
			}

			fmt.Printf("\nexpandedP.row: %d; expandedP.col: %d, char: %c\n", expandedP.row, expandedP.col, expandedChar)

			// need a copy that we can mark up for each point
			copyExpanded := make([]string, len(expanded))
			copy(copyExpanded, expanded)
			stillUnknown := 0
			if isEnclosed(expandedP, copyExpanded) {
				fmt.Println("enclosed")
				markPointsOnImage([]rowCol{{expandedP.row, expandedP.col}}, expanded, "I")
			} else {
				fmt.Println("not enclosed")
				markPointsOnImage([]rowCol{{expandedP.row, expandedP.col}}, expanded, "O")
			}

			fmt.Println("updating connected")
			changed := 0
			for {
				expanded, stillUnknown, changed = markConnectedPoints(expandedP, expanded)
				if changed == 0 {
					break
				}
			}

			fmt.Println()
			for _, line := range expanded {
				fmt.Println(line)
			}

			if stillUnknown == 0 {
				return countFromUnexpandedPoints(inputWithOnlyLoop, expanded)
			}
		}
	}

	return countFromUnexpandedPoints(inputWithOnlyLoop, expanded)
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
	if hasStraightShotOut(p, image) {
		return false
	}

	currentPoints := []rowCol{p}
	for {
		// fmt.Printf("  p: %v, current points: %d\n", p, len(currentPoints))
		// for _, line := range image {
		// 	fmt.Println(line)
		// }

		for _, cp := range currentPoints {
			if hasStraightShotOut(cp, image) {
				return false
			}
		}

		if hasEscaped(currentPoints, image) {
			return false
		}

		image = markPointsOnImage(currentPoints, image, "+")

		currentPoints = expand(currentPoints, image)
		if len(currentPoints) == 0 {
			return true
		}
	}
}

func hasStraightShotOut(p rowCol, image []string) bool {
	return canEscapeStraight(p, image, n) ||
		canEscapeStraight(p, image, e) ||
		canEscapeStraight(p, image, w) ||
		canEscapeStraight(p, image, s)
}

func canEscapeStraight(p rowCol, image []string, direction direction) bool {
	rowChange := 0
	colChange := 0

	switch direction {
	case n:
		rowChange = -1
	case e:
		colChange = 1
	case w:
		colChange = -1
	case s:
		rowChange = 1
	}

	for {
		if hasEscaped([]rowCol{p}, image) {
			return true
		}

		if image[p.row][p.col] != '.' {
			return false
		}

		p = rowCol{p.row + rowChange, p.col + colChange}
	}
}

func hasEscaped(points []rowCol, image []string) bool {
	maxRow := len(image) - 1
	maxCol := len(image[0]) - 1
	for _, p := range points {
		char := string(image[p.row][p.col])
		if char == "O" {
			return true
		}

		if char == "I" {
			return false
		}

		if char != "." {
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

func markPointsOnImage(points []rowCol, image []string, char string) []string {
	for _, p := range points {
		row := image[p.row]
		image[p.row] = row[:p.col] + char + row[p.col+1:]
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

func markConnectedPoints(p rowCol, image []string) ([]string, int, int) {
	newImage := make([]string, len(image))
	unknown := 0
	newEnclosed := 0
	changed := 0
	for r, row := range image {
		newRow := ""
		for c, _ := range row {
			char := string(image[r][c])
			if char != "." {
				newRow += char
				continue
			}

			nP := rowCol{r - 1, c}
			eP := rowCol{r, c + 1}
			wP := rowCol{r, c - 1}
			sP := rowCol{r + 1, c}

			if r > 0 {
				nChar := string(image[nP.row][nP.col])
				if strings.Contains("IO", nChar) {
					char = nChar
				}
			}

			if c < len(row)-1 {
				eChar := string(image[eP.row][eP.col])
				if strings.Contains("IO", eChar) {
					char = eChar
				}
			}

			if c > 0 {
				wChar := string(image[wP.row][wP.col])
				if strings.Contains("IO", wChar) {
					char = wChar
				}
			}

			if r < len(image)-1 {
				sChar := string(image[sP.row][sP.col])
				if strings.Contains("IO", sChar) {
					char = sChar
				}
			}

			if char == "I" {
				newEnclosed++
			}

			if char == "." {
				unknown++
			}

			if strings.Contains("IO", char) {
				changed++
			}

			newRow += char
		}
		newImage[r] = newRow
	}
	return newImage, unknown, changed
}

func countFromUnexpandedPoints(original []string, expanded []string) int {
	enclosedCount := 0
	for r, row := range original {
		for c, _ := range row {
			char := original[r][c]
			if char == '.' {
				expandedP := rowCol{r * 2, c * 2}
				if expanded[expandedP.row][expandedP.col] == 'I' {
					enclosedCount++
				}
			}
		}
	}
	return enclosedCount
}
