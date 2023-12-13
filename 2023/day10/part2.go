package day10

import (
	"fmt"
	"strings"
)

type point struct {
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
	expanded = markEscapePoints(expanded)
	changed := 0
	for {
		expanded, changed = markConnectedPoints(expanded)
		if changed == 0 {
			break
		}
	}

	fmt.Println()
	for _, line := range expanded {
		fmt.Println(line)
	}

	return countFromUnexpandedPoints(inputWithOnlyLoop, expanded)
}

func mapLoopIds(loopIds map[point]int, currentTiles []tile, input []string) map[point]int {
	for {
		newCurrentTiles := []tile{}
		for _, tile := range currentTiles {
			newTile := getNewTile(tile, input)
			newPoint := point{newTile.row, newTile.col}

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
			p := point{r, c}
			newChar := "."

			if loopIds[p] == 1 {
				newChar = input[r][c : c+1]
			}

			startingPoint := point{startingTiles[0].row, startingTiles[0].col}
			if p == startingPoint {
				newChar = getNewStartingChar(startingTiles[0].next + startingTiles[1].next)
			}

			newRow += newChar
		}
		newInput[r] = newRow
	}
	return newInput
}

func getNewStartingChar(directions direction) string {
	switch directions {
	case n + e, e + n:
		return "L"
	case n + w, w + n:
		return "J"
	case n + s, s + n:
		return "|"
	case e + w, w + e:
		return "-"
	case e + s, s + e:
		return "F"
	case w + s, s + w:
		return "7"
	default:
		panic(fmt.Sprintf("couldn't find starting char from directions: %v\n", directions))
	}
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

func markEscapePoints(expanded []string) []string {
	for _, row := range []int{0, len(expanded) - 1} {
		for col := 0; col < len(expanded[0])-1; col++ {
			if expanded[row][col] == '.' {
				expanded = markPointOnImage(point{row, col}, expanded, "O")
			}
		}
	}

	for _, col := range []int{0, len(expanded[0]) - 1} {
		for row := 0; row < len(expanded)-1; row++ {
			if expanded[row][col] == '.' {
				expanded = markPointOnImage(point{row, col}, expanded, "O")
			}
		}
	}
	return expanded
}

func markPointOnImage(p point, image []string, char string) []string {
	row := image[p.row]
	image[p.row] = row[:p.col] + char + row[p.col+1:]
	return image
}

func markConnectedPoints(image []string) ([]string, int) {
	newImage := make([]string, len(image))
	changed := 0
	for r, row := range image {
		newRow := ""
		for c, _ := range row {
			char := string(image[r][c])
			if char != "." {
				newRow += char
				continue
			}

			nP := point{r - 1, c}
			eP := point{r, c + 1}
			wP := point{r, c - 1}
			sP := point{r + 1, c}

			if r > 0 {
				char = updateCharFromNeighbor(char, nP, image)
			}

			if c < len(row)-1 {
				char = updateCharFromNeighbor(char, eP, image)
			}

			if c > 0 {
				char = updateCharFromNeighbor(char, wP, image)
			}

			if r < len(image)-1 {
				char = updateCharFromNeighbor(char, sP, image)
			}

			if char == "O" {
				changed++
			}

			newRow += char
		}
		newImage[r] = newRow
	}
	return newImage, changed
}

func updateCharFromNeighbor(char string, p point, image []string) string {
	neighborChar := string(image[p.row][p.col])
	if neighborChar == "O" {
		char = neighborChar
	}
	return char
}

func countFromUnexpandedPoints(original []string, expanded []string) int {
	enclosedCount := 0
	for r, row := range original {
		for c, _ := range row {
			char := original[r][c]
			if char == '.' {
				expandedP := point{r * 2, c * 2}
				if expanded[expandedP.row][expandedP.col] == '.' {
					enclosedCount++
				}
			}
		}
	}
	return enclosedCount
}
