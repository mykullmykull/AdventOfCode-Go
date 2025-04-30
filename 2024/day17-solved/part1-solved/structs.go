package day

import (
	"fmt"
)

type computer struct {
	a            int
	b            int
	c            int
	p            int
	instructions []int
	outputs      string
}

func (c computer) run() computer {
	for c.p < len(c.instructions) {
		println(c.toString())
		opcode := c.instructions[c.p]
		operand := c.instructions[c.p+1]
		switch opcode {
		case 0:
			c = c.adv(operand)
		case 1:
			c = c.bxl(operand)
		case 2:
			c = c.bst(operand)
		case 3:
			before := c.p
			c = c.jnz(operand)
			if c.p != before {
				continue
			}
		case 4:
			c = c.bxc(operand)
		case 5:
			c = c.out(operand)
		case 6:
			c = c.bdv(operand)
		case 7:
			c = c.cdv(operand)
		default:
			panic(fmt.Sprintf("invalid opcode %d", opcode))
		}
		c.p += 2
	}
	return c
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

func (c computer) adv(operand int) computer {
	fmt.Printf("  adv: %d / power(2, %d) = %d\n", c.a, c.combo(operand), c.a/power(2, c.combo(operand)))
	c.a = c.a / power(2, c.combo(operand))
	return c
}

func (c computer) bxl(operand int) computer {
	fmt.Printf("  bxl: %d ^ %d = %d\n", c.b, operand, bitwiseXOR(c.b, operand))
	c.b = bitwiseXOR(c.b, operand)
	return c
}

func (c computer) bst(operand int) computer {
	fmt.Printf("  bst: %d %% 8 = %d\n", c.combo(operand), c.combo(operand)%8)
	c.b = c.combo(operand) % 8
	return c
}

func (c computer) jnz(operand int) computer {
	fmt.Printf("  jnz: a = %d, p = %d, operand = %d\n", c.a, c.p, operand)
	if c.a == 0 {
		return c
	}
	c.p = operand
	return c
}

func (c computer) bxc(operand int) computer {
	fmt.Printf("  bxc: %d ^ %d = %d\n", c.b, c.c, bitwiseXOR(c.b, c.c))
	c.b = bitwiseXOR(c.b, c.c)
	return c
}

func (c computer) out(operand int) computer {
	output := c.combo(operand) % 8
	fmt.Printf("  out: %d\n", output)
	if len(c.outputs) == 0 {
		c.outputs = fmt.Sprintf("%d", output)
		return c
	}
	c.outputs += fmt.Sprintf(",%d", output)
	return c
}

func (c computer) bdv(operand int) computer {
	fmt.Printf("  bdv: %d / power(2, %d) = %d\n", c.b, c.combo(operand), c.b/power(2, c.combo(operand)))
	c.b = c.a / power(2, c.combo(operand))
	return c
}

func (c computer) cdv(operand int) computer {
	fmt.Printf("  cdv: %d / power(2, %d) = %d\n", c.c, c.combo(operand), c.c/power(2, c.combo(operand)))
	c.c = c.a / power(2, c.combo(operand))
	return c
}

func (c computer) toString() string {
	return fmt.Sprintf("a=%d b=%d c=%d p=%d instructions=%v outputs=%s", c.a, c.b, c.c, c.p, c.instructions, c.outputs)
}
