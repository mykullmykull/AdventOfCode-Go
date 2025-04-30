package day

func main(input []string, moves string) int {
	println("grid")
	printGrid(input)
	grid := expand(input)
	println("expanded grid")
	printGrid(grid)

	w := warehouse{grid: grid}.getRobot()
	for len(moves) > 0 {
		next := toConst(moves[0])
		println("next:", toString(next))
		moves = moves[1:]
		w = w.move(next)
		printGrid(w.grid)
	}

	return sumBoxesGps(w.grid)
}

func sumBoxesGps(grid []string) int {
	cols := 0
	rows := 0
	for row := range grid {
		for col := range grid[row] {
			val := toConst(grid[row][col])
			if val == boxLt {
				cols += col
				rows += row * 100
			}
		}
	}
	return cols + rows
}
