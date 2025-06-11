package day

import (
	"goAoc2021/utils"
	"strings"
)

func main(input []string) int {
	foldedDots := map[[2]int]bool{}
	fold := getFirstFold(input)
	for _, line := range input {
		if line == "" {
			break
		}
		split := strings.Split(line, ",")
		x := utils.Atoi(split[0])
		y := utils.Atoi(split[1])
		foldedDots[foldDot(x, y, fold)] = true
	}

	return len(foldedDots)
}

func getFirstFold(input []string) string {
	foundFolds := false
	for _, line := range input {
		if foundFolds {
			return strings.ReplaceAll(line, "fold along ", "")
		}
		if line == "" {
			foundFolds = true
			continue
		}
	}
	panic("No folds found in input")
}

func foldDot(x, y int, fold string) [2]int {
	split := strings.Split(fold, "=")
	dimension := split[0]
	dist := utils.Atoi(split[1])

	switch dimension {
	case "x":
		if x > dist {
			x = dist - (x - dist)
		}
	case "y":
		if y > dist {
			y = dist - (y - dist)
		}
	default:
		panic("Invalid fold dimension")
	}

	return [2]int{x, y}
}
