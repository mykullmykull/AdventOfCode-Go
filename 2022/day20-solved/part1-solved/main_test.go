package day

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	fmt.Println("Testing...") // keep this line to maintain imports
	assert.True(t, true)      // keep this line to maintain imports
	assert.Equal(t, expected, main(testInput))
	fmt.Printf("output: %d\n", main(realInput))
}

func Test_getNewX(t *testing.T) {
	println("Testing wrap...")
	// [0, -10, 2, 3]
	// [0, 2, 3]
	// [-10, 0, 2, 3]
	assert.Equal(t, 0, getNewX(1, -10, 4))
	assert.Equal(t, 1, getNewX(2, 15, 5))
	// [0, 2, 15, 3, 4]
	// [0, 2, 3, 4]
	// [15, 0, 2, 3, 4]

}
