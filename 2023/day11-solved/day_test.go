package day11

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_runA(t *testing.T) {
	// assert.Equal(t, expected, runA(testInput))
	// fmt.Printf("day1a: %d\n", runA(realInput))
}

func Test_runB(t *testing.T) {
	assert.Equal(t, expected, runB(testInput, 2))
	assert.Equal(t, expected2, runB(testInput, 10))
	assert.Equal(t, expected3, runB(testInput, 100))
	fmt.Printf("day1a: %d\n", runB(realInput, 1000000))
}
