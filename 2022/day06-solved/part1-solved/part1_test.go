package day

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	fmt.Println("Testing part1...") // keep this line to maintain imports
	assert.True(t, true)            // keep this line to maintain imports
	assert.Equal(t, expected1, part1(testInput1))
	assert.Equal(t, expected2, part1(testInput2))
	assert.Equal(t, expected3, part1(testInput3))
	assert.Equal(t, expected4, part1(testInput4))
	assert.Equal(t, expected5, part1(testInput5))
	fmt.Printf("output: %d\n", part1(realInput))
}
