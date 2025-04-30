package day

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRock(t *testing.T) {
	fmt.Println("Testing Rock...") // keep this line to maintain imports
	assert.True(t, true)           // keep this line to maintain imports
}

func getRockThatCanBlow() rock {
	return rock{[][]int{
		{2, 3},
		{3, 4},
		{4, 5},
		{5, 6},
	}}
}

func getRockThatCannotBlow() rock {
	return rock{[][]int{
		{0, 1},
		{6, 7},
	}}
}

func Test_getHeight(t *testing.T) {
	testRock := getRockThatCanBlow()
	assert.Equal(t, 4, testRock.getHeight())
}

func Test_getMinRow(t *testing.T) {
	testRock := getRockThatCanBlow()
	assert.Equal(t, 2, testRock.getMinRow(-1))
	assert.Equal(t, 2, testRock.getMinRow(3))
	assert.Equal(t, 3, testRock.getMinRow(4))
	assert.Equal(t, 4, testRock.getMinRow(5))
	assert.Equal(t, 5, testRock.getMinRow(6))
}

func Test_blowRockUnappy(t *testing.T) {
	cannotBlow := getRockThatCannotBlow()
	blown := cannotBlow.blowRock("<")
	for x, newPoint := range blown.points {
		oldPoint := cannotBlow.points[x]

		// rows are unchanged
		assert.Equal(t, oldPoint[0], newPoint[0])

		// columns are unchanged
		assert.Equal(t, oldPoint[1], newPoint[1])
	}

	blown = cannotBlow.blowRock(">")
	for x, newPoint := range blown.points {
		oldPoint := cannotBlow.points[x]

		// rows are unchanged
		assert.Equal(t, oldPoint[0], newPoint[0])

		// columns are unchanged
		assert.Equal(t, oldPoint[1], newPoint[1])
	}
}

func Test_blowRockHappy(t *testing.T) {
	canBlow := getRockThatCanBlow()
	blown := canBlow.blowRock("<")
	for x, newPoint := range blown.points {
		oldPoint := canBlow.points[x]

		// rows are unchanged
		assert.Equal(t, oldPoint[0], newPoint[0])

		// columns are decremented by 1
		assert.Equal(t, oldPoint[1]-1, newPoint[1])
	}

	blown = canBlow.blowRock(">")
	for x, newPoint := range blown.points {
		oldPoint := canBlow.points[x]

		// rows are unchanged
		assert.Equal(t, oldPoint[0], newPoint[0])

		// columns are incremented by 1
		assert.Equal(t, oldPoint[1]+1, newPoint[1])
	}
}

func Test_drop(t *testing.T) {
	canMove := getRockThatCanBlow()
	dropped := canMove.drop()
	for x, newPoint := range dropped.points {
		oldPoint := canMove.points[x]

		// rows are incremented by 1
		assert.Equal(t, oldPoint[0]+1, newPoint[0])

		// columns are unchanged
		assert.Equal(t, oldPoint[1], newPoint[1])
	}
}
