package day

import (
	"fmt"
	"goAoc2022/utils"
	"slices"
	"strings"
)

type valve struct {
	name           string
	flow           int
	adjacentValves []string
	targetsDist    map[string]int
}

type state struct {
	timeLeft     int
	position     string
	released     int
	targetsCount int
	openValves   int // the binary form of this int tracks which target valves are open (1) or closed (0)
}

func (v valve) getDistanceTo(path []string, name string, valves map[string]valve) int {
	if name == v.name {
		return 0
	}
	minDist := len(valves) + 100
	for _, adjacent := range v.adjacentValves {
		if adjacent == name {
			return 1
		}
		if utils.ContainsStr(path, adjacent) {
			continue
		}

		adjacentValve := valves[adjacent]
		adjacentDist := 1 + adjacentValve.getDistanceTo(append(path, adjacent), name, valves)
		minDist = utils.UpdateLeast(minDist, adjacentDist)
	}
	return minDist
}

func (s state) releaseMaxPressure(targets []string, nameToValve map[string]valve, cache map[state]int) (int, map[state]int) {
	maxReleased := 0
	if newlyReleased, isCached := s.getCached(cache); isCached {
		return s.released + newlyReleased, cache
	}

	unopenedTargets := s.getUnopenedTargets(targets)
	if len(unopenedTargets) == 0 {
		return s.released, cache
	}

	path := make([]int, len(unopenedTargets))
	for x := 0; x < len(path); x++ {
		path[x] = x
	}
	for len(path) > 0 {
		released, lastIndexOpened := s.getReleased(path, unopenedTargets, nameToValve)
		maxReleased = utils.UpdateMost(maxReleased, released)
		path = getNextPath(path, lastIndexOpened)
	}
	cache[s] = maxReleased
	return maxReleased, cache
}

func (s state) getUnopenedTargets(targets []string) []string {
	unopened := []string{}
	for i, target := range targets {
		if s.isValveOpen(i) {
			continue
		}
		unopened = append(unopened, target)
	}
	if len(unopened) > 15 {
		panic(fmt.Sprintf("too many unopened targets: %d", len(unopened)))
	}
	return unopened
}

func (s state) getReleased(path []int, targets []string, nameToValve map[string]valve) (int, int) {
	released := 0
	timeLeft := s.timeLeft
	start := nameToValve[s.position]
	lastIndexOpened := -1
	for x, p := range path {
		t := targets[p]
		v := nameToValve[t]
		dist := start.targetsDist[v.name]
		timeLeft = timeLeft - dist - 1
		if timeLeft <= 0 {
			lastIndexOpened = x
			break
		}
		released += v.flow * timeLeft
		start = v
	}
	return released, lastIndexOpened
}

func getNextPath(path []int, lastIndexOpened int) []int {
	indexToSwap := getIndexToSwap(path)
	if indexToSwap == -1 || (lastIndexOpened != -1 && lastIndexOpened < indexToSwap) {
		indexToSwap = lastIndexOpened
	}
	if indexToSwap == -1 {
		return []int{}
	}

	// add to the new path
	// the next digit greater than the digit to swap
	// and the remaining sorted digits
	digitToReplace, sortedAfter := getNextBiggestAndSortRemainder(indexToSwap, path)
	for digitToReplace < 0 {
		indexToSwap--
		if indexToSwap < 0 {
			return []int{}
		}
		digitToReplace, sortedAfter = getNextBiggestAndSortRemainder(indexToSwap, path)
	}

	// keep the digits before the index to swap
	newPath := append(append(path[:indexToSwap], digitToReplace), sortedAfter...)
	return newPath
}

func getIndexToSwap(path []int) int {
	// find the index before the last strictly decreasing sequence
	for x := len(path) - 1; x > 0; x-- {
		if path[x] > path[x-1] {
			return x - 1
		}
	}
	return -1
}

func getNextBiggestAndSortRemainder(indexToSwap int, path []int) (int, []int) {
	nextBiggestIndex := -1
	nextBiggestDigit := 100
	digitToReplace := path[indexToSwap]
	pathCopy := make([]int, len(path))
	copy(pathCopy, path)
	for x, d := range pathCopy[indexToSwap:] {
		if d > digitToReplace && d < nextBiggestDigit {
			nextBiggestIndex = x + indexToSwap
			nextBiggestDigit = d
		}
	}
	if nextBiggestIndex == -1 {
		return -1, []int{}
	}

	// remove the next biggest digit and sort the remaining digits
	digitsAfter := append(pathCopy[indexToSwap:nextBiggestIndex], pathCopy[nextBiggestIndex+1:]...)
	slices.Sort(digitsAfter)
	return nextBiggestDigit, digitsAfter
}

func addLeadingZeros(path int, maxPathStr string) string {
	zerosToAdd := len(maxPathStr) - len(fmt.Sprintf("%d", path))
	return strings.Repeat("0", zerosToAdd) + fmt.Sprintf("%d", path)
}

func shouldSkipPath(pathStr string, unopenedTargets []string) bool {
	// if path contains a repeated index, skip it
	found := map[rune]bool{}
	for _, r := range pathStr {
		if found[r] {
			return true
		}
		found[r] = true
	}

	// if path contains an index out of range, skip it
	for _, r := range pathStr {
		index := utils.RuneToInt(r)
		if index >= len(unopenedTargets) {
			return true
		}
	}
	return false
}

func (s state) orderTargetValves(pathStr string, unopenedTargets []string, nameToValve map[string]valve) []valve {
	targetValvesInOrder := []valve{}
	for _, r := range pathStr {
		targetIndex := utils.RuneToInt(r)
		if targetIndex >= len(unopenedTargets) {
			continue
		}
		targetValvesInOrder = append(targetValvesInOrder, nameToValve[unopenedTargets[targetIndex]])
	}
	return targetValvesInOrder
}

func (s state) getCached(cache map[state]int) (int, bool) {
	// previously released pressure and targets count is irrelevant for checking cache
	newS := state{
		timeLeft:   s.timeLeft,
		position:   s.position,
		openValves: s.openValves,
	}
	if cached, isCached := cache[newS]; isCached {
		return cached, true
	}
	return 0, false
}

func (s state) isValveOpen(index int) bool {
	indexBin := 1 << (s.targetsCount - index - 1)
	result := s.openValves&(indexBin) == indexBin
	return result
}
