package day

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	fmt.Println("Testing...") // keep this line to maintain imports
	assert.True(t, true)      // keep this line to maintain imports
	assert.Equal(t, 32, main(testInput, 50, true))
	assert.Equal(t, 31, main(testInput, 52, true))
	assert.Equal(t, 29, main(testInput, 54, true))
	assert.Equal(t, 39, main(testInput, 56, true))
	assert.Equal(t, 25, main(testInput, 58, true))
	assert.Equal(t, 23, main(testInput, 60, true))
	assert.Equal(t, 20, main(testInput, 62, true))
	assert.Equal(t, 19, main(testInput, 64, true))
	assert.Equal(t, 12, main(testInput, 66, true))
	assert.Equal(t, 14, main(testInput, 68, true))
	assert.Equal(t, 12, main(testInput, 70, true))
	assert.Equal(t, 22, main(testInput, 72, true))
	assert.Equal(t, 4, main(testInput, 74, true))
	assert.Equal(t, 3, main(testInput, 76, true))
	assert.Equal(t, 285, main(testInput, 50, false))
	fmt.Printf("output: %d\n", main(realInput, 100, false))
	// 888115 is too low
}
