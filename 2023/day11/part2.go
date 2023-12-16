package day11

import (
	"strings"
)

func runB(input []string, expansionFactor int) int {
	sum := 0
	galaxies, emptyRows, emptyCols := parse(input)

	expansionFactor -= 1
	for i, g := range galaxies {
		for j, g2 := range galaxies {
			if i <= j {
				continue
			}

			distance := abs(g.row-g2.row) + abs(g.col-g2.col)
			distance = increaseByExpansionFactor(distance, g.row, g2.row, emptyRows, expansionFactor)
			distance = increaseByExpansionFactor(distance, g.col, g2.col, emptyCols, expansionFactor)
			sum += distance
		}
	}
	return sum
}

func parse(input []string) ([]point, []int, []int) {
	galaxies := findAllGalaxies(input)
	emptyRows := []int{}
	for i, row := range input {
		if strings.Contains(row, "#") {
			continue
		}
		emptyRows = append(emptyRows, i)
	}

	emptyCols := []int{}
	for i := 0; i < len(input[0]); i++ {
		isColEmpty := true
		for _, row := range input {
			char := row[i]
			if char == '#' {
				isColEmpty = false
				break
			}
		}
		if isColEmpty {
			emptyCols = append(emptyCols, i)
		}
	}
	return galaxies, emptyRows, emptyCols
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

func increaseByExpansionFactor(distance int, g int, g2 int, empties []int, expansionFactor int) int {
	for _, rc := range empties {
		if g < rc && g2 > rc || g > rc && g2 < rc {
			distance += expansionFactor
		}
	}
	return distance
}
