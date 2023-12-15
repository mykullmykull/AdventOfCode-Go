package day13

import "fmt"

type point struct {
	row int
	col int
}

func runB(input []string) int {
	sum := 0
	patterns := parse(input)
	for _, pattern := range patterns {
		colLeftOfReflectionLine := findColLeftOfFoldWith1Smudge(pattern)
		rowAboveReflectionLine := findRowAboveFoldWith1Smudge(pattern)
		sum += colLeftOfReflectionLine + (100 * rowAboveReflectionLine)
	}
	return sum
}

func findColLeftOfFoldWith1Smudge(pattern []string) int {
	for col := 0; col < len(pattern[0]); col++ {
		smudges := countSmudgesWhenFoldingVertically(pattern, col)
		if smudges == 1 {
			return col + 1
		}
	}
	return 0
}

func countSmudgesWhenFoldingVertically(pattern []string, colLeftOfFold int) int {
	smudges := 0
	for c := 0; c <= colLeftOfFold; c++ {
		for r := 0; r < len(pattern); r++ {
			if smudges > 1 {
				return smudges
			}

			p := point{r, c}
			p2 := getReflectionPoint(p, -1, colLeftOfFold)
			if p2.col < 0 || p2.col >= len(pattern[0]) {
				break
			}

			if pattern[p.row][p.col] != pattern[p2.row][p2.col] {
				smudges++
			}
		}
	}
	return smudges
}

func getReflectionPoint(p point, rowAboveFold int, colLeftOfFold int) point {
	if rowAboveFold > -1 {
		return point{rowAboveFold + rowAboveFold - p.row + 1, p.col}
	}

	if colLeftOfFold > -1 {
		return point{p.row, colLeftOfFold + colLeftOfFold - p.col + 1}
	}
	panic(fmt.Sprintf("couldn't reflect point %v\n", p))
}

func findRowAboveFoldWith1Smudge(pattern []string) int {
	for row := 0; row < len(pattern); row++ {
		smudges := countSmudgesWhenFoldingHorizontally(pattern, row)
		if smudges == 1 {
			return row + 1
		}
	}
	return 0
}

func countSmudgesWhenFoldingHorizontally(pattern []string, rowAboveFold int) int {
	smudges := 0
	for r := 0; r <= rowAboveFold; r++ {
		for c := 0; c < len(pattern[0]); c++ {
			if smudges > 1 {
				return smudges
			}

			p := point{r, c}
			p2 := getReflectionPoint(p, rowAboveFold, -1)
			if p2.row < 0 || p2.row >= len(pattern) {
				break
			}

			if pattern[p.row][p.col] != pattern[p2.row][p2.col] {
				smudges++
			}
		}
	}
	return smudges
}

func printPattern(pattern []string) {
	fmt.Println()
	for _, line := range pattern {
		fmt.Println(line)
	}
}
