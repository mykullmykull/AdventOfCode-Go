package day1

import "testing"

func TestTopThreeCalories(t *testing.T) {
	calorieStrs := test
	expect := 45000
	result := TopThreeCalories(calorieStrs)
	if result != expect {
		t.Fatalf(`TopThreeCalories from test strings should be %v but instead returned %v`, expect, result)
	}
}

func TestMaxCalories(t *testing.T) {
	calorieStrs := test
	expect := 24000
	result := MaxCalories(calorieStrs)
	if result != expect {
		t.Fatalf(`MaxCalories from Test strings should be %v but instead returned %v`, expect, result)
	}
}
