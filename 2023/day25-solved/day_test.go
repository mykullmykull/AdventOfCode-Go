package day25

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_runA(t *testing.T) {
	assert.Equal(t, expected, runA(testInput))
	fmt.Printf("day25a: %d\n", runA(realInput))
}

// func Test_runB(t *testing.T) {
// 	assert.Equal(t, expected2, runB(testInput))
// 	fmt.Printf("day25b: %d\n", runB(realInput))
// }
