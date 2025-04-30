package day

import (
	"fmt"
	"goAoc2022/utils"
	"strconv"
)

func main(input []string) int {
	nameToValve, _, targets := parseTargets(input)
	printMap(nameToValve)
	// test(nameToValve, targets)
	// return 0

	cache := map[state]int{}

	maxOpenBin := ""
	for x := 0; x < len(targets); x++ {
		maxOpenBin += "1"
	}
	max64, _ := strconv.ParseInt(maxOpenBin, 2, 0)
	maxOpenInt := int(max64)
	// halfMax := maxOpenInt / 2
	maxReleased := 0

	// loop through all possible combinations of open valves that the other actor may have opened (eOpened)
	// only need half of max since if e opens eOpened then I will open the complement of eOpened and vice versa
	for eOpened := maxOpenInt; eOpened >= 0; eOpened-- {
		fmt.Printf("considering %b\n", eOpened)
		iOpened := maxOpenInt ^ eOpened
		myState := state{
			timeLeft:     26,
			position:     "AA",
			openValves:   eOpened,
			targetsCount: len(targets),
		}
		eState := state{
			timeLeft:     26,
			position:     "AA",
			openValves:   iOpened,
			targetsCount: len(targets),
		}
		myRelease, cacheAfterMe := myState.releaseMaxPressure(targets, nameToValve, cache)
		cache = cacheAfterMe
		eRelease, cacheAfterE := eState.releaseMaxPressure(targets, nameToValve, cache)

		cache = cacheAfterE
		maxReleased = utils.UpdateMost(maxReleased, myRelease+eRelease)
	}
	return maxReleased
}

func test(nameToValve map[string]valve, targets []string) {
	eOpened := 0b110001
	iOpened := 0b001110
	myState := state{
		timeLeft:     26,
		position:     "AA",
		openValves:   eOpened,
		targetsCount: len(targets),
	}
	eState := state{
		timeLeft:     26,
		position:     "AA",
		openValves:   iOpened,
		targetsCount: len(targets),
	}
	cache := map[state]int{}
	myRelease, cacheAfterMe := myState.releaseMaxPressure(targets, nameToValve, cache)
	cache = cacheAfterMe
	eRelease, _ := eState.releaseMaxPressure(targets, nameToValve, cache)
	println("  released", myRelease+eRelease)
}
