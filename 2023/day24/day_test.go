package day24

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_runA(t *testing.T) {
	assert.Equal(t, expected, runA(testInput, testMin, testMax))
	fmt.Printf("day24a: %f\n", runA(realInput, realMin, realMax))
}
