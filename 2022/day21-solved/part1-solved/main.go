package day

import (
	"fmt"
	"goAoc2022/utils"
	"strconv"
	"strings"
)

func main(input []string) int {
	monkeys := parse(input)
	for isWaiting(monkeys["root"]) {
		println(monkeys["root"])
		monkeys["root"] = replaceNextMonkey(monkeys)
	}
	return doMath(monkeys["root"])
}

func parse(input []string) map[string]string {
	monkeys := make(map[string]string)
	for _, line := range input {
		split := strings.Split(line, ": ")
		monkeys[split[0]] = split[1]
	}
	return monkeys
}

func isWaiting(value string) bool {
	value = removeParentheses(value)
	split := strings.Split(value, " ")
	for _, s := range split {
		if isNumber(s) || isOperator(s) {
			continue
		}
		return true
	}
	return false
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func isOperator(s string) bool {
	return s == "+" || s == "-" || s == "*" || s == "/"
}

func replaceNextMonkey(monkeys map[string]string) string {
	value := monkeys["root"]
	elements := removeParentheses(value)
	split := strings.Split(elements, " ")
	for _, s := range split {
		s = removeParentheses(s)
		if isNumber(s) || isOperator(s) {
			continue
		}
		value = strings.ReplaceAll(value, s, "("+monkeys[s]+")")
		return value
	}
	return value
}

func removeParentheses(value string) string {
	value = strings.ReplaceAll(value, "(", "")
	value = strings.ReplaceAll(value, ")", "")
	return value
}

func doMath(value string) int {
	for !isNumber(value) {
		println(value)
		nextOp := getNextOperationToSimplify(value)
		value = strings.ReplaceAll(value, nextOp, simplify(nextOp))
	}
	return utils.StrToInt(value)
}

func getNextOperationToSimplify(value string) string {
	start := -1
	for x, c := range value {
		if c == '(' {
			start = x
		}
		if c == ')' {
			return value[start : x+1]
		}
	}
	return value
}

func simplify(value string) string {
	value = removeParentheses(value)
	split := strings.Split(value, " ")
	if len(split) == 1 {
		return value
	}
	a := utils.StrToInt(split[0])
	b := utils.StrToInt(split[2])
	result := 0
	switch split[1] {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unknown operator" + value)
	}
	return fmt.Sprintf("%d", result)
}
