package day

import (
	"goAoc2022/utils"
	"strings"
)

func parse(input []string) grid {
	rockLines := []rockLine{}
	for _, line := range input {
		spots := parseSpots(line)
		newRockLines := assembleRockLines(spots)
		rockLines = append(rockLines, newRockLines...)
	}
	return constructGrid(rockLines)
}

func parseSpots(line string) []spot {
	split := strings.Split(line, " -> ")
	spots := make([]spot, len(split))
	for x, str := range split {
		colRow := strings.Split(str, ",")
		s := spot{col: utils.StrToInt(colRow[0]), row: utils.StrToInt(colRow[1])}
		spots[x] = s
	}
	return spots
}

func assembleRockLines(spots []spot) []rockLine {
	rockLines := make([]rockLine, len(spots)-1)
	for x := 0; x < len(spots)-1; x++ {
		rockLines[x] = rockLine{start: spots[x], end: spots[x+1]}
	}
	return rockLines
}

func constructGrid(rockLines []rockLine) grid {
	maxCol, maxRow := getMaximums(rockLines)
	emptyRow := ""
	for x := 0; x < maxCol; x++ {
		emptyRow += "."
	}
	rows := []string{}
	for x := 0; x < maxRow; x++ {
		rows = append(rows, emptyRow)
	}
	g := grid{origin: spot{col: 500, row: 0}, rows: rows}
	return g.addRocks(rockLines)
}

func getMaximums(rockLines []rockLine) (int, int) {
	maxCol := -1
	maxRow := -1
	for _, line := range rockLines {
		if line.start.col > maxCol {
			maxCol = line.start.col
		}
		if line.end.col > maxCol {
			maxCol = line.end.col
		}
		if line.start.row > maxRow {
			maxRow = line.start.row
		}
		if line.end.row > maxRow {
			maxRow = line.end.row
		}
	}
	return maxCol + 5, maxRow + 2
}
