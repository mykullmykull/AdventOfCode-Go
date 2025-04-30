package day

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimulator(t *testing.T) {
	fmt.Println("Testing Simulator...") // keep this line to maintain imports
	assert.True(t, true)                // keep this line to maintain imports
}

func getTestSimulator() simulator {
	winds := ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"
	st := state{
		towerCols: columns{},
		rockIndex: -1,
		windIndex: -1,
	}

	sim := simulator{
		cache:  make(map[state]counters),
		counts: counters{},
		state:  st,
		winds:  winds,
	}
	sim = sim.getRocks()
	return sim
}

func Test_getRock(t *testing.T) {
	sim := getTestSimulator()
	type testCase struct {
		index    int
		expected rock
	}
	testCases := []testCase{
		{0, rock{r1}},
		{1, rock{r2}},
		{2, rock{r3}},
		{3, rock{r4}},
		{4, rock{r5}},
	}

	for _, tc := range testCases {
		sim.state.rockIndex = tc.index
		actual := sim.getNextRock().rock
		for x, expectedPoint := range tc.expected.points {
			actualPoint := actual.points[x]
			for y, expectedValue := range expectedPoint {
				actualValue := actualPoint[y]
				assert.Equal(t, expectedValue, actualValue)
			}
		}
	}
}

func Test_getNextWind(t *testing.T) {
	sim := getTestSimulator()
	for x := 0; x < 10; x++ {
		sim = sim.getNextWind()
		expectedIndex := x % len(sim.winds)
		expectedWind := string(sim.winds[expectedIndex])
		assert.Equal(t, expectedIndex, sim.state.windIndex)
		assert.Equal(t, expectedWind, sim.wind)
	}
}

func Test_addRockToTop(t *testing.T) {
	sim := getTestSimulator()
	sim.state.rockIndex = 3
	sim = sim.getNextRock()
	sim.counts.height = 11
	oldHeight := sim.counts.height
	sim.state.towerCols = columns{
		".######",
		"..#####",
		"...####",
		"....###",
		".....##",
		"......#",
		".......",
	}
	oldTowerCols := sim.state.towerCols
	expectedAddition := 4
	sim = sim.addRockToTop()
	for x := 1; x <= 7; x++ {
		assert.Equal(t, sim.state.towerCols.getSpace(x)-oldTowerCols.getSpace(x), expectedAddition)
	}
	assert.Equal(t, sim.counts.height-oldHeight, expectedAddition)
}

func Test_addDropsAndHeight(t *testing.T) {
	sim := getTestSimulator()
	sim.counts.droppedCount = 3
	sim.counts.height = 7
	dropsPerRepeat := 4
	heightPerRepeat := 6
	dropsMax := 28
	sim = sim.addDropsAndHeight(dropsPerRepeat, heightPerRepeat, dropsMax)
	// repeats = 6
	assert.Equal(t, sim.counts.droppedCount, 27)
	assert.Equal(t, sim.counts.height, 43)
}

func Test_rockCanStillDrop(t *testing.T) {
	sim := getTestSimulator()
	type testCase struct {
		rockIndex int
		expected  bool
	}
	testCases := []testCase{
		{0, false},
		{1, true},
		{2, false},
		{3, true},
		{4, true},
	}

	for _, tc := range testCases {
		sim.state.rockIndex = tc.rockIndex
		sim = sim.getNextRock()
		sim.state.towerCols = columns{
			"##",
			"##",
			".#",
			".#",
			"##",
			"##",
			"##",
		}
		sim = sim.addRockToTop()
		actual := sim.rockCanStillDrop()
		assert.Equal(t, tc.expected, actual)
	}
}

func Test_addSettledRock_at_top(t *testing.T) {
	sim := getTestSimulator()
	sim.rock = rock{
		points: [][]int{
			{0, 1},
			{1, 1},
			{1, 2},
		},
	}
	sim.state.towerCols = columns{
		"..####",
		"...###",
		".#####",
		"..####",
		"...###",
		"....##",
		".....#",
	}
	oldTowerCols := sim.state.towerCols
	sim.counts.height = 27
	sim = sim.addSettledRock()
	assert.Equal(t, 0, sim.state.towerCols.getSpace(1))
	assert.Equal(t, 1, sim.state.towerCols.getSpace(2))
	assert.Equal(t, oldTowerCols.col3, sim.state.towerCols.col3)
	assert.Equal(t, oldTowerCols.col4, sim.state.towerCols.col4)
	assert.Equal(t, oldTowerCols.col5, sim.state.towerCols.col5)
	assert.Equal(t, oldTowerCols.col6, sim.state.towerCols.col6)
	assert.Equal(t, oldTowerCols.col7, sim.state.towerCols.col7)
}

func Test_addSettledRock_below_top(t *testing.T) {
	sim := getTestSimulator()
	sim.rock = rock{
		points: [][]int{
			{2, 1},
			{3, 1},
			{3, 2},
		},
	}
	sim.state.towerCols = columns{
		"....##",
		".....#",
		"######",
		".#####",
		"..####",
		"...###",
		"....##",
	}
	oldTowerCols := sim.state.towerCols
	sim = sim.addSettledRock()
	assert.Equal(t, 2, sim.state.towerCols.getSpace(1))
	assert.Equal(t, 3, sim.state.towerCols.getSpace(2))
	assert.Equal(t, oldTowerCols.col3, sim.state.towerCols.col3)
	assert.Equal(t, oldTowerCols.col4, sim.state.towerCols.col4)
	assert.Equal(t, oldTowerCols.col5, sim.state.towerCols.col5)
	assert.Equal(t, oldTowerCols.col6, sim.state.towerCols.col6)
	assert.Equal(t, oldTowerCols.col7, sim.state.towerCols.col7)
}

func Test_dropRock_below_top(t *testing.T) {
	sim := getTestSimulator()
	sim.rock = rock{
		points: [][]int{
			{2, 1},
			{3, 1},
			{3, 2},
		},
	}
	sim.state.towerCols = columns{
		".......##",
		".......##",
		"#########",
		".########",
		"..#######",
		"...######",
		"....#####",
	}
	oldTowerCols := sim.state.towerCols
	expectedRock := rock{
		points: [][]int{
			{3, 1},
			{4, 1},
			{4, 2},
		},
	}
	sim = sim.dropRock()
	assert.Equal(t, expectedRock, sim.rock)
	assert.Equal(t, oldTowerCols.col1, sim.state.towerCols.col1)
	assert.Equal(t, oldTowerCols.col2, sim.state.towerCols.col2)
	assert.Equal(t, oldTowerCols.col3, sim.state.towerCols.col3)
	assert.Equal(t, oldTowerCols.col4, sim.state.towerCols.col4)
	assert.Equal(t, oldTowerCols.col5, sim.state.towerCols.col5)
	assert.Equal(t, oldTowerCols.col6, sim.state.towerCols.col6)
	assert.Equal(t, oldTowerCols.col7, sim.state.towerCols.col7)
}

func Test_dropNextRock_above_top(t *testing.T) {
	sim := getTestSimulator()
	sim.rock = rock{
		points: [][]int{
			{0, 1},
			{1, 1},
			{1, 2},
		},
	}
	sim.state.towerCols = columns{
		".......##",
		".......##",
		".########",
		"..#######",
		"...######",
		"....#####",
		".....####",
	}
	expectedRock := sim.rock
	sim = sim.dropRock()
	assert.Equal(t, expectedRock, sim.rock)
	assert.Equal(t, 6, sim.state.towerCols.getSpace(1))
	assert.Equal(t, 6, sim.state.towerCols.getSpace(2))
	assert.Equal(t, 0, sim.state.towerCols.getSpace(3))
	assert.Equal(t, 1, sim.state.towerCols.getSpace(4))
	assert.Equal(t, 2, sim.state.towerCols.getSpace(5))
	assert.Equal(t, 3, sim.state.towerCols.getSpace(6))
	assert.Equal(t, 4, sim.state.towerCols.getSpace(7))
}

func Test_dropNextRock(t *testing.T) {
	sim := getTestSimulator()
	sim = sim.getRocks()
	assert.Equal(t, 5, len(sim.rocks))
	assert.Equal(t, 0, sim.counts.droppedCount)
	assert.Equal(t, 0, sim.counts.height)

	expectedHeights := []int{1, 4, 6, 7, 9, 10, 13, 15, 17}
	for x := 0; x < 9; x++ {
		sim = sim.dropNextRock()
		assert.Equal(t, x+1, sim.counts.droppedCount)
		assert.Equal(t, expectedHeights[x], sim.counts.height)
	}
}

func Test_dropRocksUntilRepeat(t *testing.T) {
	sim := getTestSimulator()
	sim = sim.getRocks()
	sim = sim.dropRocksUntilRepeat()
	assert.Equal(t, 63, sim.counts.droppedCount)
	assert.Equal(t, 102, sim.counts.height)
}

func Test_dropRocksUntilCount(t *testing.T) {
	sim := getTestSimulator()
	sim = sim.getRocks()
	sim = sim.dropRocksUntilCount(101)
	assert.Equal(t, 101, sim.counts.droppedCount)
}
