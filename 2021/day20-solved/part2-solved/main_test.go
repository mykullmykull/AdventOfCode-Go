package day

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	fmt.Println("Testing...") // keep this line to maintain imports
	assert.True(t, true)      // keep this line to maintain imports
	assert.Equal(t, expected, main(testInput, 50))
	fmt.Printf("output: %d\n", main(realInput, 50))
}
