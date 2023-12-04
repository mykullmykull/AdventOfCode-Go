package day4

import (
	"math"
	"strings"
)

func runA(input []string) int {
	sum := 0
	for _, line := range input {
		sum += int(math.Pow(2, float64(countWins(line)-1)))
	}
	return sum
}

func countWins(line string) int {
	count := 0
	withoutCardId := strings.Split(line, ": ")[1]
	split := strings.Split(withoutCardId, " | ")
	winningNumbers := strings.Split(split[0], " ")
	haveNumbers := strings.Split(split[1], " ")
	for _, have := range haveNumbers {
		for _, winning := range winningNumbers {
			if have == winning && strings.Trim(have, " ") != "" {
				count++
				break
			}
		}
	}
	return count
}

func runB(input []string) int {
	winsPerCardId := make([]int, len(input))
	for i, line := range input {
		winsPerCardId[i] = countWins(line)
	}

	return copyAndCountCards(winsPerCardId)
}

func copyAndCountCards(winsPerCardId []int) int {
	startingCardsPerId := oneOfEach(len(winsPerCardId))
	cardsPerId := copyCardsBasedOnWins(startingCardsPerId, winsPerCardId)

	count := 0
	for _, copies := range cardsPerId {
		count += copies
	}
	return count
}

func oneOfEach(n int) []int {
	cardsOfEachId := make([]int, n)
	// each id starts with 1 card
	for i, _ := range cardsOfEachId {
		cardsOfEachId[i] = 1
	}
	return cardsOfEachId
}

func copyCardsBasedOnWins(cardsPerId []int, winsPerCardId []int) []int {
	for sourceId, wins := range winsPerCardId {
		for idCopied := sourceId + 1; idCopied < sourceId+wins+1; idCopied++ {
			cardsPerId[idCopied] += cardsPerId[sourceId]
		}
	}
	return cardsPerId
}
