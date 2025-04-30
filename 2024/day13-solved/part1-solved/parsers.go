package day

import (
	"goAoc2024/utils"
	"strings"
)

func parse(input []string) []machine {
	machines := []machine{}
	for index := 0; index < len(input); index += 4 {
		lineA := input[index]
		lineB := input[index+1]
		lineP := input[index+2]
		machines = append(machines, machine{
			a:     parseButton(lineA, 3),
			b:     parseButton(lineB, 1),
			prize: parsePoint(lineP),
		})
	}
	return machines
}

func parseButton(line string, tokens int) button {
	noX := strings.ReplaceAll(line, "X+", "")
	split := strings.Split(strings.Split(noX, ": ")[1], ", Y+")
	xChange := utils.Atoi(split[0])
	yChange := utils.Atoi(split[1])
	return button{xChange, yChange, tokens}
}

func parsePoint(line string) point {
	noX := strings.ReplaceAll(line, "X=", "")
	split := strings.Split(strings.Split(noX, ": ")[1], ", Y=")
	x := utils.Atoi(split[0])
	y := utils.Atoi(split[1])
	return point{x, y}
}
