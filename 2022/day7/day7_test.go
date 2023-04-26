package day7

import "testing"

func testSumofBigDirectories(input string, expected int, t *testing.T) {
	result := sumSizeOfBigDirectories(input)
	if result != expected {
		t.Fatalf(`sumSizeOfBigDirectories from %v should be %v but instead returned %v`, input, expected, result)
	}
}
