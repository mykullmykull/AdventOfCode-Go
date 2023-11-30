package day11FirstTry

import (
	"fmt"
	"testing"
)

func Test_runA(t *testing.T) {
	// assert.Equal(t, expected1, runA(test1))
	fmt.Printf("a: %d\n", runA(real))
}

// func Test_runB(t *testing.T) {
// 	assert.Equal(t, expected1, runA(test1, 2, 5))
// 	fmt.Printf("b: %d\n", runB(real))
// }

// func Test_duplicates(t *testing.T) {
// 	type testCase struct {
// 		s1          Snapshot
// 		s2          Snapshot
// 		isDuplicate bool
// 	}

// 	hyG := Component{element: "hydrogen", isGenerator: true, floor: 1}
// 	hyM := Component{element: "hydrogen", isGenerator: false, floor: 1}
// 	liG := Component{element: "lithium", isGenerator: true, floor: 1}
// 	liM := Component{element: "lithium", isGenerator: false, floor: 1}
// 	liM2 := Component{element: "lithium", isGenerator: false, floor: 2}

// 	testCases := []testCase{
// 		// same elevator, same components
// 		{
// 			s1:          Snapshot{elevator: 1, components: []Component{hyG, hyM, liG, liM}},
// 			s2:          Snapshot{elevator: 1, components: []Component{hyG, hyM, liG, liM}},
// 			isDuplicate: true,
// 		},

// 		// different elevator, same components
// 		{
// 			s1:          Snapshot{elevator: 1, components: []Component{hyG, hyM, liG, liM}},
// 			s2:          Snapshot{elevator: 2, components: []Component{hyG, hyM, liG, liM}},
// 			isDuplicate: false,
// 		},

// 		// same elevator, different components
// 		{
// 			s1:          Snapshot{elevator: 1, components: []Component{hyG, hyM, liG, liM}},
// 			s2:          Snapshot{elevator: 1, components: []Component{hyG, hyM, liG, liM2}},
// 			isDuplicate: false,
// 		},
// 	}

// 	for _, testCase := range testCases {
// 		result := testCase.s1.isDuplicate([]Snapshot{testCase.s2}) == testCase.isDuplicate
// 		if !result {
// 			fmt.Printf("\ntest case failed: %+v\n", testCase)
// 		}
// 		assert.True(t, result)
// 	}
// }
