package day1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testPart2(t *testing.T) {
	fmt.Println("Testing part2...")
	assert.True(t, true)
	assert.Equal(t, expected, part2(testInput))
	fmt.Printf("output: %d\n", part2(realInput))
}
