package day18

import (
	"fmt"
	"goAoc2023/utils"
	"strings"
)

func runB(input []string) int64 {
	count := int64(0)
	horizontals := []horizontal{}
	row := 0
	col := 0
	for _, line := range input {
		split := strings.Split(line, " ")
		hex := split[2]
		hex = strings.Replace(hex, "(", "", -1)
		hex = strings.Replace(hex, ")", "", -1)
		hex = strings.Replace(hex, "#", "", -1)
		direction := translateDirection(utils.Atoi(string(hex[len(hex)-1])))
		distance := translateDistance(hex[:len(hex)-1])
		dRow, dCol := parseDirection(direction)
		h := horizontal{}
		if dCol == 1 {
			h = horizontal{
				row: row,
				min: col,
				max: col + distance,
			}
		}
		if dCol == -1 {
			h = horizontal{
				row: row,
				min: col - distance,
				max: col,
			}
		}
		if h.width() > 1 {
			horizontals = append(horizontals, h)
		}
		row += dRow * distance
		col += dCol * distance
	}
	count = countFromHorizontals(horizontals)
	return count
}

func translateDirection(d int) string {
	switch d {
	case 0:
		return "R"
	case 1:
		return "D"
	case 2:
		return "L"
	case 3:
		return "U"
	}
	panic("invalid direction")
}

func translateDistance(hex string) int {
	dist := 0
	for _, r := range hex {
		if dist == 0 {
			dist = hexDigitToInt(r)
			continue
		}
		dist = dist*16 + hexDigitToInt(r)
	}
	return dist
}

func hexDigitToInt(r rune) int {
	switch r {
	case 'a':
		return 10
	case 'b':
		return 11
	case 'c':
		return 12
	case 'd':
		return 13
	case 'e':
		return 14
	case 'f':
		return 15
	default:
		return utils.Atoi(string(r))
	}
}

func countFromHorizontals(horizontals []horizontal) int64 {
	count := int64(0)
	activeHorizontals := []horizontal{}
	minRow, maxRow := getRowBounds(horizontals)
	for r := minRow; r <= maxRow; r++ {
		newHorizontals := getNewHorizontals(r, horizontals)
		// debug0 := len(newHorizontals) > 0
		// debug1 := len(newHorizontals) > 0
		debug0 := false
		debug1 := false

		log("", debug1)
		log(fmt.Sprintf("\nrow: %d of %d from count: %d", r, maxRow, count), debug0)
		adderNewHorizontals, subtractorNewHorizontals := classifyNewHorizontals(activeHorizontals, newHorizontals)
		if len(newHorizontals) > 0 {
			if debug1 {
				if len(adderNewHorizontals) > 0 {
					printHorizontals(adderNewHorizontals, false)
				}
				if len(subtractorNewHorizontals) > 0 {
					printHorizontals(subtractorNewHorizontals, false)
				}
			}
			activeHorizontals = addNewHorizontals(activeHorizontals, adderNewHorizontals)
		}
		if debug1 && len(adderNewHorizontals) > 0 {
			printHorizontals(activeHorizontals, false)
		}
		for _, h := range activeHorizontals {
			if h.width() > 0 {
				count += int64(h.width())
			}
		}
		activeHorizontals = subtractNewHorizontals(activeHorizontals, subtractorNewHorizontals)
		if debug1 && len(subtractorNewHorizontals) > 0 {
			println()
			println("  active after subtracting")
			printHorizontals(activeHorizontals, false)
		}
	}
	return count
}

func getRowBounds(horizontals []horizontal) (int, int) {
	minRow := int(^uint(0) >> 1)
	maxRow := -1 * int(^uint(0)>>1)
	for _, h := range horizontals {
		if h.row < minRow {
			minRow = h.row
		}

		if h.row > maxRow {
			maxRow = h.row
		}
	}
	return minRow, maxRow
}

func getNewHorizontals(row int, horizontals []horizontal) []horizontal {
	newHorizontals := []horizontal{}
	for _, h := range horizontals {
		if h.row == row {
			newHorizontals = append(newHorizontals, h)
		}
	}
	return newHorizontals
}
func classifyNewHorizontals(activeHorizontals []horizontal, newHorizontals []horizontal) ([]horizontal, []horizontal) {
	debug := false
	adderNewHorizontals := []horizontal{}
	subtractorNewHorizontals := []horizontal{}

	for _, new := range newHorizontals {
		classification := ""
		for _, active := range activeHorizontals {
			if new.min > active.max || new.max < active.min {
				continue
			}

			if new.min > active.min && new.min < active.max && new.max < active.max && new.max > active.min {
				subtractorNewHorizontals = append(subtractorNewHorizontals, new)
				classification = "splits"
				break
			}

			if new.max == active.min {
				adderNewHorizontals = append(adderNewHorizontals, new)
				classification = "expands left edge"
				break
			}

			if new.max == active.max {
				subtractorNewHorizontals = append(subtractorNewHorizontals, new)
				classification = "shrinks right edge"
				break
			}

			if new.min == active.max {
				adderNewHorizontals = append(adderNewHorizontals, new)
				classification = "expands right edge"
				break
			}

			if new.min == active.min {
				subtractorNewHorizontals = append(subtractorNewHorizontals, new)
				classification = "shrinks left edge"
				break
			}

			panic(fmt.Sprintf("  new: %s new overlaps but doesn't align with active: %s", new.toStringWithRow(), active.toStringWithRow()))
		}

		if classification == "" {
			// brand new horizontal
			adderNewHorizontals = append(adderNewHorizontals, new)
		}
		log(fmt.Sprintf("  %s", classification), debug)
	}
	return adderNewHorizontals, subtractorNewHorizontals
}

func addNewHorizontals(activeHorizontals []horizontal, newHorizontals []horizontal) []horizontal {
	for _, new := range newHorizontals {
		expandingLeftEdge := -1
		expandingRightEdge := -1

		for i, active := range activeHorizontals {
			if new.max == active.min {
				expandingLeftEdge = i
				continue
			}
			if new.min == active.max {
				expandingRightEdge = i
				continue
			}
		}

		// brand new horizontal
		if expandingLeftEdge < 0 && expandingRightEdge < 0 {
			activeHorizontals = append(activeHorizontals, new)
			continue
		}

		// merges two horizontals
		if expandingLeftEdge >= 0 && expandingRightEdge >= 0 {
			merged := horizontal{
				min: activeHorizontals[expandingRightEdge].min,
				max: activeHorizontals[expandingLeftEdge].max,
			}
			activeHorizontals = removeFromSpan(activeHorizontals, expandingRightEdge, expandingLeftEdge)
			activeHorizontals = append(activeHorizontals, merged)
			continue
		}

		// expands left edge
		if expandingLeftEdge >= 0 {
			activeHorizontals[expandingLeftEdge].min = new.min
			continue
		}

		// expands right edge
		if expandingRightEdge >= 0 {
			activeHorizontals[expandingRightEdge].max = new.max
			continue
		}
	}
	return activeHorizontals
}

func subtractNewHorizontals(activeHorizontals []horizontal, newHorizontals []horizontal) []horizontal {
	for _, new := range newHorizontals {
		shrinkingLeftEdge := -1
		shrinkingRightEdge := -1
		splitting := -1

		for i, active := range activeHorizontals {
			if new.min == active.min {
				shrinkingLeftEdge = i
			}

			if new.max == active.max {
				shrinkingRightEdge = i
			}

			if shrinkingLeftEdge >= 0 || shrinkingRightEdge >= 0 {
				continue
			}

			if new.min < active.max && new.max > active.min {
				splitting = i
			}
		}

		// removing horizontal
		if shrinkingLeftEdge >= 0 && shrinkingRightEdge >= 0 && shrinkingLeftEdge == shrinkingRightEdge {
			activeHorizontals = removeFromSpan(activeHorizontals, shrinkingLeftEdge, -1)
			continue
		}

		// splitting horizontal
		if splitting >= 0 {
			left := horizontal{min: activeHorizontals[splitting].min, max: new.min}
			activeHorizontals[splitting].min = new.max
			activeHorizontals = append(activeHorizontals, left)
			continue
		}

		// shrinking left edge
		if shrinkingLeftEdge >= 0 {
			activeHorizontals[shrinkingLeftEdge].min = new.max
			continue
		}

		// shrinking right edge {
		if shrinkingRightEdge >= 0 {
			activeHorizontals[shrinkingRightEdge].max = new.min
			continue
		}
	}
	return activeHorizontals
}

func hasOverlappintHorizontals(horizontals []horizontal) bool {
	for i, h := range horizontals {
		for j := i + 1; j < len(horizontals); j++ {
			if h.overlaps(horizontals[j]) {
				return true
			}
		}
	}
	return false
}

func removeFromSpan(horizontals []horizontal, i int, j int) []horizontal {
	newHorizontals := []horizontal{}
	for k, h := range horizontals {
		if k == i || k == j {
			continue
		}
		newHorizontals = append(newHorizontals, h)
	}

	return newHorizontals
}

// ---------------------------------------- debugging utils ----------------------------------------

func printHorizontals(horizontals []horizontal, withRow bool) {
	for _, this := range horizontals {
		if withRow {
			println("    " + this.toStringWithRow())
		} else {
			println("    " + this.toStringNoRow())
		}
	}
}

func printHorizontalsInRowOrder(horizontals []horizontal) {
	minRow, maxRow := getRowBounds(horizontals)
	for r := minRow; r <= maxRow; r++ {
		newHorizontals := getNewHorizontals(r, horizontals)
		if len(newHorizontals) == 0 {
			continue
		}
		printHorizontals(newHorizontals, true)
	}
}

func printHorizontalsAffectingError(horizontals []horizontal, errorRowNumber int) {
	foundHorizontalsByRow := map[int]horizontal{}
	pointsOfInterest := map[int]bool{}
	pointsToAdd := map[int]bool{}

	errorRow := horizontals[0]
	for _, h := range horizontals {
		if h.row == errorRowNumber {
			errorRow = h
			break
		}
	}

	pointsToAdd[errorRow.min] = true
	pointsToAdd[errorRow.max] = true
	for {
		if len(pointsToAdd) == 0 {
			break
		}
		for p, _ := range pointsToAdd {
			for _, h := range horizontals {
				if h.min == p {
					foundHorizontalsByRow[h.row] = h
					if _, ok := pointsOfInterest[h.max]; !ok {
						pointsToAdd[h.max] = true
					}
				}
				if h.max == p {
					foundHorizontalsByRow[h.row] = h
					if _, ok := pointsOfInterest[h.min]; !ok {
						pointsToAdd[h.min] = true
					}
				}
			}
			delete(pointsToAdd, p)
			pointsOfInterest[p] = true
		}
	}
	foundHorizontals := []horizontal{}
	for _, v := range foundHorizontalsByRow {
		foundHorizontals = append(foundHorizontals, v)
	}
	fmt.Printf("found horizontals %d out of %d\n", len(foundHorizontals), len(horizontals))
	printHorizontalsInRowOrder(foundHorizontals)
}
