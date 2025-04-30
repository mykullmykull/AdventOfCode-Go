package day

import (
	"goAoc2022/utils"
)

func main(coords []int) int {
	positions := make([]int, len(coords))
	for x := 0; x < len(coords); x++ {
		positions[x] = x
	}
	coords = multiply(coords, 811589153)
	for x := 0; x < 10; x++ {
		println(x)
		coords, positions = mix(coords, positions)
	}
	coords = cycleToZero(coords)
	a := coords[1000%len(coords)]
	b := coords[2000%len(coords)]
	c := coords[3000%len(coords)]
	println("a:", a, "b:", b, "c:", c)
	return a + b + c
}

func multiply(coords []int, multiplier int) []int {
	newCoords := make([]int, len(coords))
	for x := 0; x < len(coords); x++ {
		newCoords[x] = coords[x] * multiplier
	}
	return newCoords
}

func mix(coords []int, positions []int) ([]int, []int) {
	for x := 0; x < len(positions); x++ {
		pos := positions[x]
		coord := coords[pos]
		newPos := getNewPos(pos, coord, len(coords))

		// rebuild coords
		withoutPos := append(coords[:pos], coords[pos+1:]...)
		before := copySlice(withoutPos, 0, newPos)
		withVal := append(before, coord)
		after := copySlice(withoutPos, newPos, len(withoutPos))
		coords = append(withVal, after...)

		// rebuild indexes
		min, max := sortMinMax(pos, newPos)
		for y := 0; y < len(positions); y++ {
			if y == x {
				continue
			}
			if positions[y] >= min && positions[y] <= max {
				if pos == max {
					positions[y]++
				} else {
					positions[y]--
				}
			}
		}
		positions[x] = len(before)
	}
	return coords, positions
}

func getNewPos(x int, val int, len int) int {
	if val == 0 {
		return x
	}
	abs := utils.AbsForInt(val)
	distance := abs % (len - 1)
	direction := val / abs
	newX := x + distance*direction
	toAdd := utils.DivideAndRoundUp(newX/direction, (len-1)) * (len - 1)
	return (newX + toAdd) % (len - 1)
}

func copySlice(slice []int, start int, end int) []int {
	newSlice := make([]int, end-start)
	copy(newSlice, slice[start:end])
	return newSlice
}

func sortMinMax(a int, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func cycleToZero(coords []int) []int {
	indexOf0 := -1
	for x := 0; x < len(coords); x++ {
		if coords[x] == 0 {
			indexOf0 = x
		}
	}
	return append(coords[indexOf0:], coords[:indexOf0]...)
}
