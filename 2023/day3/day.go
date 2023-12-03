package day3

import (
	"math"
	"strconv"
)

type gear struct {
	x int
	y int
}

func runB(input []string) int {
	sum := 0
	gears := getGears(input)

	for _, gear := range gears {
		minY := int(math.Max(float64(gear.y-1), 0))
		maxY := int(math.Min(float64(gear.y+1), float64(len(input)-1)))
		sum += getGearRatio(gear, input[minY:maxY+1])
	}
	return sum
}

func getGears(input []string) []gear {
	gears := []gear{}
	for y, line := range input {
		for x := 0; x < len(line); x++ {
			if input[y][x] == '*' {
				gears = append(gears, gear{x, y})
			}
		}
	}
	return gears
}

func getGearRatio(gear gear, input []string) int {
	ratio := 1
	adjacentNumbers := 0
	for i, line := range input {
		y := gear.y - 1 + i
		for x := 0; x < len(line); x++ {
			if line[x] == '.' {
				continue
			}

			n := getNum(x, y, line)
			if n.val > 0 {
				x = n.xEnd - 1
				if isAdjacentToGear(n, gear) {
					adjacentNumbers++
					ratio *= n.val
				}
			}
		}
	}
	if adjacentNumbers == 2 {
		return ratio
	}
	return 0
}

func isAdjacentToGear(n num, gear gear) bool {
	minY := int(math.Max(float64(n.y-1), 0))
	maxY := n.y + 1
	minX := int(math.Max(float64(n.xStart-1), 0))
	maxX := n.xEnd

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if y == gear.y && x == gear.x {
				return true
			}
		}
	}
	return false
}

type num struct {
	val    int
	xStart int
	xEnd   int
	y      int
}

func runA(input []string) int {
	sum := 0
	for y, line := range input {
		for x := 0; x < len(line); x++ {
			if line[x] == '.' {
				continue
			}

			n := getNum(x, y, line)
			if n.val > 0 {
				x = n.xEnd - 1
				if isAdjacentToSymbol(n, input) {
					sum += n.val
				}
			}
		}
	}

	return sum
}

func getNum(x int, y int, line string) num {
	n := num{
		xStart: x,
		y:      y,
	}

	val := 0
	for x < len(line) {
		newDigit, err := strconv.Atoi(line[x : x+1])
		if err != nil {
			break
		}

		val = val*10 + newDigit
		x++
	}

	n.val = val
	n.xEnd = x
	return n
}

func isAdjacentToSymbol(n num, input []string) bool {
	minY := int(math.Max(float64(n.y-1), 0))
	maxY := int(math.Min(float64(n.y+1), float64(len(input)-1)))
	minX := int(math.Max(float64(n.xStart-1), 0))
	maxX := int(math.Min(float64(n.xEnd), float64(len(input[0])-1)))

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if isSymbol(input[y][x : x+1]) {
				return true
			}
		}
	}
	return false
}

func isSymbol(str string) bool {
	if str == "." {
		return false
	}

	_, err := strconv.Atoi(str)
	return err != nil
}
