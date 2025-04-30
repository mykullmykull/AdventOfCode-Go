package day

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	fmt.Println("Testing...") // keep this line to maintain imports
	assert.True(t, true)      // keep this line to maintain imports
	// assert.Equal(t, expected, main(testInput))
	fmt.Printf("output: %s\n", main(realInput))
}

func runTest() {
	c := computer{0, 2024, 43690, 0, []int{4, 0}, ""}
	c = c.run()
	println(c.b)
}
