package day15

import (
	"fmt"
	"goAoc2023/utils"
	"strings"
)

type lens struct {
	label       string
	focalLength int
}

func runB(input string) int {
	boxes := map[int][]lens{}
	steps := parse(input)
	for _, step := range steps {
		boxes = swapLenses(boxes, step)
	}
	// printBoxes(boxes)
	return totalFocusingPower(boxes)
}

func parse(input string) []lens {
	noDashes := strings.Replace(input, "-", "", -1)
	stepStrings := strings.Split(noDashes, ",")
	steps := make([]lens, len(stepStrings))
	for i, str := range stepStrings {
		split := strings.Split(str, "=")
		label := split[0]
		focalLength := -1
		if len(split) == 2 {
			focalLength = utils.Atoi(split[1])
		}
		steps[i] = lens{label, focalLength}
	}
	return steps
}

func swapLenses(boxes map[int][]lens, step lens) map[int][]lens {
	boxId := hashString(step.label)
	lenses := boxes[boxId]
	newLenses := []lens{}
	if step.focalLength > 0 {
		newLenses = replaceOrAddLens(lenses, step)
		boxes[boxId] = newLenses
		return boxes
	}

	for _, lens := range lenses {
		if lens.label == step.label {
			continue
		}
		newLenses = append(newLenses, lens)
	}
	boxes[boxId] = newLenses
	return boxes
}

func replaceOrAddLens(lenses []lens, step lens) []lens {
	index := -1
	for i, lens := range lenses {
		if lens.label == step.label {
			index = i
			break
		}
	}
	if index > -1 {
		lenses[index] = step
	} else {
		lenses = append(lenses, step)
	}
	return lenses
}

func totalFocusingPower(boxes map[int][]lens) int {
	sum := 0
	for id, lenses := range boxes {
		for num, lens := range lenses {
			sum += (id + 1) * (num + 1) * lens.focalLength
			// fmt.Printf("box id: %d, lens num: %d, label: %s, focalLength: %d, sum: %d\n", id+1, num+1, lens.label, lens.focalLength, sum)
		}
	}
	return sum
}

func printBoxes(boxes map[int][]lens) {
	for id, lenses := range boxes {
		fmt.Printf("Box %d: %v\n", id, lenses)
	}
}
