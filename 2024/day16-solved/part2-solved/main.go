package day

import "fmt"

func main(gridOg []string, best int) int {
	bestTurns := best / 1000
	bestSteps := best % 1000

	mOg := maze{grid: gridOg}
	println("original")
	mOg.printGrid()
	gridOg = mOg.removeDeadEnds().grid
	println("filtered out dead ends")
	mOg.printGrid()

	valid := make([]string, len(gridOg))
	mValid := maze{grid: valid}
	copy(mValid.grid, gridOg)

	for row := range gridOg {
		for col := range gridOg[row] {
			s := spot{row, col, none}
			g := make([]string, len(gridOg))
			m := maze{spots: []spot{s}, grid: g}
			copy(m.grid, gridOg)
			m = m.growToBestTurns(bestTurns)
			if m.hasBestTurns(bestTurns) {
				mValid.updateGrid(s, "T")
				println(fmt.Sprintf("found point with best turns: %v", s))
				mValid.printGrid()
			}
		}
	}

	mValid = mValid.markUnreachable()
	println("filtered by turns")
	mValid.printGrid()
	mValid = mValid.removeDeadEnds()
	println("filtered out dead ends")
	mValid.printGrid()
	filteredByTurns := make([]string, len(gridOg))
	copy(filteredByTurns, mValid.grid)

	for row := range gridOg {
		for col := range gridOg[row] {
			s := spot{row, col, none}
			g := make([]string, len(gridOg))
			m := maze{spots: []spot{s}, grid: g}
			copy(m.grid, filteredByTurns)

			if m.hasBestSteps(bestSteps) {
				mValid.updateGrid(s, "X")
				println(fmt.Sprintf("found point with best steps: %v", s))
				mValid.printGrid()
			}
		}
	}
	println("filtered by steps")
	mValid.printGrid()
	mValid = mValid.filterOutBackTrackers()
	println("filtered out back trackers")
	mValid.printGrid()
	return mValid.countXs() + 2
}
