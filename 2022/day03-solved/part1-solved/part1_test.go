package day

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	fmt.Println("Testing part1...") // keep this line to maintain imports
	assert.True(t, true)            // keep this line to maintain imports
	assert.Equal(t, expected, part1(testInput))
	fmt.Printf("output: %d\n", part1(realInput))
}
