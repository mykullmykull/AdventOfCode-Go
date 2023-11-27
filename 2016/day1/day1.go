package day1

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

func runB(input string) int {
	x := 0
	y := 0
	facing := "+y"
	var visited []point
	isFinished := false
	for i := 0; i < len(input); i++ {
		if isFinished {
			break
		}

		char := input[i]
		if char == ',' || char == ' ' {
			continue
		}

		if char == 'R' || char == 'L' {
			facing = turn(facing, char)
			continue
		}

		nextComma := strings.Index(input[i:], ",")
		if nextComma == -1 {
			nextComma = len(input[i:])
		}

		distance, err := strconv.Atoi(input[i : i+nextComma])
		if err != nil {
			panic(fmt.Sprintf("Could not parse distance from %s", input[i:i+nextComma]))
		}
		i += nextComma

		for j := 0; j < distance; j++ {
			switch facing {
			case "+y":
				y += 1
			case "-y":
				y -= 1
			case "+x":
				x += 1
			case "-x":
				x -= 1
			}
			if alreadyVisited(visited, point{x, y}) {
				isFinished = true
				break
			}

			visited = append(visited, point{x, y})
		}
	}

	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func alreadyVisited(visited []point, p point) bool {
	for _, old := range visited {
		if old.x == p.x && old.y == p.y {
			return true
		}
	}
	return false
}

func runA(input string) int {
	x := 0
	y := 0
	facing := "+y"
	for i := 0; i < len(input); i++ {
		char := input[i]
		if char == ',' || char == ' ' {
			continue
		}

		if char == 'R' || char == 'L' {
			facing = turn(facing, char)
			continue
		}

		nextComma := strings.Index(input[i:], ",")
		if nextComma == -1 {
			nextComma = len(input[i:])
		}

		distance, err := strconv.Atoi(input[i : i+nextComma])
		if err != nil {
			panic(fmt.Sprintf("Could not parse distance from %s", input[i:i+nextComma]))
		}
		i += nextComma

		switch facing {
		case "+y":
			y += distance
		case "-y":
			y -= distance
		case "+x":
			x += distance
		case "-x":
			x -= distance
		}
	}

	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func turn(facing string, direction byte) string {
	switch facing {
	case "+y":
		if direction == 'R' {
			return "+x"
		} else {
			return "-x"
		}
	case "-y":
		if direction == 'R' {
			return "-x"
		} else {
			return "+x"
		}
	case "+x":
		if direction == 'R' {
			return "-y"
		} else {
			return "+y"
		}
	case "-x":
		if direction == 'R' {
			return "+y"
		} else {
			return "-y"
		}
	}
	panic(fmt.Sprintf("Could not turn from %s to the %b", facing, direction))
}
