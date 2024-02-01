package day3

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_runA(t *testing.T) {
	assert.Equal(t, expected, runA(testInput))
	fmt.Printf("day3a: %d\n", runA(realInput))
}

func Test_runB(t *testing.T) {
	assert.Equal(t, expected2, runB(testInput))
	fmt.Printf("day3b: %d\n", runB(realInput))
}
