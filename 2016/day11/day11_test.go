package day11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_runA(t *testing.T) {
	assert.Equal(t, expected1, runA(test1))
	// fmt.Printf("a: %d\n", runA(real, 17, 61))
}

// func Test_runB(t *testing.T) {
// 	assert.Equal(t, expected1, runA(test1, 2, 5))
// 	fmt.Printf("b: %d\n", runB(real))
// }
