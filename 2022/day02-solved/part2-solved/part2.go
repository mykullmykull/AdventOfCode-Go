package day

import "fmt"

func part2(input []string) int {
	score := 0
	for _, round := range input {
		score += getRoundScore(round)
	}
	return score
}

func getRoundScore(round string) int {
	score := 0
	theirMove := round[:1]
	outcome := round[2:]

	// add score from outcome
	switch outcome {
	case "X":
		score += 0
	case "Y":
		score += 3
	case "Z":
		score += 6
	default:
		panic(fmt.Sprintf("%s isn't a valid outcome.", outcome))
	}

	// add score from my shape
	return score + scoreMyShape(outcome, theirMove)
}

func scoreMyShape(outcome string, theirMove string) int {
	invalidError := fmt.Sprintf("%s isn't a valid outcome.", outcome)
	if outcome == "X" {
		switch theirMove {
		case "A":
			return 3
		case "B":
			return 1
		case "C":
			return 2
		default:
			panic(invalidError)
		}
	}

	if outcome == "Y" {
		switch theirMove {
		case "A":
			return 1
		case "B":
			return 2
		case "C":
			return 3
		default:
			panic(invalidError)
		}
	}

	if outcome == "Z" {
		switch theirMove {
		case "A":
			return 2
		case "B":
			return 3
		case "C":
			return 1
		default:
			panic(invalidError)
		}
	}

	panic(fmt.Sprintf("%s is an invalid outcome.", outcome))
}
