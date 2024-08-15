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
	tail      spot
	head      spot
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
	newGrid = t.head.updateGrid(newGrid, "H")
	newGrid = t.tail.updateGrid(newGrid, "T")
	for _, row := range newGrid {
		println(row)
	}
}

func (t thread) move(line string) thread {
	split := strings.Split(line, " ")
	direction := split[0]
	distance := utils.StrToInt(split[1])
	for x := 0; x < distance; x++ {
		t = t.moveHead(direction)
		t = t.moveTail()
	}
	return t
}

func (t thread) moveHead(direction string) thread {
	switch direction {
	case "U":
		t.head.r--
	case "R":
		t.head.c++
	case "D":
		t.head.r++
	case "L":
		t.head.c--
	default:
		panic(fmt.Sprintf("Invalid direction %s\n", direction))
	}
	return t
}

func (t thread) moveTail() thread {
	rowGap := t.head.r - t.tail.r
	colGap := t.head.c - t.tail.c
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
		t.tail.r--
	}
	if dn || dnAndRt || dnAndLt {
		t.tail.r++
	}
	if rt || upAndRt || dnAndRt {
		t.tail.c++
	}
	if lt || upAndLt || dnAndLt {
		t.tail.c--
	}

	t.tailSpots[t.tail] = true
	return t
}
