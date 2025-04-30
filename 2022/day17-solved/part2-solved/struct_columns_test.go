package day

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getTestSpaces() columns {
	return columns{
		".######",
		"..#####",
		"...####",
		"....###",
		".....##",
		"......#",
		".......",
	}
}

func TestSpaces(t *testing.T) {
	fmt.Println("Testing Spaces...") // keep this line to maintain imports
	assert.True(t, true)             // keep this line to maintain imports
}

func Test_getSpace(t *testing.T) {
	columns := getTestSpaces()
	assert.Equal(t, 1, columns.getSpace(1))
}

func Test_getMin(t *testing.T) {
	columns := getTestSpaces()
	assert.Equal(t, 1, columns.getMinSpace())
}

func Test_addHeight(t *testing.T) {
	columns := getTestSpaces()
	columns = columns.addHeight(5)
	assert.Equal(t, 6, columns.getSpace(1))
	assert.Equal(t, 7, columns.getSpace(2))
	assert.Equal(t, 8, columns.getSpace(3))
	assert.Equal(t, 9, columns.getSpace(4))
	assert.Equal(t, 10, columns.getSpace(5))
	assert.Equal(t, 11, columns.getSpace(6))
	assert.Equal(t, 12, columns.getSpace(7))
}

func Test_markRock(t *testing.T) {
	columns := getTestSpaces()
	columns = columns.markRock([]int{3, 6})
	assert.Equal(t, 3, columns.getSpace(6))
}
