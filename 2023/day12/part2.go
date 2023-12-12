package day12

func runB(input []string) int {
	sum := 0
	springsShort, groupsShort := parseInput(input)
	springs, groups := expand(springsShort, groupsShort)
	for i := 0; i < len(springs); i++ {
		sum += countSolutions(springs[i], groups[i])
	}
	return sum
}

func expand(springsShort []string, groupsShort [][]int) ([]string, [][]int) {
	springs := make([]string, len(springsShort))
	for i, s := range springsShort {
		newSpring := ""
		for j := 0; j < 5; j++ {
			if j > 0 {
				newSpring += "?"
			}
			newSpring += s
		}
		springs[i] = newSpring
	}

	groups := make([][]int, len(groupsShort))
	for i, g := range groupsShort {
		newGroup := []int{}
		for j := 0; j < 5; j++ {
			newGroup = append(newGroup, g...)
		}
		groups[i] = newGroup
	}
	return springs, groups
}
