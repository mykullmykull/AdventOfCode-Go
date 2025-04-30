package day

func main(input []string) int {
	maze := maze{grid: input}
	maze = maze.initialize()
	for len(maze.racers) > 0 {
		maze = maze.process()
	}
	return maze.scores[maze.end.toInt()]
}
