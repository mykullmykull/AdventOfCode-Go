package day17

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_runA(t *testing.T) {
	assert.Equal(t, expected, runA(testInput))
	fmt.Printf("day17a: %d\n", runA(realInput))
}

func Test_runB(t *testing.T) {
	assert.Equal(t, expected2, runB(testInput))
	assert.Equal(t, expected3, runB(testInput3))
	fmt.Printf("day17b: %d\n", runB(realInput))
}
