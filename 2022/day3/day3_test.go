package day3

import "testing"

func TestSumPrioritiesOfBadges(t *testing.T) {
	input := test
	expected := 70
	result := SumPrioritiesOfBadges(input)
	if result != expected {
		t.Fatalf(`SumPrioritiesOfBadges from test should be %v but instead returned %v`, expected, result)
	}
}

func TestSumPrioritiesOfMatchingTypes(t *testing.T) {
	input := test
	expected := 157
	result := SumPrioritiesOfMatchingTypes(input)
	if result != expected {
		t.Fatalf(`SumPrioritiesOfMatchingTypes from test should be %v but instead returned %v`, expected, result)
	}
}
