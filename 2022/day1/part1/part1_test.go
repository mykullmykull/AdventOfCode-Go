package day1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	fmt.Println("Testing part1...")
	assert.True(t, true)
	assert.Equal(t, expected, part1(testInput))
	fmt.Printf("output: %d\n", part1(realInput))
}
