package day6

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func Test_runA(t *testing.T) {
// 	assert.Equal(t, expected1, runA(test1))
// 	fmt.Printf("a: %s\n", runA(real))
// }

func Test_runB(t *testing.T) {
	assert.Equal(t, expected2, runB(test1))
	fmt.Printf("b: %s\n", runB(real))
}
