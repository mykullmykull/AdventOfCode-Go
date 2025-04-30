package day

import (
	"goAoc2024/utils"
	"strconv"
)

func main(input string) int {
	total := 0
	x := 0
	shouldMultiply := true
	for x < len(input) {
		total, x, shouldMultiply = updateState(input, total, x, shouldMultiply)
	}
	return total
}

func updateState(input string, total, x int, shouldMultiply bool) (int, int, bool) {
	if len(input)-x < 7 {
		return total, x + 1, shouldMultiply
	}

	if input[x:x+4] == "do()" {
		return total, x + 4, true
	}

	if input[x:x+7] == "don't()" {
		return total, x + 7, false
	}

	if !shouldMultiply {
		return total, x + 1, shouldMultiply
	}

	if input[x:x+4] != "mul(" {
		return total, x + 1, shouldMultiply
	}
	x += 4
	a := -1
	b := -1
	nextNumber := ""
	maxX := x + 12
	for x < maxX {
		char := input[x]
		if (a == -1 && len(nextNumber) == 0) && char == ',' {
			return total, x + 1, shouldMultiply
		}
		if char == ',' {
			a = utils.Atoi(nextNumber)
			nextNumber = ""
			x++
			continue
		}
		if (a == -1 || len(nextNumber) == 0) && char == ')' {
			return total, x + 1, shouldMultiply
		}
		if char == ')' {
			b = utils.Atoi(nextNumber)
			break
		}
		_, err := strconv.Atoi(input[x : x+1])
		if err != nil {
			return total, x + 1, shouldMultiply
		}
		nextNumber += input[x : x+1]
		x++
	}
	return total + a*b, x + 1, shouldMultiply
}
