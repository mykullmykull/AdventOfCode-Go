package day

import (
	"goAoc2022/utils"
	"math"
	"strings"
)

type cpu struct {
	x                int
	cycle            int
	instructionIndex int
}

type screen struct {
	pixels []string
}

func (c cpu) advance(instructions []string, crt screen) (cpu, screen) {
	crt = crt.updatePixel(c)
	instruction := strings.Split(instructions[c.instructionIndex], " ")
	c.cycle++
	if instruction[0] == "noop" {
		c.instructionIndex++
		return c, crt
	}
	crt = crt.updatePixel(c)
	c.cycle++
	c.x += utils.StrToInt(instruction[1])
	c.instructionIndex++
	return c, crt
}

func (crt screen) createCrt() screen {
	emptyRow := crt.getEmptyRow()
	return screen{
		pixels: []string{emptyRow, emptyRow, emptyRow, emptyRow, emptyRow, emptyRow},
	}
}

func (crt screen) getEmptyRow() string {
	emptyRow := ""
	for cycle := 1; cycle <= 40; cycle++ {
		emptyRow += "."
	}
	return emptyRow
}

func (crt screen) updatePixel(c cpu) screen {
	crtRow := c.cycle / 40
	crtCol := c.cycle % 40
	if math.Abs(float64(crtCol-c.x)) < 2 {
		crt = crt.markPixel(crtRow, crtCol, "#")
	}
	return crt
}

func (crt screen) markPixel(row int, col int, newChar string) screen {
	if row >= 0 && row < len(crt.pixels) &&
		col >= 0 && col < len(crt.pixels[row]) {
		crt.pixels[row] = crt.pixels[row][:col] + newChar + crt.pixels[row][col+1:]
	}
	return crt
}

func (crt screen) printCrt(cycle int) {
	crt.pixels = []string{crt.getEmptyRow()}
	crt.markPixel(0, cycle%40, "0")
	crt.printScreen()
}

func (crt screen) printSprite(x int) {
	crt.pixels = []string{crt.getEmptyRow()}
	crt.markPixel(0, x-1, "*")
	crt.markPixel(0, x+0, "*")
	crt.markPixel(0, x+1, "*")
	crt.printScreen()
}

func (crt screen) printScreen() {
	for _, row := range crt.pixels {
		println(row)
	}
}
