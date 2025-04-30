package day

import (
	"goAoc2024/utils"
)

func main(input []string) int {
	comp, program := parse(input)
	nextAs := []int{0}
	for x := len(program) - 1; x >= 0; x-- {
		if program[x] == ',' {
			continue
		}
		targetOut := utils.Atoi(program[x : x+1])
		newNextAs := []int{}
		for _, nextA := range nextAs {
			for a := 8 * nextA; a < 8*(nextA+1); a++ {
				comp = comp.reset(a)
				comp = comp.goToJmp()
				if comp.output == targetOut {
					newNextAs = append(newNextAs, a)
				}
			}
		}
		nextAs = newNextAs
	}
	return nextAs[0]
}

func (comp computer) goToJmp() computer {
	for comp.pointer < len(comp.instructions) {
		code := comp.instructions[comp.pointer]
		operand := comp.instructions[comp.pointer+1]
		comp = comp.followInstruction(code, operand)
		if code == 3 {
			return comp
		}
	}
	panic("didn't find a jmp")
}
