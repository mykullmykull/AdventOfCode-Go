package day13

func runB(input []string) int {
	sum := 0
	patterns := parse(input)
	for _, pattern := range patterns {
		// fmt.Println()
		// for _, line := range pattern {
		// 	fmt.Println(line)
		// }

		oldValue := getValueOfReflectionLine(pattern)
		oldRowValue := oldValue / 100
		oldColValue := oldValue % 100

		foundSmudge := false

		for r, row := range pattern {
			for c, _ := range row {
				char := string(pattern[r][c])
				newChar := "."
				if char == newChar {
					newChar = "#"
				}

				// need a copy that we can mark up to mess with smudges
				copiedPattern := make([]string, len(pattern))
				copy(copiedPattern, pattern)
				row := copiedPattern[r]
				row = row[:c] + newChar + row[c+1:]

				copiedPattern[r] = row
				// fmt.Println()
				// for _, line := range copiedPattern {
				// 	fmt.Println(line)
				// }

				newValue := getValueOfReflectionLine(copiedPattern)
				// fmt.Printf("new value: %d\n", newValue)
				newRowValue := newValue / 100
				newColValue := newValue % 100

				if newRowValue == oldRowValue && newColValue == oldColValue {
					continue
				}

				if newRowValue == oldRowValue {
					sum += newColValue
					foundSmudge = true
					break
				} else {
					sum += newRowValue * 100
					foundSmudge = true
					break
				}
			}
			if foundSmudge {
				break
			}
		}

	}
	return sum
}
