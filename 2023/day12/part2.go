package day12

import "fmt"

func runB(input []string) int {
	if false {
		return runTest()
	}

	solutions := 0
	springsShort, groupsShort := parseInput(input)
	springs, groups := expand(springsShort, groupsShort)
	for i := 0; i < len(springs); i++ {
		log(fmt.Sprintf("i: %d out of %d", i, len(springs)), true)
		solutions += countSolutionsByNodes(springs[i], groups[i])
	}
	return solutions
}

func runTest() int {
	debug := true
	debug2 := true
	node := node{"YNYNYNYNYNYNNYY"}
	springs := "????.#...#...?????.#...#...?????.#...#...?????.#...#...?????.#...#..."
	groups := []int{4, 1, 1, 4, 1, 1, 4, 1, 1, 4, 1, 1, 4, 1, 1}
	printNode(node, springs, groups, debug, debug2)
	countSolutionsByNodes(springs, groups)
	return 0
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

func countSolutionsByNodes(springs string, groups []int) int {
	unknowns := countUnknowns(springs)
	solutions := map[node]int{}
	alreadyCounted := map[remainder]int{}
	node := node{""}
	for {
		remainder := node.getRemainder(springs, groups)
		if count, ok := alreadyCounted[remainder]; ok {
			solutions[node] = count
			node = node.parent()
			continue
		}

		if len(node.guessed) == unknowns {
			if node.fails(springs, groups, unknowns) {
				solutions[node] = 0
			} else {
				solutions[node] = 1
			}

			node = node.parent()
			continue
		}

		yesCount, yesDone := solutions[node.yesChild()]
		noCount, noDone := solutions[node.noChild()]

		if yesDone && noDone {
			count := yesCount + noCount
			if remainder.groupsLeft > 0 {
				alreadyCounted[remainder] = count
			}
			if node.guessed == "" {
				return count
			}
			delete(solutions, node.yesChild())
			delete(solutions, node.noChild())
			solutions[node] = count
			node = node.parent()
			continue
		}

		if !yesDone && node.fails(springs, groups, unknowns) {
			solutions[node] = 0
			node = node.parent()
			continue
		}

		if !yesDone {
			node = node.yesChild()
			continue
		}

		node = node.noChild()
		continue
	}
}

func countUnknowns(springs string) int {
	unknowns := 0
	for _, c := range springs {
		if c == '?' {
			unknowns++
		}
	}
	return unknowns
}

func printNode(n node, springs string, groups []int, debug bool, debug2 bool) {
	guessed := ""
	guessedIndex := 0
	for _, rune := range springs {
		if rune == '?' && guessedIndex < len(n.guessed) {
			guessed += string(n.guessed[guessedIndex])
			guessedIndex++
			continue
		}
		guessed += string(rune)
	}
	msg := fmt.Sprintf("  groups: %v, node: %s", groups, n.guessed)
	if debug2 {
		msg += fmt.Sprintf(", springs: %s", guessed)
	}
	log(msg, debug)
}

func log(msg string, shouldLog bool) {
	if !shouldLog {
		return
	}

	fmt.Println(msg)
}
