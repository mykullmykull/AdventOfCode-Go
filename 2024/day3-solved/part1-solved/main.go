package day

import (
	"goAoc2024/utils"
	"strconv"
)

func main(input string) int {
	total := 0
	x := 0
	for x < len(input) {
		total, x = addNextValidProduct(input, total, x)
		println("total=", total)
	}
	return total
}

func addNextValidProduct(input string, total, x int) (int, int) {
	maxX := x + 12
	if len(input)-x < 4 || input[x:x+4] != "mul(" {
		return total, x + 1
	}
	x += 4
	a := -1
	b := -1
	nextNumber := ""
	for x < maxX {
		char := input[x]
		if (a == -1 && len(nextNumber) == 0) && char == ',' {
			return total, x + 1
		}
		if char == ',' {
			a = utils.Atoi(nextNumber)
			println("  found , a=", a)
			nextNumber = ""
			x++
			continue
		}
		if (a == -1 || len(nextNumber) == 0) && char == ')' {
			return total, x + 1
		}
		if char == ')' {
			b = utils.Atoi(nextNumber)
			break
		}
		_, err := strconv.Atoi(input[x : x+1])
		if err != nil {
			return total, x + 1
		}
		nextNumber += input[x : x+1]
		x++
	}
	return total + a*b, x + 1
}
