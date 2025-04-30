package day

import (
	"fmt"
	"math"
)

func main(input []string) string {
	totalDecimal := 0
	for _, snafu := range input {
		decimal := convertToDecimal(snafu)
		println("snafu", snafu, "decimal", decimal)
		totalDecimal += decimal
	}
	return "0"
}

func convertToDecimal(snafu string) int {
	decimal := 0
	for x := 0; x < len(snafu); x++ {
		index := len(snafu) - x - 1
		fives := int(math.Pow(5, float64(x)))
		char := snafu[index]
		digit := 0
		switch char {
		case '2':
			digit = 2
		case '1':
			digit = 1
		case '0':
			digit = 0
		case '-':
			digit = -1
		case '=':
			digit = -2
		default:
			panic(fmt.Sprintf("invalid snafu char: %c\n", char))
		}
		decimal += digit * fives
		// fmt.Printf("x %d, index %d, fives %d, char %c, digit %d, toAdd %d\n", x, index, fives, char, digit, digit*fives)
	}
	return decimal
}

func convertToSnafu(decimal int) string {
	println("decimal", decimal)
	snafu := ""
	fives := 5
	remaining := decimal
	for remaining > 0 {
		println("  remaining", remaining, "snafu", snafu)
		toAdd := remaining % fives
		println("    toAdd before", toAdd)
		if toAdd > 2*fives/5 {
			toAdd = toAdd - fives
		}
		println("    toAdd after", toAdd)
		remaining -= toAdd
		println("    new remaining", remaining)
		switch toAdd {
		case 2:
			snafu = "2" + snafu
		case 1:
			snafu = "1" + snafu
		case 0:
			snafu = "0" + snafu
		case -1:
			snafu = "-" + snafu
		case -2:
			snafu = "=" + snafu
		default:
			panic(fmt.Sprintf("invalid toAdd: %d\n", toAdd))
		}
		fives *= 5
	}
	println("  dec", decimal, "= snafu", snafu)
	return snafu
}
