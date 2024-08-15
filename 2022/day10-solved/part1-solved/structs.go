package day

import (
	"goAoc2022/utils"
	"strings"
)

type cpu struct {
	x                int
	cycle            int
	instructionIndex int
}

func (c cpu) advanceToTime(time int, instructions []string) cpu {
	for c.cycle < time-2 {
		// if time == 220 {
		// 	println("cycle", c.cycle, "x", c.x, "instruction", instructions[c.instructionIndex])
		// }
		split := strings.Split(instructions[c.instructionIndex], " ")
		switch split[0] {
		case "noop":
			c.cycle++
		case "addx":
			if time == c.cycle+1 {
				return c
			}

			c.cycle += 2
			c.x += utils.StrToInt(split[1])
		}
		c.instructionIndex++
	}
	// println("cycle", c.cycle, "x", c.x, "instruction", instructions[c.instructionIndex])
	return c
}
