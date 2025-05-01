package day

import (
	"goAoc2021/utils"
	"strings"
)

func main(input []string) int {
	x := 0
	y := 0
	for _, line := range input {
		parts := strings.Split(line, " ")
		direction := parts[0]
		distance := utils.Atoi(parts[1])
		switch direction {
		case "forward":
			x += distance
		case "down":
			y += distance
		case "up":
			y -= distance
		default:
			panic("Unknown direction: " + direction)
		}
	}
	return x * y
}
