package day

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart2(t *testing.T) {
	fmt.Println("Testing part2...") // keep this line to maintain imports
	assert.True(t, true)            // keep this line to maintain imports
	assert.Equal(t, expected, part2(testStacks, testMoves))
	fmt.Printf("output: %s\n", part2(realStacks, realMoves))
}
