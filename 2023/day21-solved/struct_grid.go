package day21

import (
	"fmt"
)

type grid struct {
	rows                   []string
	start                  spot
	maxSteps               int
	maxFullSubFarmDistance int
	edgesOddity            int
	centerOddity           int
	subFarmSize            int
	evenSpotsCount         int
	oddSpotsCount          int
}

// ----------------------------------------- parser functions ----------------------------------------//
func (g grid) getStart() spot {
	for r, row := range g.rows {
		for c, cell := range row {
			if cell == 'S' {
				return spot{r, c}
			}
		}
	}
	return spot{-1, -1}
}

func (g grid) getCenterOddity() int {
	return g.maxSteps % 2
}

func (g grid) getSubFarmSize() int {
	width := len(g.rows[0])
	height := len(g.rows)
	if width != height {
		panic(fmt.Sprintf("subFarmWidth: %d != subFarmHeight: %d", width, height))
	}

	return width
}

func (g grid) getMaxFullSubFarmDistance() int {
	// we assume that if you can reach the center of the next subFarm,
	// then you can reach all of the previous subFarm
	return (g.maxSteps / g.subFarmSize) - 1
}

func (g grid) getEdgesOddity() int {
	return (g.maxFullSubFarmDistance + 1 + g.centerOddity) % 2
}

func (g grid) countEvenOdd() (int, int) {
	start := g.getStart()
	edgeSpots := []spot{start}

	// start is an even spot, but isn't accounted for in steps loop
	evenSpotsCount := 1
	oddSpotsCount := 0
	p := prev{}
	p = p.addToPrev(edgeSpots)
	step := 1
	for len(edgeSpots) > 0 {
		edgeSpots = g.moveOneStep(edgeSpots, p)
		evenSpotsCount, oddSpotsCount = addToEvenOdd(step, edgeSpots, evenSpotsCount, oddSpotsCount)
		step++
	}
	return evenSpotsCount, oddSpotsCount
}

// --------------------------------------- utility functions ---------------------------------------//
func (g grid) getSpot(l spot) string {
	return g.rows[l.r][l.c : l.c+1]
}

func (g grid) updateSpot(l spot, val string) {
	g.rows[l.r] = g.rows[l.r][:l.c] + val + g.rows[l.r][l.c+1:]
}

func (g grid) printGrid(edgeSpots []spot) {
	rowsToPrint := make([]string, len(g.rows))
	copy(rowsToPrint, g.rows)
	g.rows = rowsToPrint
	for _, s := range edgeSpots {
		g.updateSpot(s, "0")
	}

	for i, row := range g.rows {
		newRow := ""
		counter := 0
		for _, char := range row {
			newRow += string(char)
			counter++
			if counter == g.subFarmSize {
				newRow += " "
				counter = 0
			}
		}
		g.rows[i] = newRow
	}

	rowIndex := 0
	for rowIndex < len(g.rows) {
		if rowIndex%g.subFarmSize == 0 {
			println()
		}
		println(g.rows[rowIndex])
		rowIndex++
	}
}

func (g grid) printGridFromPrev(p prev) {
	g.printGrid(p.prevSpots())
}

// ----------------------------------------- part2 functions ----------------------------------------//

func (g grid) spotsFromEdgeSubFarm(stepsLeft int, edgeSpots []spot, e expander) int {
	evenSpotsCount, oddSpotsCount := len(edgeSpots), 0

	g.rows = e.expandRows()
	edgeSpots = e.adjustEdgeSpots(edgeSpots)

	p := prev{}
	p = p.addToPrev(edgeSpots)
	for step := 1; step < stepsLeft; step++ {
		toAdd := g.moveOneStep(edgeSpots, p)
		evenSpotsCount, oddSpotsCount = addToEvenOdd(step, toAdd, evenSpotsCount, oddSpotsCount)
		edgeSpots = toAdd
		p = p.addToPrev(toAdd)
	}
	// g.printGridFromPrev(p)
	// println("evenSpotsCount", evenSpotsCount, "oddSpotsCount", oddSpotsCount)
	if g.edgesOddity == 0 {
		return evenSpotsCount
	} else {
		return oddSpotsCount
	}
}

func (g grid) moveOneStep(edgeSpots []spot, p prev) []spot {
	newEdges := []spot{}
	for _, s := range edgeSpots {
		toAdd := s.spotsAfterOneStep(g, p)
		newEdges = append(newEdges, toAdd...)
		p = p.addToPrev(toAdd)
	}
	edgeSpots = newEdges

	return edgeSpots
}

func addToEvenOdd(step int, toAdd []spot, evenSpotsCount, oddSpotsCount int) (int, int) {
	if step%2 == 0 {
		evenSpotsCount += len(toAdd)
		return evenSpotsCount, oddSpotsCount
	}
	oddSpotsCount += len(toAdd)
	return evenSpotsCount, oddSpotsCount
}

func (g grid) swapFullSpots(currentFullSpots int) int {
	if currentFullSpots == 0 {
		if g.centerOddity == 0 {
			return g.evenSpotsCount
		}
		return g.oddSpotsCount
	}
	if currentFullSpots == g.evenSpotsCount {
		return g.oddSpotsCount
	}
	return g.evenSpotsCount
}

func (g grid) getEdgesOddityFromCurrentFullSpots(currentFullSpots int) int {
	if currentFullSpots == g.evenSpotsCount {
		return 0
	}
	return 1
}
