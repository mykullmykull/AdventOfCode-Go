package day1

import (
	"strconv"
)

func runB(input []string) int {
	sum := 0

	// index / 2 + 1 = the int translation
	validDigits := []string{
		"1",
		"one",
		"2",
		"two",
		"3",
		"three",
		"4",
		"four",
		"5",
		"five",
		"6",
		"six",
		"7",
		"seven",
		"8",
		"eight",
		"9",
		"nine",
	}

	for _, line := range input {
		firstDigit := -1
		lastDigit := -1
		for i := 0; i < len(line); i++ {
			for j, key := range validDigits {
				value := j/2 + 1
				if firstDigit < 0 {
					if len(key)+i > len(line) {
						continue
					}

					if line[i:i+len(key)] == key {
						firstDigit = value
					}
				}

				if lastDigit < 0 {
					if len(key)+i > len(line) {
						continue
					}

					if line[len(line)-len(key)-i:len(line)-i] == key {
						lastDigit = value
					}
				}

				if firstDigit > 0 && lastDigit > 0 {
					break
				}
			}
			if firstDigit > 0 && lastDigit > 0 {
				sum += firstDigit*10 + lastDigit
				break
			}
		}
	}
	return sum
}

func runA(input []string) int {
	sum := 0
	for _, line := range input {
		firstDigit := -1
		lastDigit := -1
		for i := 0; i < len(line); i++ {
			if firstDigit < 0 {
				firstDigit = getFirstDigitA(line, i)
			}

			if lastDigit < 0 {
				lastDigit = getLastDigitA(line, i)
			}

			if firstDigit >= 0 && lastDigit >= 0 {
				sum += firstDigit*10 + lastDigit
				break
			}
		}
	}
	return sum
}

func getFirstDigitA(line string, i int) int {
	n, err := strconv.Atoi(line[i : i+1])
	if err != nil {
		return -1
	}
	return n
}

func getLastDigitA(line string, i int) int {
	n, err := strconv.Atoi(line[len(line)-i-1 : len(line)-i])
	if err != nil {
		return -1
	}
	return n
}
