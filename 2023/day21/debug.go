package day21

func printGrid(grid []string, locations []location) {
	for _, l := range locations {
		row := []byte(grid[l.row])
		if row[l.col] == '#' {
			panic("location: " + l.toString() + " has a rock")
		}
		row[l.col] = 'O'
		grid[l.row] = string(row)
	}
	for _, row := range grid {
		println(row)
	}
}

func log(s string, debug bool) {
	if debug {
		println(s)
	}
}
