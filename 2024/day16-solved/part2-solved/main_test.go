package day

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	fmt.Println("Testing...") // keep this line to maintain imports
	assert.True(t, true)      // keep this line to maintain imports
	assert.Equal(t, expected1, main(test1, score1))
	assert.Equal(t, expected2, main(test2, score2))
	fmt.Printf("output: %d\n", main(realInput, realScore))
}
