package day10

import "fmt"

func runB(input []string) int {
	// fmt.Println("\n-----------------------------------------------")
	startingTiles := findStartTiles(input)
	startingKey := startingTiles[0].row*1000 + startingTiles[0].col

	// key = row * 1000 + col
	// value = distance
	loopIds := map[int]int{}
	loopIds[startingKey] = 1
	loopIds = mapLoopIds(loopIds, startingTiles, input)
	inputWithOnlyLoop := redrawInputWithOnlyLoop(startingTiles, loopIds, input)

	enclosedCount := 0
	for r, row := range input {
		for c, _ := range row {
			if isEnclosed(r, c, loopIds, inputWithOnlyLoop) {
				enclosedCount++
			}
		}
	}
	return enclosedCount
}

func mapLoopIds(loopIds map[int]int, currentTiles []tile, input []string) map[int]int {
	for {
		newCurrentTiles := []tile{}
		for _, tile := range currentTiles {
			newTile := getNewTile(tile, input)
			newTileKey := newTile.row*1000 + newTile.col

			if loopIds[newTileKey] == 1 {
				return loopIds
			}
			loopIds[newTileKey] = 1
			newCurrentTiles = append(newCurrentTiles, newTile)
		}
		currentTiles = newCurrentTiles
	}
}

func redrawInputWithOnlyLoop(startingTiles []tile, loopIds map[int]int, input []string) []string {
	newInput := make([]string, len(input))
	for r := 0; r < len(input); r++ {
		newRow := ""
		for c := 0; c < len(input[0]); c++ {
			key := r*1000 + c
			newChar := "."

			if loopIds[key] == 1 {
				newChar = input[r][c : c+1]
			}

			if key == startingTiles[0].row*1000+startingTiles[0].col {
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

func isEnclosed(r int, c int, loopIds map[int]int, inputWithOnlyLoop []string) bool {
	currentKeys := []int{r*1000 + c}
	previousKeys := map[int]int{}
	previousKeys[r*1000+c] = 1
	maxRow := len(inputWithOnlyLoop) - 1
	maxCol := len(inputWithOnlyLoop[0]) - 1

	for {
		if hasAnyEscapes(currentKeys, maxRow, maxCol) {
			return false
		}

		currentKeys = expand(currentKeys, previousKeys, loopIds, inputWithOnlyLoop)
		previousKeys = addToMap(previousKeys, currentKeys)
	}
}

func hasAnyEscapes(keys []int, maxRow int, maxCol int) bool {
	for _, k := range keys {
		row := k / 1000
		col := k % 1000
		if row == 0 ||
			row == maxRow ||
			col == 0 ||
			col == maxCol {
			return true
		}
	}
	return false
}

func expand(keys []int, previousKeys map[int]int, loopIds map[int]int, input []string) []int {
	newKeys := []int{}
	maxRow := len(input) - 1
	maxCol := len(input[0]) - 1
	for _, k := range keys {
		row := k / 1000
		col := k % 1000

		nKey := (row-1)*1000 + col
		eKey := row*1000 + col + 1
		wKey := row*1000 + col - 1
		sKey := (row+1)*1000 + col
		possibleKeys := []int{}

		if row > 0 {
			possibleKeys = append(possibleKeys, nKey)
		}

		if col < maxCol {
			possibleKeys = append(possibleKeys, eKey)
		}

		if col > 0 {
			possibleKeys = append(possibleKeys, wKey)
		}

		if row < maxRow {
			possibleKeys = append(possibleKeys, sKey)
		}

		possibleKeys = filterOutPrevious(possibleKeys, previousKeys)
		keysToAdd := filterOutBlocks(possibleKeys, loopIds, input)
		newKeys = append(newKeys, keysToAdd...)
		previousKeys = addToMap(previousKeys, keysToAdd)
	}
	return newKeys
}

func filterOutPrevious(keys []int, previous map[int]int) []int {
	filtered := []int{}
	for _, k := range keys {
		if previous[k] == 1 {
			continue
		}
		filtered = append(filtered, k)
	}
	return filtered
}

func filterOutBlocks(keys []int, loopIds map[int]int, input []string) []int {
	filtered := []int{}
	for _, k := range keys {
		if loopIds[k] == 1 {
			row := k / 1000
			col := k % 1000
			loopChar := input[row][col : col+1]

		}
	}
}

func addToMap(keyMap map[int]int, keys []int) map[int]int {
	for _, k := range keys {
		keyMap[k] = 1
	}
	return keyMap
}
