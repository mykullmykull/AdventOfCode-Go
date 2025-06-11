package day

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	fmt.Println("Testing...") // keep this line to maintain imports
	assert.True(t, true)      // keep this line to maintain imports
	// assert.Equal(t, expected1, main(test1))
	// assert.Equal(t, expected2, main(test2))
	// assert.Equal(t, expected3, main(test3))
	// assert.Equal(t, expected4, main(test4))
	// assert.Equal(t, expected5, main(test5))
	// assert.Equal(t, expected6, main(test6))
	// assert.Equal(t, expected7, main(test7))
	// assert.Equal(t, expected8, main(test8))
	fmt.Printf("output: %d\n", main(realInput))
}
