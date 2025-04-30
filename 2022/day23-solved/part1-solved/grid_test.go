package day

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Grid(t *testing.T) {
	fmt.Println("Testing Grid...") // keep this line to maintain imports
	assert.True(t, true)           // keep this line to maintain imports
}

func Test_expand(t *testing.T) {
	grid := []string{
		"#",
	}
	expected := []string{
		"...",
		".#.",
		"...",
	}
	assert.Equal(t, expected, expand(grid))
}

func Test_isEmpty(t *testing.T) {
	grid := []string{
		".....",
		"..#..",
		"..##.",
		".....",
	}
	assert.Equal(t, true, isEmpty(grid, 0, -1))
	assert.Equal(t, false, isEmpty(grid, 1, -1))
	assert.Equal(t, false, isEmpty(grid, 2, -1))
	assert.Equal(t, true, isEmpty(grid, 3, -1))
	assert.Equal(t, true, isEmpty(grid, -1, 0))
	assert.Equal(t, true, isEmpty(grid, -1, 1))
	assert.Equal(t, false, isEmpty(grid, -1, 2))
	assert.Equal(t, false, isEmpty(grid, -1, 3))
	assert.Equal(t, true, isEmpty(grid, -1, 4))
}

func Test_countEmpty(t *testing.T) {
	grid := []string{
		".....",
		"..#..",
		"..##.",
		".....",
	}
	expected := 17
	assert.Equal(t, expected, countEmpty(grid))
}

func Test_lookAround(t *testing.T) {
	grid := []string{
		".....",
		"..#..",
		"..##.",
		".....",
	}
	expected := []string{
		"#..",
		"##.",
		"...",
	}
	assert.Equal(t, expected, lookAround(grid, 2, 3))
}

func Test_shrink(t *testing.T) {
	grid := []string{
		".....",
		".....",
		"...#.",
		".....",
	}
	expected := []string{
		"#",
	}
	assert.Equal(t, expected, shrink(grid))
}
