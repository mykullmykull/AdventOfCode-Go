package day

import (
	"goAoc2021/utils"
	"strings"
)

func main(input []string) int {
	x, y, aim := 0, 0, 0
	for _, line := range input {
		parts := strings.Split(line, " ")
		direction := parts[0]
		distance := utils.Atoi(parts[1])
		switch direction {
		case "forward":
			x += distance
			y += aim * distance
		case "down":
			aim += distance
		case "up":
			aim -= distance
		default:
			panic("Unknown direction: " + direction)
		}
	}
	return x * y
}
