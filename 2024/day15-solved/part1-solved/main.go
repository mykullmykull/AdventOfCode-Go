package day

func main(grid []string, moves string) int {
	printGrid(grid)
	robot := findRobot(grid)
	for _, move := range moves {
		println("move", string(move))
		change := point{0, 0}
		switch move {
		case '^':
			change.row = -1
		case '>':
			change.col = 1
		case 'v':
			change.row = 1
		case '<':
			change.col = -1
		default:
			panic("invalid move " + string(move))
		}

		grid, robot = moveIfNotBlocked(grid, robot, change)
		printGrid(grid)
	}
	sum := sumBoxesGps(grid)
	return sum
}

func findRobot(grid []string) point {
	for row, line := range grid {
		for col, cell := range line {
			if cell == '@' {
				return point{row, col}
			}
		}
	}
	panic("robot not found")
}

func moveIfNotBlocked(grid []string, robot point, change point) ([]string, point) {
	row := robot.row
	col := robot.col
	for {
		row += change.row
		col += change.col
		if grid[row][col] == 'O' {
			continue
		}
		if grid[row][col] == '#' {
			return grid, robot
		}
		for row != robot.row || col != robot.col {
			newRow := row - change.row
			newCol := col - change.col
			toMove := grid[newRow][newCol : newCol+1]
			grid[row] = grid[row][:col] + toMove + grid[row][col+1:]
			row = newRow
			col = newCol
		}
		grid[row] = grid[row][:col] + "." + grid[row][col+1:]
		row += change.row
		col += change.col
		return grid, point{row, col}
	}
}

func sumBoxesGps(grid []string) int {
	sum := 0
	for r, line := range grid {
		for c, cell := range line {
			if cell == 'O' {
				sum += 100*r + c
			}
		}
	}
	return sum
}

func printGrid(grid []string) {
	for _, line := range grid {
		println(line)
	}
	println()
}
