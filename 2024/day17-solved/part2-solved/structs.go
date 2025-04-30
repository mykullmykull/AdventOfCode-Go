package day

import (
	"fmt"
	"goAoc2024/utils"
	"strings"
)

type computer struct {
	a            int
	b            int
	c            int
	pointer      int
	output       int
	instructions []int
}

func (c computer) followInstruction(code int, operand int) computer {
	switch code {
	case 0:
		c = c.adv(operand)
	case 1:
		c = c.bxl(operand)
	case 2:
		c = c.bst(operand)
	case 3:
		c = c.jnz(operand)
	case 4:
		c = c.bxc()
	case 5:
		c = c.out(operand)
	case 6:
		c = c.bdv(operand)
	case 7:
		c = c.cdv(operand)
	default:
		panic(fmt.Sprintf("invalid code %d", code))
	}
	c.pointer += 2
	return c
}

func (c computer) adv(operand int) computer {
	exp := c.combo(operand)
	c.a = c.a / power(2, exp)
	return c
}

func (c computer) bxl(operand int) computer {
	c.b = bitwiseXOR(c.b, operand)
	return c
}

func (c computer) bst(operand int) computer {
	x := c.combo(operand)
	c.b = x % 8
	return c
}

func (c computer) jnz(operand int) computer {
	if c.a == 0 {
		return c
	}
	c.pointer = operand - 2
	return c
}

func (c computer) bxc() computer {
	c.b = bitwiseXOR(c.b, c.c)
	return c
}

func (c computer) out(operand int) computer {
	c.output = c.combo(operand) % 8
	return c
}

func (c computer) bdv(operand int) computer {
	c.b = c.a / power(2, c.combo(operand))
	return c
}

func (c computer) cdv(operand int) computer {
	c.c = c.a / power(2, c.combo(operand))
	return c
}

func (c computer) toString() string {
	return fmt.Sprintf("a=%d b=%d c=%d p=%d outputs=%d", c.a, c.b, c.c, c.pointer, c.output)
}

func (c computer) combo(operand int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return c.a
	case 5:
		return c.b
	case 6:
		return c.c
	default:
		panic(fmt.Sprintf("invalid operand %d", operand))
	}
}

func parse(input []string) (computer, string) {
	a := utils.Atoi(strings.Split(input[0], ": ")[1])
	b := utils.Atoi(strings.Split(input[1], ": ")[1])
	c := utils.Atoi(strings.Split(input[2], ": ")[1])
	program := strings.Split(input[4], ": ")[1]
	strs := strings.Split(program, ",")
	instructions := make([]int, len(strs))
	for x, str := range strs {
		instructions[x] = utils.Atoi(str)
	}
	return computer{a, b, c, 0, -1, instructions}, program
}

func (c computer) reset(a int) computer {
	return computer{
		a:            a,
		b:            0,
		c:            0,
		pointer:      0,
		output:       -1,
		instructions: c.instructions,
	}
}
