package day4

import (
	"fmt"
	"strconv"
	"strings"
)

var sortedLetters []rune = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

func runB(input []string) {
	for _, row := range input {
		name, sectorId := sectorIdIfRealRoom(row)
		if sectorId == 0 {
			continue
		}

		fmt.Println(decryptName(name, sectorId) + ", sectorId: " + strconv.Itoa(sectorId))
	}
	return
}

func decryptName(name string, sectorId int) string {
	decrypted := ""
	for _, letter := range name {
		if letter == '-' {
			decrypted += " "
			continue
		}

		decrypted += string(shift(letter, sectorId))
	}
	return decrypted
}

func shift(letter rune, sectorId int) rune {
	index := strings.IndexRune(string(sortedLetters), letter)
	if index == -1 {
		panic(fmt.Sprintf("letter %s not found in sortedLetters", string(letter)))
	}

	newIndex := (index + sectorId) % len(sortedLetters)
	return sortedLetters[newIndex]
}

func runA(input []string) int {
	sum := 0
	for _, row := range input {
		_, sectorId := sectorIdIfRealRoom(row)
		sum += sectorId
	}
	return sum
}

func sectorIdIfRealRoom(row string) (string, int) {
	name, sectorID, checksum := parseRow(row)
	calculatedChecksum := calculateChecksum(name)
	if checksumsMatch(checksum, calculatedChecksum) {
		return name, sectorID
	}

	return "", 0
}

func parseRow(row string) (string, int, string) {
	splitChecksum := strings.Split(strings.Split(row, "]")[0], "[")
	checksum := splitChecksum[1]
	splitSectorID := strings.Split(splitChecksum[0], "-")
	sectorID, err := strconv.Atoi(splitSectorID[len(splitSectorID)-1])
	if err != nil {
		panic(fmt.Sprintf("cannot convert %s to int", splitSectorID[len(splitSectorID)-1]))
	}

	name := strings.Join(splitSectorID[:len(splitSectorID)-1], "-")
	return name, sectorID, checksum
}

func calculateChecksum(name string) string {
	countLetters := make(map[rune]int)
	for _, letter := range name {
		if letter == '-' {
			continue
		}

		countLetters[letter]++
	}
	fiveMostCommon := findFiveMostCommon(countLetters)
	return fiveMostCommon
}

func findFiveMostCommon(countLetters map[rune]int) string {
	fiveMostCommon := ""
	for i := 0; i < 5; i++ {
		max := 0
		var maxLetter rune
		for _, letter := range sortedLetters {
			count := countLetters[letter]
			if count > max {
				max = count
				maxLetter = letter
			}
		}
		fiveMostCommon += string(maxLetter)
		delete(countLetters, maxLetter)
	}
	return fiveMostCommon
}

func checksumsMatch(checksum, calculatedChecksum string) bool {
	matches := true
	for i := 0; i < 5; i++ {
		letterMatches := false
		for j := 0; j < 5; j++ {
			if checksum[i] == calculatedChecksum[j] {
				letterMatches = true
				break
			}
		}
		if !letterMatches {
			matches = false
			break
		}
	}
	return matches
}
