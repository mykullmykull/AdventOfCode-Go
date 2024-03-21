package day21

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_runA(t *testing.T) {
	// assert.Equal(t, expected, runA(testInput, testSteps))
	// fmt.Printf("day1a: %d\n", runA(realInput, realSteps))
}

func Test_runB(t *testing.T) {
	assert.Equal(t, e2, runB(testInput, s2))
	assert.Equal(t, e3, runB(testInput, s3))
	assert.Equal(t, e4, runB(testInput, s4))
	assert.Equal(t, e5, runB(testInput, s5))
	assert.Equal(t, e6, runB(testInput, s6))
	// assert.Equal(t, e7, runB(testInput, s7))
	// assert.Equal(t, e8, runB(testInput, s8))
	// fmt.Printf("day1b: %d\n", runB(realInput, realSteps))
}
