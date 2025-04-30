package day

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getTestBranch() [24]int {
	return [24]int{0, 1, 2, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
}

func getTestState() state {
	testCost := map[int][4]int{
		ore: {2, 0, 0, 0},
		cly: {3, 0, 0, 0},
		obs: {4, 2, 0, 0},
		geo: {5, 0, 3, 0},
	}
	return state{
		branch:    getTestBranch(),
		materials: [4]int{1, 2, 3, 4},
		robots:    [4]int{4, 3, 2, 1},
		costs:     testCost,
	}
}

func Test_Structs(t *testing.T) {
	fmt.Println("Testing Structs...") // keep this line to maintain imports
	assert.True(t, true)              // keep this line to maintain imports
}

func Test_getMaxTime(t *testing.T) {
	s := getTestState()

	// Test with 4 robots
	s.robots = [4]int{2, 0, 2, 0}
	time := s.getMaxTime()
	assert.Equal(t, 3, time)
}

func Test_getNextRobotAndCost(t *testing.T) {
	s := getTestState()

	// Test with 2 robots
	s.robots = [4]int{2, 0, 0, 0}
	robot, cost := s.getNextRobotAndCost()
	assert.Equal(t, 1, robot)
	assert.Equal(t, s.costs[1], cost)

	// Test with 3 robots
	s.robots = [4]int{2, 1, 0, 0}
	robot, cost = s.getNextRobotAndCost()
	assert.Equal(t, 2, robot)
	assert.Equal(t, s.costs[2], cost)
}

func Test_mineGeosToTheEnd(t *testing.T) {
	s := getTestState()
	expectedGeos := s.materials[geo] + s.robots[geo]*24
	s = s.mineGeosToTheEnd()
	assert.Equal(t, expectedGeos, s.materials[geo])
	assert.Equal(t, 24, s.time)
}

func Test_update(t *testing.T) {
	s := getTestState()
	time := 7
	expectedOre := (s.materials[ore] + s.robots[ore]*time) - s.costs[ore][ore]
	expectedCly := (s.materials[cly] + s.robots[cly]*time) - s.costs[ore][cly]
	expectedObs := (s.materials[obs] + s.robots[obs]*time) - s.costs[ore][obs]
	expectedGeo := (s.materials[geo] + s.robots[geo]*time) - s.costs[ore][geo]
	expectedMaterials := [4]int{expectedOre, expectedCly, expectedObs, expectedGeo}
	expectedRobots := [4]int{5, 3, 2, 1}
	expectedTime := s.time + time

	s = s.update(time)
	assert.Equal(t, expectedMaterials, s.materials)
	assert.Equal(t, expectedRobots, s.robots)
	assert.Equal(t, expectedTime, s.time)
}

func Test_resetWithNextBranch(t *testing.T) {
	s := getTestState()
	s = s.resetWithNextBranch()
	assert.Equal(t, initialMaterials, s.materials)
	assert.Equal(t, initialRobots, s.robots)
}

func Test_getNextIndexToBuild(t *testing.T) {
	s := getTestState()
	x := s.getNextIndexToBuild()
	assert.Equal(t, 9, x)

	// Test with 0 robots
	s.robots = [4]int{}
	x = s.getNextIndexToBuild()
	assert.Equal(t, 24, x)

	// Test with initial robots
	s.robots = initialRobots
	x = s.getNextIndexToBuild()
	assert.Equal(t, 24, x)
}

func Test_incrementBranch(t *testing.T) {
	s := getTestState()

	// Test incrementing index 0
	s = s.incrementBranch(0)
	expected := []int{1, 0, 0, 0}
	assert.Equal(t, expected, s.branch[:4])

	// Test incrementing index 1
	s.branch = getTestBranch()
	s = s.incrementBranch(1)
	expected = []int{0, 2, 0, 0}
	assert.Equal(t, expected, s.branch[:4])

	// Test incrementing index 2
	s.branch = getTestBranch()
	s = s.incrementBranch(2)
	expected = []int{0, 1, 3, 0}
	assert.Equal(t, expected, s.branch[:4])

	// Test incrementing index 3 with maxed robot
	s.branch = getTestBranch()
	s = s.incrementBranch(3)
	expected = []int{0, 1, 3, 0}
	assert.Equal(t, expected, s.branch[:4])

	// Test incrementing with bad branch
	s.branch = badBranch
	s = s.incrementBranch(3)
	expected = []int{0, 0, 0, 0}
	assert.Equal(t, expected, s.branch[:4])

	// Test if first robot is geo
	s.branch = getTestBranch()
	s.branch[0] = geo
	s = s.incrementBranch(0)
	expected = []int{4, 1, 2, 3}
	assert.Equal(t, expected, s.branch[:4])
}
