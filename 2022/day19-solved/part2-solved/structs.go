package day

import (
	"fmt"
	"goAoc2022/utils"
	"strings"
)

const (
	ore = iota
	cly
	obs
	geo
)

var initialMaterials = [4]int{0, 0, 0, 0}
var initialRobots = [4]int{1, 0, 0, 0}
var badBranch = [32]int{-1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
var exampleBranch = [32]int{1, 1, 1, 2, 1, 2, 3, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type factory struct {
	id    int
	costs map[int][4]int
}

type state struct {
	branch    [32]int        // index of next robot = total robots - 1, int = type
	time      int            // current time
	materials [4]int         // index = type
	robots    [4]int         // index = type
	costs     map[int][4]int // key = type, value.index = type, value.value = cost
}

func (f factory) getMaxGeos() int {
	maxGeos := 0
	cache := map[int]int{} // key = time, value = maxPossibleGeos
	s := f.initializeState(badBranch)
	// can't build obs or geo robot before cly
	count := 0
	for s.branch[0] < 2 {
		// for s.branch[4] < 1 { // fastest so far = 7.385s
		// for s.branch[3] < 1 { // fastest so far = 22.886s
		count++
		if count%100000 == 0 {
			println("branch:", s.branchToString())
		}
		s, cache = s.buildRobots(cache)
		maxGeos = utils.UpdateMost(maxGeos, s.materials[geo])
		s = s.resetWithNextBranch()
	}
	println("branch:", s.branchToString())
	return maxGeos
}

func (f factory) initializeState(branch [32]int) state {
	return state{costs: f.costs, branch: branch}.resetWithNextBranch()
}

func (s state) buildRobots(cache map[int]int) (state, map[int]int) {
	for s.time < 32 {
		possibleGeos := s.getPossibleGeos()
		if cache[s.time] > possibleGeos {
			return s, cache
		}
		cache[s.time] = possibleGeos
		maxTime := s.getMaxTime()
		if s.time+maxTime > 32 {
			return s.mineGeosToTheEnd(), cache
		}
		s = s.update(maxTime)
	}
	return s, cache
}

func (s state) getPossibleGeos() int {
	timeLeft := 32 - s.time
	robotsMining := s.robots[geo]
	fromOldRobots := robotsMining * timeLeft
	return s.materials[geo] + fromOldRobots
}

func (s state) willNotBeatMax(maxGeos int) bool {
	if !strings.Contains(s.branchToString(), "3") {
		return true
	}

	robotsMining := s.robots[geo]
	timeLeft := 32 - s.time
	fromOldRobots := robotsMining * timeLeft
	fromNewRobots := 0
	for x := 0; x < timeLeft; x++ {
		fromNewRobots += x
	}
	possibleGeos := s.materials[geo] + fromOldRobots + fromNewRobots
	return possibleGeos <= maxGeos
}

func (s state) getMaxTime() int {
	_, costs := s.getNextRobotAndCost()
	maxTime := 1
	for x := ore; x <= geo; x++ {
		stillNeeded := costs[x] - s.materials[x]
		if stillNeeded <= 0 {
			continue
		}
		robotsMining := s.robots[x]
		if robotsMining == 0 {
			return 32
		}
		time := utils.DivideAndRoundUp(stillNeeded, robotsMining) + 1
		maxTime = utils.UpdateMost(maxTime, time)
	}
	return maxTime
}

func (s state) getNextRobotAndCost() (int, [4]int) {
	nextIndex := s.getNextIndexToBuild()
	if nextIndex == 32 {
		nextIndex = 0
	}
	nextRobot := s.branch[nextIndex]
	costs := s.costs[nextRobot]
	return nextRobot, costs
}

func (s state) mineGeosToTheEnd() state {
	robotsMining := s.robots[geo]
	timeLeft := 32 - s.time
	s.materials[geo] += robotsMining * timeLeft
	s.time = 32
	return s
}

func (s state) update(timeToAdd int) state {
	nextRobot, costs := s.getNextRobotAndCost()
	newMaterials := [4]int{s.materials[0], s.materials[1], s.materials[2], s.materials[3]}
	for x := ore; x <= geo; x++ {
		toAdd := (s.robots[x] * timeToAdd) - costs[x]
		newMaterials[x] += toAdd
	}
	s.materials = newMaterials
	s.robots[nextRobot]++
	s.time += timeToAdd
	return s
}

func (s state) updateWithPanic(timeToAdd int) state {
	nextRobot, costs := s.getNextRobotAndCost()
	newMaterials := [4]int{s.materials[0], s.materials[1], s.materials[2], s.materials[3]}
	for x := ore; x <= geo; x++ {
		// mine materials until the start of the last minute
		toAdd := s.robots[x] * (timeToAdd - 1)
		newMaterials[x] += toAdd

		// start building the robot, and verify that we have enough
		toAdd = -1 * costs[x]
		newMaterials[x] += toAdd
		if newMaterials[x] < 0 {
			panic("negative materials of type " + fmt.Sprint(x))
		}

		// mine materials for the last minute
		toAdd = s.robots[x]
		newMaterials[x] += toAdd
	}
	s.materials = newMaterials
	s.robots[nextRobot]++
	s.time += timeToAdd
	return s
}

func (s state) resetWithNextBranch() state {
	indexToIncrement := s.getNextIndexToBuild() - 1
	s = s.incrementBranch(indexToIncrement)
	s.materials = initialMaterials
	s.robots = initialRobots
	s.time = 0
	return s
}

func (s state) getNextIndexToBuild() int {
	countRobots := s.robots[ore] + s.robots[cly] + s.robots[obs] + s.robots[geo]
	if countRobots < 2 {
		return 32
	}
	return countRobots - 1
}

func (s state) incrementBranch(x int) state {
	newBranch := [32]int{}
	if s.branch[0] == badBranch[0] {
		s.branch = newBranch
		return s
	}

	for y := 0; y < 32; y++ {
		newBranch[y] = s.branch[y]
	}

	newBranch[x]++
	s.branch = newBranch

	if s.branch[x] > geo && x == 0 {
		return s
	}

	if s.branch[x] > geo {
		s.branch[x] = 0
		return s.incrementBranch(x - 1)
	}
	x++
	for x < 32 {
		s.branch[x] = 0
		x++
	}
	return s
}
