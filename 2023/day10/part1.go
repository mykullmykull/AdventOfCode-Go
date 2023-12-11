package day10

import (
	"fmt"
	"strings"
)

type direction string

const (
	n direction = "north"
	e direction = "east"
	w direction = "west"
	s direction = "south"
)

type tile struct {
	row  int
	col  int
	next direction
}

func runA(input []string) int {
	startingTiles := findStartTiles(input)
	startingKey := startingTiles[0].row*1000 + startingTiles[0].col

	// key = row * 1000 + col
	// value = distance
	loopDistances := map[int]int{}
	loopDistances[startingKey] = 0

	return findMostDistantLoopTile(loopDistances, startingTiles, input)
}

func findStartTiles(input []string) []tile {
	row, col := findStartRowAndCol(input)
	tiles := []tile{}

	// is n a valid next
	if row-1 >= 0 &&
		strings.Contains("|7F", string(input[row-1][col])) {
		tiles = append(tiles, tile{row: row, col: col, next: n})
	}

	// is e a valid next
	if len(input[0]) > col+1 &&
		strings.Contains("-j7", string(input[row][col+1])) {
		tiles = append(tiles, tile{row: row, col: col, next: e})
	}

	// is w a valid next
	if col-1 >= 0 &&
		strings.Contains("-LF", string(input[row][col-1])) {
		tiles = append(tiles, tile{row: row, col: col, next: w})
	}

	// is s a valid next
	if len(input) > row+1 &&
		strings.Contains("|LJ", string(input[row+1][col])) {
		tiles = append(tiles, tile{row: row, col: col, next: s})
	}

	if len(tiles) != 2 {
		panic(fmt.Sprintf("There aren't exactly 2 starting directions, there are %d\n", len(tiles)))
	}

	return tiles
}

func findStartRowAndCol(input []string) (int, int) {
	for i, row := range input {
		for j, col := range row {
			if col == 'S' {
				return i, j
			}
		}
	}
	panic("Couldn't ever find the start.")
}

func findMostDistantLoopTile(loopDistances map[int]int, currentTiles []tile, input []string) int {
	maxDistance := 0
	for {
		maxDistance++
		newCurrentTiles := []tile{}
		for _, tile := range currentTiles {
			newTile := getNewTile(tile, input)
			newTileKey := newTile.row*1000 + newTile.col
			if loopDistances[newTileKey] > 0 {
				return maxDistance
			}
			loopDistances[newTileKey] = maxDistance
			newCurrentTiles = append(newCurrentTiles, newTile)
		}

		currentTiles = newCurrentTiles
	}
}

func getNewTile(this tile, input []string) tile {
	if this.next == n {
		nextTile := tile{row: this.row - 1, col: this.col}
		switch input[nextTile.row][nextTile.col] {
		case '|':
			nextTile.next = n
		case '7':
			nextTile.next = w
		case 'F':
			nextTile.next = e
		default:
			panic(fmt.Sprintf("going n, but next tile is %c/n", input[nextTile.row][nextTile.col]))
		}
		return nextTile
	}

	if this.next == e {
		nextTile := tile{row: this.row, col: this.col + 1}
		switch input[nextTile.row][nextTile.col] {
		case '-':
			nextTile.next = e
		case 'J':
			nextTile.next = n
		case '7':
			nextTile.next = s
		default:
			panic(fmt.Sprintf("going e, but next tile is %c/n", input[nextTile.row][nextTile.col]))
		}
		return nextTile
	}

	if this.next == w {
		nextTile := tile{row: this.row, col: this.col - 1}
		switch input[nextTile.row][nextTile.col] {
		case '-':
			nextTile.next = w
		case 'L':
			nextTile.next = n
		case 'F':
			nextTile.next = s
		default:
			panic(fmt.Sprintf("going w, but next tile is %c/n", input[nextTile.row][nextTile.col]))
		}
		return nextTile
	}

	if this.next == s {
		nextTile := tile{row: this.row + 1, col: this.col}
		switch input[nextTile.row][nextTile.col] {
		case '|':
			nextTile.next = s
		case 'L':
			nextTile.next = e
		case 'J':
			nextTile.next = w
		default:
			panic(fmt.Sprintf("going s, but next tile is %c/n", input[nextTile.row][nextTile.col]))
		}
		return nextTile
	}

	panic(fmt.Sprintf("couldn't find next tile from %v/n", this))
}
