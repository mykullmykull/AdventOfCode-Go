package day21

import (
	"fmt"
	"testing"
)

func Test_runA(t *testing.T) {
	// assert.Equal(t, expected, runA(testInput, testSteps))
	// fmt.Printf("day1a: %d\n", runA(realInput, realSteps))

	// TEST MAP
	// fmt.Printf("with 4 steps: %d\n", runA(testInput, 4))
	// fmt.Printf("with 6 steps: %d\n", runA(testInput, 6))
	// fmt.Printf("with 11 steps: %d\n", runA(testInput, 11))
	// fmt.Printf("with 500 steps: %d\n", runA(testInput, 500)) // 167004 which matches AOC

	// REAL MAP
	// fmt.Printf("with 60 steps: %d\n", runA(realInput, 60)) // 3082
	// fmt.Printf("with 61 steps: %d\n", runA(realInput, 61)) // 3231
	// fmt.Printf("with 75 steps: %d\n", runA(realInput, 75)) // 5029
	// fmt.Printf("with 76 steps: %d\n", runA(realInput, 76)) // 5127
	// fmt.Printf("with 140 steps: %d\n", runA(realInput, 140)) // 17067
	// fmt.Printf("with 141 steps: %d\n", runA(realInput, 141)) // 17306
	// fmt.Printf("with 500 steps: %d\n", runA(realInput, 500)) // 216464
	// fmt.Printf("with 501 steps: %d\n", runA(realInput, 501)) // 217314
	// fmt.Printf("with 700 steps: %d\n", runA(realInput, 700)) // 423075
	fmt.Printf("with 701 steps: %d\n", runA(realInput, 701)) // 423914

}
