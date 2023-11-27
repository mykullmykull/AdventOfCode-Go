package day9

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_runA(t *testing.T) {
	assert.Equal(t, expected1, runA(test1))
	assert.Equal(t, expected2, runA(test2))
	assert.Equal(t, expected3, runA(test3))
	assert.Equal(t, expected4, runA(test4))
	assert.Equal(t, expected5, runA(test5))
	fmt.Printf("a: %d\n", runA(real))
}

func Test_runB(t *testing.T) {
	assert.Equal(t, expected3, runB(test3))
	assert.Equal(t, expected5b, runB(test5))
	assert.Equal(t, expected6, runB(test6))
	assert.Equal(t, expected7, runB(test7))
	fmt.Printf("b: %d\n", runB(real))
}
