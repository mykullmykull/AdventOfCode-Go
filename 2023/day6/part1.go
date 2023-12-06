package day6

import (
	"strconv"
)

func runA(input []string) int {
	raceDurations := parseListOfSpacedInts(input[0])
	recordDistances := parseListOfSpacedInts(input[1])
	waysToWinEachRace := countWaysToWinEachRace(raceDurations, recordDistances)
	return multiplyWaysToWin(waysToWinEachRace)
}

func parseListOfSpacedInts(str string) []int {
	numbers := []int{}
	num := 0
	for i := 0; i < len(str); i++ {
		char := str[i : i+1]
		if char == " " {
			if num > 0 {
				numbers = append(numbers, num)
			}
			num = 0
		}

		n, err := strconv.Atoi(char)
		if err == nil {
			num = num*10 + n
		}
	}
	numbers = append(numbers, num)
	return numbers
}

func countWaysToWinEachRace(raceDurations []int, recordDistances []int) []int {
	waysToWinEachRace := make([]int, len(raceDurations))
	for i := 0; i < len(raceDurations); i++ {
		duration := raceDurations[i]
		waysToWin := 0
		for buttonTime := 0; buttonTime < duration; buttonTime++ {
			speed := buttonTime
			timeLeft := duration - buttonTime
			distance := speed * timeLeft
			if distance > recordDistances[i] {
				waysToWin++
			}
		}
		waysToWinEachRace[i] = waysToWin
	}
	return waysToWinEachRace
}

func multiplyWaysToWin(waysToWinEachRace []int) int {
	product := 1
	for i := 0; i < len(waysToWinEachRace); i++ {
		product *= waysToWinEachRace[i]
	}
	return product
}
