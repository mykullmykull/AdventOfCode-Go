package day10

import (
	"fmt"
)

type point struct {
	row   int
	col   int
	vHalf int
	hHalf int
}

type move struct {
	oldP   point
	newP   point
	moving direction
}

func runB_failed(input []string) int {
	// fmt.Println("\n-----------------------------------------------")
	startingTiles := findStartTiles(input)
	starting := point{row: startingTiles[0].row, col: startingTiles[0].col}

	// This is a map so that it is quickly searchable. All values should be 1.
	loopPoints := map[point]int{}
	loopPoints[starting] = 1
	loopPoints = mapLoopIds(loopPoints, startingTiles, input)
	inputWithOnlyLoop := redrawInputWithOnlyLoop(startingTiles, loopPoints, input)

	for _, line := range inputWithOnlyLoop {
		fmt.Println(line)
	}

	enclosedCount := 0
	for r, row := range input {
		for c, _ := range row {
			p := point{r, c, 0, 0}
			char := inputWithOnlyLoop[r][c]
			if char != '.' {
				continue
			}

			if r == 4 && c == 7 {
				fmt.Printf("\nis r: %d, c: %d enclosed?\n", r, c)
			}
			if isEnclosed_failed(p, loopPoints, inputWithOnlyLoop) {
				enclosedCount++
				fmt.Printf("  enclosed count: %d\n", enclosedCount)
			}
		}
	}
	fmt.Printf("final enclosed count: %d\n", enclosedCount)
	return enclosedCount
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

func isEnclosed_failed(p point, loopIds map[point]int, inputWithOnlyLoop []string) bool {
	currentPoints := []point{p}
	previousPoints := map[point]int{}
	previousPoints = addToMap(previousPoints, currentPoints)

	for {
		if hasAnyEscapes_failed(currentPoints, inputWithOnlyLoop) {
			if p.row == 4 && p.col == 7 {
				fmt.Printf("  no, escaped\n")
			}
			return false
		}

		currentPoints = expand_failed(currentPoints, previousPoints, loopIds, inputWithOnlyLoop)
		if p.row == 4 && p.col == 7 {
			// fmt.Printf("  current points: %v\n", currentPoints)
		}

		if len(currentPoints) == 0 {
			return true
		}
		previousPoints = addToMap(previousPoints, currentPoints)
	}
}

func hasAnyEscapes_failed(points []point, input []string) bool {
	maxRow := len(input) - 1
	maxCol := len(input[0]) - 1
	for _, p := range points {
		if input[p.row][p.col] != '.' {
			continue
		}
		if p.row == 0 ||
			p.row == maxRow ||
			p.col == 0 ||
			p.col == maxCol {
			fmt.Printf("  excape p: %v\n", p)
			return true
		}
	}
	return false
}

func expand_failed(points []point, previousPoints map[point]int, loopPoints map[point]int, input []string) []point {
	newPoints := []point{}
	maxRow := len(input) - 1
	maxCol := len(input[0]) - 1
	for _, p := range points {
		pointsToAdd := []point{}

		nMove := move{p, point{p.row - 1, p.col, 0, p.hHalf}, n}
		eMove := move{p, point{p.row, p.col + 1, p.vHalf, 0}, e}
		wMove := move{p, point{p.row, p.col - 1, p.vHalf, 0}, w}
		sMove := move{p, point{p.row + 1, p.col, 0, p.hHalf}, s}

		if p.row > 0 &&
			isNew_failed(nMove, previousPoints) &&
			isPassable(nMove, input) {
			newPoint := squeezeIfNeeded(nMove, input)
			fmt.Printf("  nMove.oldP %v => %v\n", nMove.oldP, newPoint)
			pointsToAdd = append(pointsToAdd, newPoint)
		}

		if p.col < maxCol &&
			isNew_failed(eMove, previousPoints) &&
			isPassable(eMove, input) {
			newPoint := squeezeIfNeeded(eMove, input)
			fmt.Printf("  eMove.oldP %v => %v\n", eMove.oldP, newPoint)
			pointsToAdd = append(pointsToAdd, newPoint)
		}

		if p.col > 0 &&
			isNew_failed(wMove, previousPoints) &&
			isPassable(wMove, input) {
			newPoint := squeezeIfNeeded(wMove, input)
			fmt.Printf("  wMove.oldP %v => %v\n", wMove.oldP, newPoint)
			pointsToAdd = append(pointsToAdd, newPoint)
		}

		if p.row < maxRow &&
			isNew_failed(sMove, previousPoints) &&
			isPassable(sMove, input) {
			newPoint := squeezeIfNeeded(sMove, input)
			fmt.Printf("  sMove.oldP %v => %v\n", sMove.oldP, newPoint)
			pointsToAdd = append(pointsToAdd, newPoint)
		}

		newPoints = append(newPoints, pointsToAdd...)
		if len(pointsToAdd) > 0 {
			fmt.Println()
		}
		previousPoints = addToMap(previousPoints, pointsToAdd)
	}
	return newPoints
}

func isNew_failed(m move, previous map[point]int) bool {
	m.newP.hHalf = 0
	m.newP.vHalf = 0
	return previous[m.newP] != 1
}

func isPassable(m move, input []string) bool {
	oldChar := input[m.oldP.row][m.oldP.col]
	newChar := input[m.newP.row][m.newP.col]

	switch m.moving {
	case n, s:
		if oldChar == '-' || newChar == '-' {
			return false
		}

		if newChar == '.' {
			return true
		}

		return canSqueeze(m, input)
	case e, w:
		if oldChar == '|' || newChar == '|' {
			return false
		}

		if newChar == '.' {
			return true
		}

		return canSqueeze(m, input)
	}

	panic(fmt.Sprintf("cannot verify passability of m: %v\n", m))
}

func canSqueeze(m move, input []string) bool {
	newChar := string(input[m.newP.row][m.newP.col])
	if newChar == "|" {
		return (m.moving == n || m.moving == s) && m.oldP.hHalf != 0
	}

	if newChar == "-" {
		return (m.moving == e || m.moving == w) && m.oldP.vHalf != 0
	}

	switch m.moving {
	case n:
		if m.newP.hHalf > 0 {
			return newChar != "F"
		}
		if m.newP.hHalf < 0 {
			return newChar != "7"
		}
		return canNewVSqueeze(m, input)
	case s:
		if m.newP.hHalf > 0 {
			return newChar != "L"
		}
		if m.newP.hHalf < 0 {
			return newChar != "J"
		}
		return canNewVSqueeze(m, input)
	case e:
		if m.newP.vHalf > 0 {
			return newChar != "J"
		}
		if m.newP.vHalf < 0 {
			return newChar != "7"
		}
		return canNewHSqueeze(m, input)
	case w:
		if m.newP.vHalf > 0 {
			return newChar != "L"
		}
		if m.newP.vHalf < 0 {
			return newChar != "F"
		}
		return canNewHSqueeze(m, input)
	}

	panic(fmt.Sprintf("cannot verify if m can squeeze. m: %v\n", m))
}

func canNewVSqueeze(m move, input []string) bool {
	return (m.newP.col == 0 || input[m.newP.row][m.newP.col-1] != '-') &&
		(m.newP.col == len(input[0])-1 || input[m.newP.row][m.newP.col+1] != '-')
}

func canNewHSqueeze(m move, input []string) bool {
	return (m.newP.row == 0 || input[m.newP.row-1][m.newP.col] != '|') &&
		(m.newP.row == len(input)-1 || input[m.newP.row+1][m.newP.col] != '|')
}

func addToMap(pointMap map[point]int, points []point) map[point]int {
	for _, p := range points {
		p.hHalf = 0
		p.vHalf = 0
		pointMap[p] = 1
	}
	return pointMap
}

func squeezeIfNeeded(m move, input []string) point {
	newChar := input[m.newP.row][m.newP.col]
	if m.newP.hHalf != 0 || m.newP.vHalf != 0 {
		if newChar == '.' {
			m.newP.vHalf = 0
			m.newP.hHalf = 0
		}
		return m.newP
	}

	if newChar == '.' {
		return m.newP
	}

	vHalf, hHalf := 1, 1

	switch newChar {
	case 'L':
		vHalf = -1
		hHalf = -1
	case 'F':
		hHalf = -1
	case 'J':
		vHalf = -1
	}

	switch m.moving {
	case n, s:
		m.newP.hHalf = hHalf
	case e, w:
		m.newP.vHalf = vHalf
	}

	return m.newP
}
