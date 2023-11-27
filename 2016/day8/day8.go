package day8

import (
	"fmt"
	"strconv"
	"strings"
)

func runA(input []string) int {
	screen := createScreen()
	screen = followInstructions(screen, input)
	printScreen(screen)
	return countLitPixels(screen)
}

func followInstructions(screen []string, instructions []string) []string {
	for _, instruction := range instructions {
		screen = followInstruction(screen, instruction)
	}
	return screen
}

func followInstruction(screen []string, instruction string) []string {
	split := strings.Split(instruction, " ")
	if split[0] == "rect" {
		return rect(screen, split[1])
	}

	if split[0] == "rotate" {
		index := convertAtoi(strings.Split(split[2], "=")[1])
		amount := convertAtoi(split[4])
		return rotate(screen, split[1], index, amount)
	}

	panic(fmt.Sprintf("Unknown instruction: %s", instruction))
}

func rect(screen []string, size string) []string {
	split := strings.Split(size, "x")
	width, err := strconv.Atoi(split[0])
	if err != nil {
		panic(fmt.Sprintf("Could not parse width: %s", split[0]))
	}

	height, err := strconv.Atoi(split[1])
	if err != nil {
		panic(fmt.Sprintf("Could not parse height: %s", split[1]))
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			screen[i] = replaceAtIndex(screen[i], j, '#')
		}
	}
	return screen
}

func rotate(screen []string, rowCol string, index int, amount int) []string {
	new := createScreen()
	height := len(screen)
	width := len(screen[0])
	for y := 0; y < len(new); y++ {
		for x := 0; x < width; x++ {
			if rowCol == "row" && y == index {
				indexToMove := (x - amount + width) % width
				new[y] = replaceAtIndex(new[y], x, rune(screen[y][indexToMove]))
			} else if rowCol == "column" && x == index {
				indexToMove := (y - amount + height) % height
				new[y] = replaceAtIndex(new[y], x, rune(screen[indexToMove][x]))
			} else {
				new[y] = replaceAtIndex(new[y], x, rune(screen[y][x]))
			}
		}
	}
	return new
}

func createScreen() []string {
	screen := make([]string, 6)
	blankRow := ""
	for i := 0; i < 50; i++ {
		blankRow += "."
	}
	for i := range screen {
		screen[i] = blankRow
	}
	return screen
}

func countLitPixels(screen []string) int {
	count := 0
	for _, row := range screen {
		for _, pixel := range row {
			if pixel == '#' {
				count++
			}
		}
	}
	return count
}

func convertAtoi(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Sprintf("Could not parse %s to int", str))
	}
	return i
}

func replaceAtIndex(str string, index int, replacement rune) string {
	return str[:index] + string(replacement) + str[index+1:]
}

func printScreen(screen []string) {
	for _, row := range screen {
		fmt.Println(row)
	}
	fmt.Println()
}
