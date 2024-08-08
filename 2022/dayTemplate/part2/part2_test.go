package day1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart2(t *testing.T) {
	fmt.Println("Testing part2...") // keep this line to maintain imports
	assert.True(t, true)            // keep this line to maintain imports
	assert.Equal(t, expected, part2(testInput))
	fmt.Printf("output: %d\n", part2(realInput))
}
