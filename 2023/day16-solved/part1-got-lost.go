package day16

// import (
// 	"fmt"
// 	"goAoc2023/utils"
// )

// func runA_got_lost(grid []string) int {
// 	printGrid(grid)
// 	return countEnergized(grid)
// }

// func countEnergized(grid []string) int {
// 	energized := initiateEnergized(grid)
// 	beams := []beam{{point{0, 0}, rt}}
// 	for {
// 		fmt.Printf("len(beams): %d\n", len(beams))
// 		fmt.Printf("beams[0]: %v\n", beams[0])
// 		if len(beams) == 0 {
// 			break
// 		}
// 		beams, grid, energized = moveAndSplit(beams, grid, energized)
// 		printGrid(energized)
// 	}

// 	count := 0
// 	for _, row := range energized {
// 		for _, char := range row {
// 			if char != '.' {
// 				count++
// 			}
// 		}
// 	}
// 	return count
// }

// func initiateEnergized(grid []string) []string {
// 	energized := []string{}
// 	baseRow := ""
// 	for i := 0; i < len(grid[0]); i++ {
// 		baseRow += "."
// 	}
// 	for i := 0; i < len(grid); i++ {
// 		energized = append(energized, baseRow)
// 	}
// 	return energized
// }

// func moveAndSplit(beams []beam, grid []string, energized []string) ([]beam, []string, []string) {
// 	newBeams := []beam{}
// 	// examine the each beam and its current grid char
// 	// movers: . > < v ^ or digit mean that direction continues the same and grid char gets updated
// 	// mirrors: / \ mean that direction changes
// 	// spliters: - | mean that either direction continues or beam spawns two with directions perpendicular to current
// 	// each combo needs to update newBeams, grid, and energized
// 	for _, b := range beams {
// 		char := grid[b.location.row][b.location.col]
// 		// fmt.Printf("moveAndSplit char: %s\n", string(char))
// 		switch char {
// 		case '.', '>', '<', 'v', '^', '2', '3', '4', '5', '6', '7', '8', '9':
// 			newBeams, energized = updateNewBeamsEnergized(b, b.direction, newBeams, energized, grid)
// 			grid = updateGrid(grid, b.location, b.direction)
// 		case '-':
// 			switch b.direction {
// 			case up, dn:
// 				newBeams, energized = updateNewBeamsEnergized(b, lt, newBeams, energized, grid)
// 				newBeams, energized = updateNewBeamsEnergized(b, rt, newBeams, energized, grid)
// 			case lt, rt:
// 				newBeams, energized = updateNewBeamsEnergized(b, b.direction, newBeams, energized, grid)
// 			}
// 		case '|':
// 			switch b.direction {
// 			case lt, rt:
// 				newBeams, energized = updateNewBeamsEnergized(b, up, newBeams, energized, grid)
// 				newBeams, energized = updateNewBeamsEnergized(b, dn, newBeams, energized, grid)
// 			case up, dn:
// 				newBeams, energized = updateNewBeamsEnergized(b, b.direction, newBeams, energized, grid)
// 			}
// 		case '/':
// 			switch b.direction {
// 			case up:
// 				newBeams, energized = updateNewBeamsEnergized(b, rt, newBeams, energized, grid)
// 			case lt:
// 				newBeams, energized = updateNewBeamsEnergized(b, dn, newBeams, energized, grid)
// 			case dn:
// 				newBeams, energized = updateNewBeamsEnergized(b, lt, newBeams, energized, grid)
// 			case rt:
// 				newBeams, energized = updateNewBeamsEnergized(b, up, newBeams, energized, grid)
// 			}
// 		case '\\':
// 			switch b.direction {
// 			case up:
// 				newBeams, energized = updateNewBeamsEnergized(b, lt, newBeams, energized, grid)
// 			case lt:
// 				newBeams, energized = updateNewBeamsEnergized(b, up, newBeams, energized, grid)
// 			case dn:
// 				newBeams, energized = updateNewBeamsEnergized(b, rt, newBeams, energized, grid)
// 			case rt:
// 				newBeams, energized = updateNewBeamsEnergized(b, dn, newBeams, energized, grid)
// 			}
// 		default:
// 			panic(fmt.Sprintf("couldn't update from char %c\n", char))
// 		}
// 	}
// 	newBeams = filterOutDuplicateBeams(newBeams)
// 	return newBeams, grid, energized
// }

// func updateNewBeamsEnergized(b beam, newD direction, beams []beam, energized []string, grid []string) ([]beam, []string) {
// 	nextB := beam{}
// 	switch newD {
// 	case up:
// 		nextB = moveUp(b)
// 	case lt:
// 		nextB = moveLt(b)
// 	case dn:
// 		nextB = moveDn(b)
// 	case rt:
// 		nextB = moveRt(b)
// 	}

// 	if len(filterOutFinishedBeams([]beam{nextB}, grid, energized)) > 0 {
// 		beams = append(beams, nextB)
// 		energized = updateEnergized(energized, nextB)
// 	}

// 	return beams, energized
// }

// func moveUp(b beam) beam {
// 	b.location.row -= 1
// 	b.direction = up
// 	return b
// }

// func moveLt(b beam) beam {
// 	b.location.col -= 1
// 	b.direction = lt
// 	return b
// }

// func moveDn(b beam) beam {
// 	b.location.row += 1
// 	b.direction = dn
// 	return b
// }

// func moveRt(b beam) beam {
// 	b.location.col += 1
// 	b.direction = rt
// 	return b
// }

// func updateGrid(grid []string, p point, d direction) []string {
// 	char := string(grid[p.row][p.col])
// 	// fmt.Printf("updateGrid p: %v, char: %s\n", p, string(char))
// 	newChar := char
// 	switch char {
// 	case ".":
// 		switch d {
// 		case up:
// 			newChar = "^"
// 		case lt:
// 			newChar = "<"
// 		case dn:
// 			newChar = "v"
// 		case rt:
// 			newChar = ">"
// 		}
// 	case "^", "<", "v", ">":
// 		newChar = "2"
// 	case "2", "3", "4", "5", "6", "7", "8":
// 		newChar = fmt.Sprintf("%d", utils.Atoi(string(char))+1)
// 	}
// 	row := grid[p.row]
// 	row = row[:p.col] + string(newChar) + row[p.col+1:]
// 	// fmt.Printf("new row: %s\n", row)
// 	grid[p.row] = row
// 	return grid
// }

// func filterOutFinishedBeams(beams []beam, grid []string, energized []string) []beam {
// 	filtered := []beam{}
// 	maxRow := len(grid) - 1
// 	maxCol := len(grid[0]) - 1
// 	for _, b := range beams {
// 		if b.location.row < 0 || b.location.row > maxRow ||
// 			b.location.col < 0 || b.location.col > maxCol {
// 			continue
// 		}

// 		char := energized[b.location.row][b.location.col]

// 		if char != '.' && utils.Atoi(string(char)) == int(b.direction) {
// 			continue
// 		}

// 		filtered = append(filtered, b)
// 	}
// 	return filtered
// }

// func filterOutDuplicateBeams(beams []beam) []beam {
// 	filtered := []beam{}
// 	found := map[beam]int{}
// 	for _, b := range beams {
// 		if found[b] == 1 {
// 			continue
// 		}
// 		found[b] = 1
// 		filtered = append(filtered, b)
// 	}
// 	return filtered
// }

// func updateEnergized(energized []string, p beam) []string {
// 	if energized[p.location.row][p.location.col] != '.' {
// 		return energized
// 	}
// 	row := energized[p.location.row]
// 	row = row[:p.location.col] + fmt.Sprintf("%d", p.direction) + row[p.location.col+1:]
// 	energized[p.location.row] = row
// 	return energized
// }

// func printGrid(grid []string) {
// 	fmt.Println()
// 	for _, line := range grid {
// 		fmt.Println(line)
// 	}
// }
