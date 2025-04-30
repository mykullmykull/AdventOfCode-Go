package day

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Main(t *testing.T) {
	fmt.Println("Testing Main...") // keep this line to maintain imports
	assert.True(t, true)           // keep this line to maintain imports
	assert.Equal(t, expectedPart1Test, main(testInput, 2022))
	assert.Equal(t, expectedPart1Real, main(realInput, 2022))
	assert.Equal(t, expectedPart2Test, main(testInput, 1000000000000))
	fmt.Printf("part2 real output: %d\n", main(realInput, 1000000000000))
}
