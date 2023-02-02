package day5

import "testing"

func TestTopCratesAfterMovesPart1(t *testing.T) {
	crates := testCrates
	moves := testMoves
	expected := "CMZ"
	result := TopCratesAfterMoves(crates, moves, 1)
	if result != expected {
		t.Fatalf(`TopCratesAfterMoves from part 1 test should be %v but instead returned %v`, expected, result)
	}
}

func TestTopCratesAfterMovesPart2(t *testing.T) {
	crates := testCrates
	moves := testMoves
	expected := "MCD"
	result := TopCratesAfterMoves(crates, moves, 2)
	if result != expected {
		t.Fatalf(`TopCratesAfterMoves from part 2 test should be %v but instead returned %v`, expected, result)
	}
}
