package day

import (
	"fmt"
	"goAoc2022/utils"
	"strings"
)

type spot struct {
	r int
	c int
}

type thread struct {
	spots     [10]spot
	tailSpots map[spot]bool
}

func (s spot) updateGrid(grid []string, newStr string) []string {
	grid[s.r] = grid[s.r][:s.c] + newStr + grid[s.r][s.c+1:]
	return grid
}

func (t thread) printGrid(grid []string) {
	newGrid := make([]string, len(grid))
	copy(newGrid, grid)
	for s := range t.tailSpots {
		newGrid = s.updateGrid(newGrid, "*")
	}
	for x, s := range t.spots {
		newGrid = s.updateGrid(newGrid, fmt.Sprintf("%d", x))
	}
	for _, row := range newGrid {
		println(row)
	}
}

func (t thread) move(line string) thread {
	split := strings.Split(line, " ")
	direction := split[0]
	distance := utils.StrToInt(split[1])
	for x := 0; x < distance; x++ {
		for y := range t.spots {
			if y == 0 {
				t.spots[0] = t.spots[0].moveHead(direction)
				continue
			}
			t = t.moveTail(y)
		}
	}
	return t
}

func (s spot) moveHead(direction string) spot {
	switch direction {
	case "U":
		s.r--
	case "R":
		s.c++
	case "D":
		s.r++
	case "L":
		s.c--
	default:
		panic(fmt.Sprintf("Invalid direction %s\n", direction))
	}
	return s
}

func (t thread) moveTail(spotIndex int) thread {
	head := t.spots[spotIndex-1]
	tail := t.spots[spotIndex]
	rowGap := head.r - tail.r
	colGap := head.c - tail.c
	if rowGap <= 1 && rowGap >= -1 && colGap <= 1 && colGap >= -1 {
		return t
	}

	up := rowGap < 0 && colGap == 0
	rt := rowGap == 0 && colGap > 0
	dn := rowGap > 0 && colGap == 0
	lt := rowGap == 0 && colGap < 0
	upAndRt := rowGap < 0 && colGap > 0
	dnAndRt := rowGap > 0 && colGap > 0
	dnAndLt := rowGap > 0 && colGap < 0
	upAndLt := rowGap < 0 && colGap < 0

	if up || upAndRt || upAndLt {
		t.spots[spotIndex].r--
	}
	if dn || dnAndRt || dnAndLt {
		t.spots[spotIndex].r++
	}
	if rt || upAndRt || dnAndRt {
		t.spots[spotIndex].c++
	}
	if lt || upAndLt || dnAndLt {
		t.spots[spotIndex].c--
	}

	if spotIndex == 9 {
		t.tailSpots[t.spots[9]] = true
	}
	return t
}
