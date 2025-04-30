package day

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	fmt.Println("Testing...") // keep this line to maintain imports
	assert.True(t, true)      // keep this line to maintain imports
	// assert.Equal(t, expected, main(testInput, testInstructions))
	fmt.Printf("output: %d\n", main(realInput, realInstructions))
	test_move(t)
}

func test_move(t *testing.T) {
	fmt.Println("Testing a particular move...") // keep this line to maintain imports
	b := board{
		grid:         realInput,
		instructions: realInstructions,
		p: point{
			row:   117,
			col:   4,
			side:  11,
			toRow: 0,
			toCol: -1,
		},
	}
	b.move("36")
	assert.Equal(t, -1, b.p.toCol)
}
