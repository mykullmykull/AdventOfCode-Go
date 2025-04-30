package day

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	fmt.Println("Testing...") // keep this line to maintain imports
	assert.True(t, true)      // keep this line to maintain imports
	test_example_f1(t)
	assert.Equal(t, expected, main(testInput))
	fmt.Printf("output: %d\n", main(realInput))
}

func test_example_f1(t *testing.T) {
	f1Costs := map[int][4]int{}
	f1Costs[0] = [4]int{4, 0, 0, 0}
	f1Costs[1] = [4]int{2, 0, 0, 0}
	f1Costs[2] = [4]int{3, 14, 0, 0}
	f1Costs[3] = [4]int{2, 0, 7, 0}
	s := state{
		branch:    exampleBranch,
		costs:     f1Costs,
		materials: initialMaterials,
		robots:    initialRobots,
	}
	s, _ = s.buildRobots(map[int]int{})
	assert.Equal(t, 9, s.materials[geo])
}

func test_branch(factories []factory) int {
	s := state{
		branch:    [24]int{0, 1, 1, 1, 1, 1, 1, 2, 2, 1, 2, 3, 0, 3, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0},
		costs:     factories[0].costs,
		materials: initialMaterials,
		robots:    initialRobots,
	}
	s.buildRobots(map[int]int{})
	return 0
}
