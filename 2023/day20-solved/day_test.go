package day20

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_runA(t *testing.T) {
	assert.Equal(t, expected1, runA(testInput1))
	assert.Equal(t, expected2, runA(testInput2))
	fmt.Printf("day1a: %d\n", runA(realInput))
}

func Test_runB(t *testing.T) {
	fmt.Printf("day1b: %d\n", runB(realInput))
}
