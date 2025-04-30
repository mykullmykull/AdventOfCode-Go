package day

import (
	"goAoc2022/utils"
)

func main(coords []int) int {
	indexes := make([]int, len(coords))
	for x := 0; x < len(coords); x++ {
		indexes[x] = x
	}
	coords = mix(coords, indexes)
	coords = cycleToZero(coords)
	a := coords[1000%len(coords)]
	b := coords[2000%len(coords)]
	c := coords[3000%len(coords)]
	println("a:", a, "b:", b, "c:", c)
	return a + b + c
}

func mix(coords []int, indexes []int) []int {
	for len(indexes) > 0 {
		println(len(indexes))
		x := indexes[0]
		coord := coords[x]
		indexes = indexes[1:]
		newX := getNewX(x, coord, len(coords))

		// rebuild coords
		withoutX := append(coords[:x], coords[x+1:]...)
		before := copySlice(withoutX, 0, newX)
		withVal := append(before, coord)
		after := copySlice(withoutX, newX, len(withoutX))
		coords = append(withVal, after...)

		// rebuild indexes
		min, max := sortMinMax(x, newX)
		for y := 0; y < len(indexes); y++ {
			if indexes[y] >= min && indexes[y] <= max {
				if x == max {
					indexes[y]++
				} else {
					indexes[y]--
				}
			}
		}

	}
	return coords
}

func getNewX(x int, val int, len int) int {
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
