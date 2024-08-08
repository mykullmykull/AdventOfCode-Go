package day21

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_runB(t *testing.T) {
	// My solution was built with assumptions that don't necessarily hold for the testInput
	// specifically: center row and center column have no rocks
	// therefore: the shortest path to any subfarm is to its center edge
	// and: the first spots encountered in any subfarm are always reached by an even number of steps

	// assert.Equal(t, e2, runB(testInput, s2))
	// assert.Equal(t, e3, runB(testInput, s3))
	// assert.Equal(t, e4, runB(testInput, s4))
	// assert.Equal(t, e5, runB(testInput, s5))
	// assert.Equal(t, e6, runB(testInput, s6))
	// assert.Equal(t, e7, runB(testInput, s7))
	// assert.Equal(t, e8, runB(testInput, s8))

	// REAL MAP - targets obtained by running runA from part1 with the same values
	// Within center subFarm
	assert.Equal(t, 3082, runB(realInput, 60))
	assert.Equal(t, 3231, runB(realInput, 61))

	// Just outside center subFarm
	assert.Equal(t, 5029, runB(realInput, 75))
	assert.Equal(t, 5127, runB(realInput, 76))

	// With full center subfarm and corners
	assert.Equal(t, 17067, runB(realInput, 140))
	assert.Equal(t, 17306, runB(realInput, 141))

	// With edges
	assert.Equal(t, 216464, runB(realInput, 500))
	assert.Equal(t, 217314, runB(realInput, 501))

	// With more edges
	assert.Equal(t, 423075, runB(realInput, 700))
	assert.Equal(t, 423914, runB(realInput, 701))

	fmt.Printf("day1b: %d\n", runB(realInput, realSteps2))
}
