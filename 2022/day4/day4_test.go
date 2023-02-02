package day4

import "testing"

func TestCountOverlapsPart2(t *testing.T) {
	input := test
	expected := 4
	result := CountOverlaps(input, 2)
	if result != expected {
		t.Fatalf(`CountOverlaps part 2 from test should be %v but instead returned %v`, expected, result)
	}
}

func TestCountOverlapsPart1(t *testing.T) {
	input := test
	expected := 2
	result := CountOverlaps(input, 1)
	if result != expected {
		t.Fatalf(`CountOverlaps part 1 from test should be %v but instead returned %v`, expected, result)
	}
}
