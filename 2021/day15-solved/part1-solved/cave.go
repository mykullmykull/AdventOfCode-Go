package day

import (
	"fmt"
	"goAoc2021/utils"
	"math"
)

type point struct {
	row int
	col int
}

type cave struct {
	grid    []string
	riskMap map[point]int
}

func (c *cave) isInBounds(p point) bool {
	if p.row < 0 || p.row >= len(c.grid) {
		return false
	}
	if p.col < 0 || p.col >= len(c.grid[p.row]) {
		return false
	}
	return true
}

func (c *cave) checkBounds(p point) {
	if !c.isInBounds(p) {
		panic(fmt.Sprintf("Point out of bounds: %v", p))
	}
}

func (c *cave) pointValue(p point) string {
	c.checkBounds(p)
	return c.grid[p.row][p.col : p.col+1]
}

func (c *cave) pointRisk(p point) int {
	c.checkBounds(p)
	return utils.Atoi(c.pointValue(p))

}

func (c *cave) markGrid(p point) {
	c.checkBounds(p)
	row := c.grid[p.row]
	c.grid[p.row] = row[:p.col] + "X" + row[p.col+1:]
}

func (c *cave) gridSize() int {
	rows := len(c.grid)
	cols := len(c.grid[0])
	return rows * cols
}

func (c *cave) update() {
	min := c.findMinAdjacentValue()
	minAdjacentPoints := c.findMinAdjacentPoints(min)
	for p := range minAdjacentPoints {
		c.riskMap[p] = c.getMinRiskTotal(p)
		fmt.Printf("  Point: %v, Risk: %d\n", p, c.riskMap[p])
		c.markGrid(p)
	}
	println()
	println()
}

func (c *cave) minimizeRiskMap() bool {
	madeChange := false
	for r := 0; r < len(c.grid); r++ {
		for col := 0; col < len(c.grid[r]); col++ {
			p := point{r, col}
			if p.row == 0 && p.col == 0 {
				continue // Skip the starting point
			}
			currentRisk := c.riskMap[p]
			value := c.pointRisk(p)
			min := c.findMinAdjacentRisk(p)
			if currentRisk != value+min {
				madeChange = true
				c.riskMap[p] = value + min
			}
		}
	}
	return madeChange
}

func (c *cave) findMinAdjacentRisk(p point) int {
	c.checkBounds(p)
	min := math.MaxInt
	for _, adj := range c.getAdjacents(p) {
		risk := c.riskMap[adj]
		if risk < min {
			min = risk
		}
	}
	return min
}

func (c *cave) findMinAdjacentValue() int {
	min := math.MaxInt
	for r := 0; r < len(c.grid); r++ {
		for col := 0; col < len(c.grid[r]); col++ {
			p := point{r, col}
			if c.pointValue(p) == "X" {
				continue
			}
			if c.isAdjacentToX(p) {
				risk := c.pointRisk(p)
				if risk < min {
					min = risk
				}
			}
		}
	}
	return min
}

func (c *cave) isAdjacentToX(p point) bool {
	c.checkBounds(p)
	for _, adj := range c.getAdjacents(p) {
		if c.pointValue(adj) == "X" {
			return true
		}
	}
	return false
}

func (c *cave) getAdjacents(p point) []point {
	c.checkBounds(p)
	candidates := []point{
		{p.row - 1, p.col}, // Up
		{p.row + 1, p.col}, // Down
		{p.row, p.col - 1}, // Left
		{p.row, p.col + 1}, // Right
	}
	adjacents := []point{}
	for _, adj := range candidates {
		if c.isInBounds(adj) {
			adjacents = append(adjacents, adj)
		}
	}
	return adjacents
}

func (c *cave) findMinAdjacentPoints(min int) map[point]bool {
	minAdjacentPoints := make(map[point]bool)
	for r := 0; r < len(c.grid); r++ {
		for col := 0; col < len(c.grid[r]); col++ {
			p := point{r, col}
			if c.pointValue(p) == "X" {
				continue
			}
			if c.isAdjacentToX(p) && c.pointRisk(p) == min {
				minAdjacentPoints[p] = true
			}
		}
	}

	return minAdjacentPoints
}

func (c *cave) getMinRiskTotal(p point) int {
	c.checkBounds(p)
	risk := c.pointRisk(p)
	minAdjacentRisk := math.MaxInt
	for _, adj := range c.getAdjacents(p) {
		if c.pointValue(adj) != "X" {
			continue
		}
		adjacentRisk := c.riskMap[adj]
		if adjacentRisk < minAdjacentRisk {
			minAdjacentRisk = adjacentRisk
		}
	}
	return risk + minAdjacentRisk
}

func (c *cave) printGrid() {
	for _, row := range c.grid {
		fmt.Println(row)
	}
	println()
}

func (c *cave) printMap() {
	for r := 0; r < len(c.grid); r++ {
		for col := 0; col < len(c.grid[r]); col++ {
			p := point{r, col}
			risk := c.riskMap[p]
			fmt.Printf("Point: %v, Risk: %d\n", p, risk)
		}
	}
	println()
}
