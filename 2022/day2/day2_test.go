package day2

import "testing"

func TestFinalScoreB(t *testing.T) {
	input := test
	expected := 12
	result := FinalScoreB(input)
	if result != expected {
		t.Fatalf(`TestFinalScoreB after test input should be %v but instead was %v`, expected, result)
	}
}

func TestFinalScore(t *testing.T) {
	input := test
	expected := 15
	result := FinalScore(input)
	if result != expected {
		t.Fatalf(`TestFinalScore after test input should be %v but instead was %v`, expected, result)
	}
}
