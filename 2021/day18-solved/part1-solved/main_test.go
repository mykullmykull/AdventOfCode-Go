package day

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	fmt.Println("Testing...") // keep this line to maintain imports
	assert.True(t, true)      // keep this line to maintain imports
	assert.Equal(t, hwM, main(hwT), "Main should match the homework test case")
	fmt.Printf("output: %d\n", main(realInput))
}

func TestSum(t *testing.T) {
	testCases := []struct {
		name     string
		testData []string
		expected string
	}{
		{"sum #1", sum1T, sum1E},
		{"sum #2", sum2T, sum2E},
		{"sum #3", sum3T, sum3E},
		{"sum #4", sum4T, sum4E},
		{"hw", hwT, hwS},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, sum(tc.testData), "Sum should match")
		})
	}
}

func TestReduce(t *testing.T) {
	testCases := []struct {
		name     string
		testData []string
		expected string
	}{
		{"steps", stepT, stepE},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, reduce(tc.testData), "Reduction should match")
		})
	}
}

func TestGetMagnitude(t *testing.T) {
	testCases := []struct {
		name     string
		testData string
		expected int
	}{
		{"magnitude #1", mag1T, mag1E},
		{"magnitude #2", mag2T, mag2E},
		{"magnitude #3", mag3T, mag3E},
		{"magnitude #4", mag4T, mag4E},
		{"magnitude #5", mag5T, mag5E},
		{"magnitude #6", mag6T, mag6E},
		{"magnitude #7", mag7T, mag7E},
		{"magnitude #8", mag8T, mag8E},
		{"magnitude #9", mag9T, mag9E},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, getMagnitude(tc.testData), "Magnitude should match")
		})
	}
}

func TestSplit(t *testing.T) {
	testCases := []struct {
		name     string
		testData string
		expected string
	}{
		{"split #1", split1T, split1E},
		{"split #2", split2T, split2E},
		{"split #3", split3T, split3E},
		{"split #4", split4T, split4E},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, split(tc.testData), "Split should match")
		})
	}
}

func TestExplode(t *testing.T) {
	testCases := []struct {
		name     string
		testData string
		expected string
	}{
		{"explode #1", explode1T, explode1E},
		{"explode #2", explode2T, explode2E},
		{"explode #3", explode3T, explode3E},
		{"explode #4", explode4T, explode4E},
		{"explode #5", explode5T, explode5E},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, explode(tc.testData), "Explosion should match")
		})
	}
}

func TestAdd(t *testing.T) {
	testCases := []struct {
		name     string
		testData []string
		expected string
	}{
		{"add #1", add1T, add1E},
		{"add #2", add2T, add2E},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, add(tc.testData), "Addition should match")
		})
	}
}
