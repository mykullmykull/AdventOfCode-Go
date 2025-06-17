package day

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	fmt.Println("Testing...") // keep this line to maintain imports
	assert.True(t, true)      // keep this line to maintain imports
	assert.Equal(t, expect1, main(test1))
	fmt.Printf("output: %d\n", main(realInput))
}

func TestGetMinVx(t *testing.T) {
	testCases := []struct {
		name     string
		p        probe
		expected int
	}{
		{"x between 20-30", probe{minX: 20, maxX: 30}, 6},
		{"x is 0", probe{minX: 0, maxX: 0}, 0},
		{"x between 1000 and 2000", probe{minX: 1000, maxX: 2000}, 45},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.p.getMinVx()
			assert.Equal(t, tc.expected, tc.p.velX, "Min Vx should match")
		})
	}
}

func TestGetMaxVy(t *testing.T) {
	testCases := []struct {
		input    string
		minVelX  int
		expected int
	}{
		{"target area: x=20..30, y=-10..-5", 6, 9},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			p := newProbe(tc.input)
			p.getMaxVy()
			p.velX = tc.minVelX
			assert.Equal(t, tc.expected, p.velY, "Max Vy should match")
		})
	}
}
