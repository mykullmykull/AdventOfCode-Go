package day

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExplodePair(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		index    int
		expected string
	}{
		{"simple pair", "[1,2]", 0, "0"},
		{"left pair", "[[1,2],[3,4]]", 1, "[0,[5,4]]"},
		{"middle pair", "[[[1,2],[3,1]],[4,5]]", 8, "[[[1,5],0],[5,5]]"},
		{"right pair", "[3,[6,5]]", 3, "[9,0]"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, explodePair(tc.input, tc.index))
		})
	}
}

func TestSplitPairFromRight(t *testing.T) {
	testCases := []struct {
		name          string
		input         string
		expectedPair  string
		expectedRight string
	}{
		{"simple pair", "1,2]", "1,2", ""},
		{"nested pair", "1,2],[3,4]]", "1,2", ",[3,4]]"},
		{"complex pair", "2,3]],1],[4,5]]", "2,3", "],1],[4,5]]"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			pair, right := splitPairFromRight(tc.input)
			assert.Equal(t, tc.expectedPair, pair)
			assert.Equal(t, tc.expectedRight, right)
		})
	}
}

func TestAddToLeft(t *testing.T) {
	testCases := []struct {
		name     string
		leftStr  string
		val      int
		expected string
	}{
		{"add 5 to 4", "[[1, [3, 4]],[", 5, "[[1, [3, 9]],["},
		{"add 1 to 2", "[2", 1, "[3"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, addToLeft(tc.leftStr, tc.val), "adding to left should match")
		})
	}
}

func TestAddToRight(t *testing.T) {
	testCases := []struct {
		name     string
		rightStr string
		val      int
		expected string
	}{
		{"add 1 to 2", "]],2],4],", 1, "]],3],4],"},
		{"add 5 to 4", "]]],4],[1,2]]", 5, "]]],9],[1,2]]"},
		{"add 3 to 0", "],[0,1]]", 3, "],[3,1]]"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, addToRight(tc.rightStr, tc.val), "adding to right should match")
		})
	}
}
