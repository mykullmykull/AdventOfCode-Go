package day

func main(input []string) screen {
	c := cpu{
		cycle:            0,
		x:                1,
		instructionIndex: 0,
	}
	crt := screen{}.createCrt()
	for c.instructionIndex < len(input) {
		c, crt = c.advance(input, crt)
	}
	return crt
}
