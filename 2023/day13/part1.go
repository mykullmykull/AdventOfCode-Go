package day13

import "fmt"

func runA(input []string) int {
	sum := 0
	patterns := parse(input)
	for _, pattern := range patterns {
		sum += getValueOfReflectionLine(pattern)
	}
	return sum
}

func parse(input []string) [][]string {
	patterns := [][]string{}
	pattern := []string{}
	for _, line := range input {
		if line == "" {
			patterns = append(patterns, pattern)
			pattern = []string{}
			continue
		}
		pattern = append(pattern, line)
	}
	return patterns
}

func getValueOfReflectionLine(pattern []string) int {
	debug := false
	rowsAboveReflection := findRowsAboveReflection(pattern)
	if debug {
		fmt.Printf("rows: %d\n", rowsAboveReflection)
	}
	colsLeftOfReflection := findColsLeftOfReflection(pattern)
	if debug {
		fmt.Printf("cols: %d\n", colsLeftOfReflection)
	}

	return colsLeftOfReflection + (100 * rowsAboveReflection)
}

func findColsLeftOfReflection(pattern []string) int {
	debug := false
	for col := 0; col < len(pattern[0])-1; col++ {
		if debug {
			fmt.Printf("col: %d\n", col)
		}
		if colIsLeftOfReflection(col, pattern) {
			if debug {
				fmt.Println("is left of reflection")
			}
			// need the number of columns, not the index
			return col + 1
		}
	}
	return 0
}

func colIsLeftOfReflection(col int, pattern []string) bool {
	debug := false
	d := 0
	for {
		if col-d < 0 || col+d >= len(pattern[0])-1 {
			break
		}
		if debug {
			fmt.Printf("  col: %d, d: %d\n", col, d)
		}
		if !columnsReflect(col-d, col+d+1, pattern) {
			if debug {
				fmt.Println("  failed")
			}
			return false
		}
		d++
	}
	if debug {
		fmt.Println("  passed")
	}
	return true
}

func columnsReflect(c1 int, c2 int, pattern []string) bool {
	debug := false
	for row := 0; row < len(pattern); row++ {
		if debug {
			fmt.Printf("    row: %d, c1: %d, c2: %d, char1: %c, char2: %c\n", row, c1, c2, pattern[row][c1], pattern[row][c2])
		}
		if pattern[row][c1] != pattern[row][c2] {
			return false
		}
	}
	return true
}

func findRowsAboveReflection(pattern []string) int {
	debug := false
	for row := 0; row < len(pattern)-1; row++ {
		if debug {
			fmt.Printf("row: %d\n", row)
		}
		if rowIsAboveReflection(row, pattern) {
			if debug {
				fmt.Println("is above reflection")
			}
			// need the number of rows, not the index
			return row + 1
		}
	}
	return 0
}

func rowIsAboveReflection(row int, pattern []string) bool {
	debug := false
	d := 0
	for {
		if row-d < 0 || row+d >= len(pattern)-1 {
			break
		}
		if debug {
			fmt.Printf("  row: %d, d: %d\n", row, d)
		}
		if !rowsReflect(row-d, row+d+1, pattern) {
			if debug {
				fmt.Println("  failed")
			}
			return false
		}
		d++
	}
	if debug {
		fmt.Println("  passed")
	}
	return true
}

func rowsReflect(r1 int, r2 int, pattern []string) bool {
	debug := false
	for col := 0; col < len(pattern[0]); col++ {
		if debug {
			fmt.Printf("    col: %d, r1: %d, r2: %d, char1: %c, char2: %c\n", col, r1, r2, pattern[r1][col], pattern[r2][col])
		}
		if pattern[r1][col] != pattern[r2][col] {
			return false
		}
	}
	return true
}
