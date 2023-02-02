package day5

import (
	"log"
	"strconv"
	"strings"
)

func TopCratesAfterMoves(crateRows []string, moves []string, part int) string {
	crates := parseCrates(crateRows)
	crates = moveCrates(crates, moves, part)
	tops := getTops(crates)
	return tops
}

func moveCrates(crates [][]string, moves []string, part int) [][]string {
	for _, move := range moves {
		crates = executeSingleMove(crates, move, part)
	}
	return crates
}

func executeSingleMove(crates [][]string, move string, part int) [][]string {
	fields := strings.Fields(move)
	count, e1 := strconv.Atoi(fields[1])
	stackAName, e2 := strconv.Atoi(fields[3])
	stackBName, e3 := strconv.Atoi(fields[5])

	if e1 != nil || e2 != nil || e3 != nil {
		log.Fatal("unable to convert something to int")
	}

	var moves int
	if part == 1 {
		moves = count
		count = 1
	} else {
		moves = 1
	}

	for i := 0; i < moves; i++ {
		var stackACopy []string
		stackACopy = append(stackACopy, crates[stackAName-1]...)
		crate := stackACopy[0:count]
		crates[stackBName-1] = append(crate, crates[stackBName-1]...)
		crates[stackAName-1] = crates[stackAName-1][count:]
	}

	return crates
}

func getTops(crates [][]string) string {
	var tops string
	for _, crate := range crates {
		tops = tops + crate[0]
	}
	return tops
}

func parseCrates(crateRows []string) [][]string {
	var crates [][]string
	maxIndex := len(crateRows[0])
	for i := 0; i < maxIndex; i = i + 4 {
		var stack []string
		for _, row := range crateRows {
			if row == crateRows[len(crateRows)-1] {
				break
			}
			crateStr := row[i : i+3]
			crate := string(crateStr[1])
			if crate != " " {
				stack = append(stack, crate)
			}
		}
		crates = append(crates, stack)
	}
	return crates
}
