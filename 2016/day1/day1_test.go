package day1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_runA(t *testing.T) {
	assert.Equal(t, e1, runA(i1))
	assert.Equal(t, e2, runA(i2))
	assert.Equal(t, e3, runA(i3))
	fmt.Printf("day1a: %d\n", runA(real))
}

func Test_runB(t *testing.T) {
	assert.Equal(t, e4, runB(i4))
	fmt.Printf("day1b: %d\n", runB(real))
}
