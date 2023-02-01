package day2

func FinalScoreB(rounds []string) int {
	score := 0
	for _, v := range rounds {
		points := 0
		switch v {
		case "A X":
			points = 3 + 0
		case "A Y":
			points = 1 + 3
		case "A Z":
			points = 2 + 6
		case "B X":
			points = 1 + 0
		case "B Y":
			points = 2 + 3
		case "B Z":
			points = 3 + 6
		case "C X":
			points = 2 + 0
		case "C Y":
			points = 3 + 3
		case "C Z":
			points = 1 + 6
		}
		score = score + points
	}
	return score
}

func FinalScore(rounds []string) int {
	score := 0
	for _, v := range rounds {
		points := 0
		switch v {
		case "A X":
			points = 1 + 3
		case "A Y":
			points = 2 + 6
		case "A Z":
			points = 3 + 0
		case "B X":
			points = 1 + 0
		case "B Y":
			points = 2 + 3
		case "B Z":
			points = 3 + 6
		case "C X":
			points = 1 + 6
		case "C Y":
			points = 2 + 0
		case "C Z":
			points = 3 + 3
		}
		score = score + points
	}
	return score
}
