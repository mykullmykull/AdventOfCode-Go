package day6

import "testing"

func TestFirstIndexAfterFourUnique(t *testing.T) {
	testFour(test1, 7, t)
	testFour(test2, 5, t)
	testFour(test3, 6, t)
	testFour(test4, 10, t)
	testFour(test5, 11, t)

	testFourteen(test1, 19, t)
	testFourteen(test2, 23, t)
	testFourteen(test3, 23, t)
	testFourteen(test4, 29, t)
	testFourteen(test5, 26, t)
}

func testFour(input string, expected int, t *testing.T) {
	result := FirstIndexAfterFourUnique(input)
	if result != expected {
		t.Fatalf(`FirstIndexAfterFourUnique from %v should be %v but instead returned %v`, input, expected, result)
	}
}

func testFourteen(input string, expected int, t *testing.T) {
	result := FirstIndexAfterFourteenUnique(input)
	if result != expected {
		t.Fatalf(`FirstIndexAfterFourteenUnique from %v should be %v but instead returned %v`, input, expected, result)
	}
}
