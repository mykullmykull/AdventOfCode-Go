package day

func expand(input []string) []string {
	newGrid := []string{}
	for row := range input {
		newRow := ""
		for col := range input[row] {
			val := input[row][col]
			switch val {
			case '#':
				newRow += "##"
			case 'O':
				newRow += "[]"
			case '.':
				newRow += ".."
			case '@':
				newRow += "@."
			default:
				panic("Invalid value " + string(val))
			}
		}
		newGrid = append(newGrid, newRow)
	}
	return newGrid
}
