package day

import (
	"fmt"
	"goAoc2021/utils"
	"strings"
)

type vent struct {
	startX int
	startY int
	endX   int
	endY   int
}

func main(input []string) int {
	ventMap := createVentMap(input)
	// loop through each line and update the seaFloorMap
	for _, line := range input {
		v := parse(line)
		// skip diagonal lines
		if v.startX != v.endX && v.startY != v.endY {
			continue
		}
		ventMap = v.update(ventMap)
		fmt.Printf("Updating ventMap with %v\n", v)
		// printMap(ventMap)
	}
	// count the number of points >= 2
	over2 := 0
	for _, row := range ventMap {
		for _, col := range row {
			if col >= 2 {
				over2++
			}
		}
	}
	return over2
}

func createVentMap(input []string) [][]int {
	maxX, maxY := getDimensions(input)
	ventMap := make([][]int, maxY)
	for y := range ventMap {
		row := make([]int, maxX)
		ventMap[y] = row
	}
	return ventMap
}

func getDimensions(input []string) (int, int) {
	maxX := 0
	maxY := 0
	for _, line := range input {
		v := parse(line)
		if v.startX > maxX {
			maxX = v.startX
		}
		if v.endX > maxX {
			maxX = v.endX
		}
		if v.startY > maxY {
			maxY = v.startY
		}
		if v.endY > maxY {
			maxY = v.endY
		}
	}
	return maxX + 1, maxY + 1
}

func parse(str string) vent {
	v := vent{}
	startAndEnd := strings.Split(str, " -> ")

	startXY := strings.Split(startAndEnd[0], ",")
	v.startX = utils.Atoi(startXY[0])
	v.startY = utils.Atoi(startXY[1])

	endXY := strings.Split(startAndEnd[1], ",")
	v.endX = utils.Atoi(endXY[0])
	v.endY = utils.Atoi(endXY[1])

	return v
}

func (v vent) update(ventMap [][]int) [][]int {
	xDiff := v.endX - v.startX
	xChange := 0
	if xDiff != 0 {
		xChange = xDiff / utils.Abs(xDiff)
	}

	yDiff := v.endY - v.startY
	yChange := 0
	if yDiff != 0 {
		yChange = yDiff / utils.Abs(yDiff)
	}

	y := v.startY
	x := v.startX
	for x != v.endX+xChange || y != v.endY+yChange {
		ventMap[y][x]++
		x += xChange
		y += yChange
	}
	return ventMap
}

func printMap(ventMap [][]int) {
	for _, row := range ventMap {
		for _, col := range row {
			fmt.Printf("%d", col)
		}
		println()
	}
	println()
}
