package day

import "fmt"

func part1(input []string) int {
	score := 0
	for _, round := range input {
		score += getRoundScore(round)
	}
	return score
}

func getRoundScore(round string) int {
	score := 0
	theirMove := round[:1]
	myMove := round[2:]

	// increase score by shape
	switch myMove {
	case "X":
		score += 1
	case "Y":
		score += 2
	case "Z":
		score += 3
	default:
		panic(fmt.Sprintf("%s isn't a valid move for me.", myMove))
	}

	if theirMove == "A" {
		switch myMove {
		case "X":
			score += 3
		case "Y":
			score += 6
		case "Z":
			score += 0
		}
		return score
	}

	if theirMove == "B" {
		switch myMove {
		case "X":
			score += 0
		case "Y":
			score += 3
		case "Z":
			score += 6
		}
		return score
	}

	if theirMove == "C" {
		switch myMove {
		case "X":
			score += 6
		case "Y":
			score += 0
		case "Z":
			score += 3
		}
		return score
	}

	panic(fmt.Sprintf("%s isn't a valid move for them.", theirMove))
}
