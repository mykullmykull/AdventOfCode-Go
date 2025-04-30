package day

import (
	"goAoc2024/utils"
	"strings"
)

func main(input []string) string {
	c := parse(input)
	c = c.run()
	return c.outputs
}

func parse(input []string) computer {
	a := utils.Atoi(strings.Split(input[0], ": ")[1])
	b := utils.Atoi(strings.Split(input[1], ": ")[1])
	c := utils.Atoi(strings.Split(input[2], ": ")[1])
	split := strings.Split(strings.Split(input[4], ": ")[1], ",")
	var instructions []int
	for _, s := range split {
		instructions = append(instructions, utils.Atoi(s))
	}
	return computer{a, b, c, 0, instructions, ""}
}
