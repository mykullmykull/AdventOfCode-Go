package day4

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_runA(t *testing.T) {
	assert.Equal(t, ea, runA(i1))
	fmt.Printf("a: %d\n", runA(real))
}

func Test_runB(t *testing.T) {
	// assert.Equal(t, eb, runB(i1))
	fmt.Print("b: ")
	runB(real)
	fmt.Println()
}
