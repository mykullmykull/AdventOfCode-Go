package day15

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_runA(t *testing.T) {
	assert.Equal(t, expected, runA(testInput))
	assert.Equal(t, expected2, runA(testInput2))
	fmt.Printf("day15a: %d\n", runA(realInput))
}

func Test_runB(t *testing.T) {
	assert.Equal(t, expected3, runB(testInput2))
	fmt.Printf("day15b: %d\n", runB(realInput))
}
