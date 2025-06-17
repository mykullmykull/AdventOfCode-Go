package day

import (
	"fmt"
)

func add(numbers []string) string {
	if len(numbers) != 2 {
		panic("add function requires exactly two numbers")
	}
	return fmt.Sprintf("[%s,%s]", numbers[0], numbers[1])
}
