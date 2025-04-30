package day

import (
	"goAoc2022/utils"
	"strings"
)

func parseFactories(input []string) []factory {
	factories := []factory{}
	for _, line := range input {
		factory := parseFactory(line)
		factories = append(factories, factory)
	}
	return factories
}

func parseFactory(line string) factory {
	f := factory{costs: map[int][4]int{}}
	ints := cleanAndSplitLineToInts(line)
	f.id = ints[0]
	oreCost := [4]int{ints[1], 0, 0, 0}
	clyCost := [4]int{ints[2], 0, 0, 0}
	obsCost := [4]int{ints[3], ints[4], 0, 0}
	geoCost := [4]int{ints[5], 0, ints[6], 0}
	f.costs[ore] = oreCost
	f.costs[cly] = clyCost
	f.costs[obs] = obsCost
	f.costs[geo] = geoCost
	return f
}

func cleanAndSplitLineToInts(line string) []int {
	clean := strings.Replace(
		strings.Replace(
			strings.Replace(
				strings.Replace(
					strings.Replace(
						strings.Replace(
							strings.Replace(
								strings.Replace(
									strings.Replace(
										strings.Replace(
											strings.Replace(
												strings.Replace(
													strings.Replace(
														line,
														"Blueprint ", "", -1),
													":", "", -1),
												"Each ", "", -1),
											"robot costs ", "", -1),
										"ore ", "", -1),
									"clay ", "", -1),
								"obsidian ", "", -1),
							"geode ", "", -1),
						"ore. ", "", -1),
					"clay. ", "", -1),
				"obsidian.", "", -1),
			"geode. ", "", -1),
		"and ", "", -1)

	split := strings.Split(clean, " ")
	if len(split) < 7 {
		panic("there should be 7 ints left in the line: " + utils.JoinArray(split, " "))
	}

	ints := []int{}
	for x := 0; x < 7; x++ {
		ints = append(ints, utils.StrToInt(split[x]))
	}

	return ints
}
