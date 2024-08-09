package day

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	fmt.Println("Testing part1...") // keep this line to maintain imports
	assert.True(t, true)            // keep this line to maintain imports
	assert.Equal(t, expected1, part2(testInput1))
	assert.Equal(t, expected2, part2(testInput2))
	assert.Equal(t, expected3, part2(testInput3))
	assert.Equal(t, expected4, part2(testInput4))
	assert.Equal(t, expected5, part2(testInput5))
	fmt.Printf("output: %d\n", part2(realInput))
}
