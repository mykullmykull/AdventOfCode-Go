package day

import (
	"fmt"
	"goAoc2022/utils"
	"math"
	"strings"
)

type columns struct {
	col1 string
	col2 string
	col3 string
	col4 string
	col5 string
	col6 string
	col7 string
}

func (c columns) isEmpty() bool {
	for col := 1; col <= 7; col++ {
		if c.getSpace(col) != 0 {
			return false
		}
	}
	return true
}

func (c columns) getSpace(col int) int {
	array := c.toArray()
	column := array[col-1]
	space := strings.Index(column, "#")
	if space == -1 {
		return len(column)
	}
	return space
}

func (c columns) getMinSpace() int {
	min := math.MaxInt
	for col := 1; col <= 7; col++ {
		min = utils.UpdateLeast(min, c.getSpace(col))
	}
	return min
}

func (c columns) addHeight(height int) columns {
	array := c.toArray()
	for col := 1; col <= 7; col++ {
		str := array[col-1]
		for x := 0; x < height; x++ {
			str = "." + str
		}
		array[col-1] = str
	}
	return arrayToColumns(array)
}

func (c columns) moveUp() columns {
	array := c.toArray()
	for col := 1; col <= 7; col++ {
		str := array[col-1]
		if str[0] == '#' {
			panic("cannot delete the rock at col " + fmt.Sprintf("%d", col))
		}
		str = str[1:]
		array[col-1] = str
	}
	return arrayToColumns(array)
}

func (c columns) getMaxSpace() int {
	max := 0
	for col := 1; col <= 7; col++ {
		max = utils.UpdateMost(max, c.getSpace(col))
	}
	return max
}

func (c columns) markRock(point []int) columns {
	col := point[1]
	row := point[0]
	array := c.toArray()

	// expand column to fit point
	for len(c.col1) < row {
		for x := 0; x < 7; x++ {
			array[x] = "." + array[x]
		}
	}

	str := array[col-1]

	endCol := ""
	if row < len(str)-1 {
		endCol = str[row+1:]
	}

	str = str[:row] + "#" + endCol
	array[col-1] = str
	return arrayToColumns(array)
}

func (c columns) toArray() []string {
	return []string{c.col1, c.col2, c.col3, c.col4, c.col5, c.col6, c.col7}
}

func arrayToColumns(array []string) columns {
	return columns{
		col1: array[0],
		col2: array[1],
		col3: array[2],
		col4: array[3],
		col5: array[4],
		col6: array[5],
		col7: array[6],
	}
}

func (c columns) shrinkToMaxSpace() columns {
	max := c.getMaxSpace() + 10
	array := c.toArray()
	for col := 1; col <= 7; col++ {
		max = utils.UpdateLeast(max, len(array[col-1]))
		str := array[col-1][:max]
		array[col-1] = str
	}
	return arrayToColumns(array)
}

func (c columns) isBlocked(col int, row int) bool {
	array := c.toArray()
	str := array[col-1]
	return row >= len(str) || string(array[col-1][row]) == "#"
}
