package day6

import "fmt"

func runB(input []string) string {
	signal := ""
	for i := 0; i < len(input[0]); i++ {
		countLetters := make(map[rune]int)
		for _, row := range input {
			letter := rune(row[i])
			countLetters[letter]++
		}
		least := 1000
		leastLetter := ' '
		for letter, count := range countLetters {
			if count < least {
				least = count
				leastLetter = letter
			}
		}
		signal += string(leastLetter)
	}
	return signal
}

func runA(input []string) string {
	signal := ""
	for i := 0; i < len(input[0]); i++ {
		countLetters := make(map[rune]int)
		max := 0
		maxLetter := ' '
		for _, row := range input {
			letter := rune(row[i])
			countLetters[letter]++
			if countLetters[letter] > max {
				max = countLetters[letter]
				maxLetter = letter
			}
			fmt.Printf("i: %d, letter: %c, count: %d, max: %d, maxLetter: %c\n", i, letter, countLetters[letter], max, maxLetter)
		}
		signal += string(maxLetter)
	}
	return signal
}
