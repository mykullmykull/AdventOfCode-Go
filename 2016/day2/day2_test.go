package day2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_runA(t *testing.T) {
	assert.Equal(t, ea, runA(i1))
	fmt.Printf("a: %s\n", runA(real))
}

func Test_runB(t *testing.T) {
	assert.Equal(t, eb, runB(i1))
	fmt.Printf("b: %s\n", runB(real))
}
