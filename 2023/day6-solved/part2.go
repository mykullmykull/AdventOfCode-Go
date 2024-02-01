package day6

import (
	"strconv"
)

func runB(input []string) int {
	raceDuration := parseRaceData(input[0])
	recordDistance := parseRaceData(input[1])
	return countWaysToWin(raceDuration, recordDistance)
}

func parseRaceData(str string) int {
	num := 0
	for i := 0; i < len(str); i++ {
		n, err := strconv.Atoi(str[i : i+1])
		if err == nil {
			num = num*10 + n
		}
	}
	return num
}

func countWaysToWin(raceDuration int, recordDistance int) int {
	waysToWin := 0
	for buttonTime := 0; buttonTime < raceDuration; buttonTime++ {
		speed := buttonTime
		timeLeft := raceDuration - buttonTime
		distance := speed * timeLeft
		if distance > recordDistance {
			waysToWin++
		}
	}
	return waysToWin
}
