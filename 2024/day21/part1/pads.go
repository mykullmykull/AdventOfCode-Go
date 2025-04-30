package day

import (
	"fmt"
	"goAoc2024/utils"
	"strings"
)

/*

nPad:
7   8   9
4   5   6
1   2   3
	0   A

dPad:
	^   A
<   v   >

*/

type button struct {
	row int
	col int
}

type pad struct {
	rows    int
	cols    int
	buttons []string
	spots   map[string]button
	paths   map[pair][]string
}

var nPad = pad{
	rows:    4,
	cols:    3,
	buttons: []string{" ", "7", "8", "9", "4", "5", "6", "1", "2", "3", "0", "A"},
	spots: map[string]button{
		" ": {3, 0},
		"A": {3, 2},
		"0": {3, 1},
		"1": {2, 0},
		"2": {2, 1},
		"3": {2, 2},
		"4": {1, 0},
		"5": {1, 1},
		"6": {1, 2},
		"7": {0, 0},
		"8": {0, 1},
		"9": {0, 2},
	},
}

var dPad = pad{
	rows:    2,
	cols:    3,
	buttons: []string{" ", "^", "A", "<", "v", ">"},
	spots: map[string]button{
		" ": {0, 0},
		"A": {0, 2},
		"^": {0, 1},
		"<": {1, 0},
		"v": {1, 1},
		">": {1, 2},
	},
}

func (p pad) mapPaths() pad {
	paths := map[pair][]string{}
	for _, start := range p.buttons {
		for _, end := range p.buttons {
			if start == " " || end == " " {
				continue
			}
			paths[pair{start, end}] = p.getPossiblePaths(start, end)
		}
	}

	p.paths = paths
	return p
}

func (p pad) getPossiblePaths(start, end string) []string {
	possiblePaths := []string{}
	vert := p.spots[end].row - p.spots[start].row
	vStr := "v"
	if vert < 0 {
		vStr = "^"
	}
	horz := p.spots[end].col - p.spots[start].col
	hStr := ">"
	if horz < 0 {
		hStr = "<"
	}
	toGap := p.getPathToGap(start)
	binaries := utils.AllBinaries(vert, horz)
	println("start", start, "end", end)
	fmt.Printf("  vert %d, horz %d, binaries %v\n", vert, horz, binaries)
	for _, b := range binaries {
		b = strings.Replace(b, "0", vStr, -1)
		b = strings.Replace(b, "1", hStr, -1)
		if toGap != "" && strings.Index(b, toGap) == 0 {
			continue
		}
		possiblePaths = append(possiblePaths, b)
	}
	return possiblePaths
}

func (p pad) getPathToGap(start string) string {
	gapSpot := p.spots[" "]
	vert := gapSpot.row - p.spots[start].row
	horz := gapSpot.col - p.spots[start].col
	if vert == 0 {
		hStr := ">"
		if horz < 0 {
			hStr = "<"
		}
		return strings.Repeat(hStr, utils.Abs(horz))
	}
	if horz == 0 {
		vStr := "v"
		if vert < 0 {
			vStr = "^"
		}
		return strings.Repeat(vStr, utils.Abs(vert))
	}
	return ""
}
