package day8

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_runA(t *testing.T) {
	assert.Equal(t, expected, runA(testInput))
	assert.Equal(t, expected2, runA(testInput2))
	fmt.Printf("day8a: %d\n", runA(realInput))
}

func Test_runB(t *testing.T) {
	assert.Equal(t, expected3, runB(testInput3))
	fmt.Printf("day8b: %d\n", runB(realInput))
}
