package day

import (
	"fmt"
	"goAoc2022/utils"
	"strconv"
	"strings"
)

func main(input []string) int {
	monkeys := parse(input)
	rootSplit := strings.Split(monkeys["root"], " ")
	monkeyA := rootSplit[0]
	monkeyB := rootSplit[2]

	for isWaiting(monkeys[monkeyA]) {
		monkeys[monkeyA] = replaceNextMonkey(monkeys, monkeyA)
	}

	for isWaiting(monkeys[monkeyB]) {
		monkeys[monkeyB] = replaceNextMonkey(monkeys, monkeyB)
	}
	monkeys[monkeyA] = simplify(monkeys[monkeyA])
	monkeys[monkeyB] = simplify(monkeys[monkeyB])
	simplified := monkeys[monkeyA] + " = " + monkeys[monkeyB]
	solution := solve(simplified)
	return solution
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
		if isNumber(s) || isOperator(s) || s == "humn" {
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

func replaceNextMonkey(monkeys map[string]string, target string) string {
	value := monkeys[target]
	elements := removeParentheses(value)
	split := strings.Split(elements, " ")
	for _, s := range split {
		s = removeParentheses(s)
		if isNumber(s) || isOperator(s) || s == "humn" {
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

func simplify(value string) string {
	oldValue := ""
	for oldValue != value {
		nextOp := getNextStringToSimplify(value)
		oldValue = value
		value = strings.ReplaceAll(value, nextOp, doMath(nextOp))
	}
	return value
}

func getNextStringToSimplify(value string) string {
	start := -1
	end := -1
	nextOp := ""
	for x, c := range value {
		if c == ')' && start < 0 && x == len(value)-1 {
			return ""
		}
		if c == ')' && start < 0 {
			return getNextStringToSimplify(value[x+1:])
		}
		if c == '(' {
			start = x
		}
		if c == ')' {
			end = x + 1
			nextOp = value[start:end]
			break
		}
	}
	if nextOp == "" {
		return value
	}
	if strings.Contains(nextOp, "humn") {
		return getNextStringToSimplify(value[end:])
	}
	return nextOp
}

func doMath(value string) string {
	value = removeParentheses(value)
	split := strings.Split(value, " ")
	if len(split) == 1 {
		return value
	}
	if !isNumber(split[0]) || !isNumber(split[2]) {
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

func solve(str string) int {
	split := strings.Split(str, " = ")
	left := split[0]
	right := utils.StrToInt(split[1])
	for left != "humn" {
		outsideStr, insideStr := getOutsideStr(left)
		right = reverseOperation(outsideStr, right)
		left = insideStr
	}
	return right
}

func replaceMinuses(str string) string {
	newStr := ""
	for _, c := range str {
		if c == '-' {
			newStr += "+ (-1 *"
			continue
		}
		newStr += string(c)
	}
	return newStr
}

func getOutsideStr(str string) (string, string) {
	for x, c := range str {
		if x == 0 && c == '(' {
			break
		}
		if c == '(' {
			return str[:x], str[x+1 : len(str)-1]
		}
	}

	for x := len(str) - 1; x >= 0; x-- {
		if x == len(str)-1 && str[x] == ')' {
			break
		}
		if str[x] == ')' {
			return str[x+1:], str[1:x]
		}
	}

	split := strings.Split(str, " ")
	if split[0] == "humn" {
		return str[4:], "humn"
	}
	if split[2] == "humn" {
		return str[:len(str)], "humn"
	}

	panic("no parentheses or humn found in " + str)
}

func reverseOperation(str string, result int) int {
	str = strings.Trim(str, " ")
	split := strings.Split(str, " ")
	opp := split[1]
	num, errA := strconv.Atoi(split[0])
	if errA != nil {
		num, _ = strconv.Atoi(split[1])
		opp = split[0]
	}
	if split[1] == "-" {
		return -1 * (result - num)
	}

	switch opp {
	case "+":
		return result - num
	case "-":
		return result + num
	case "*":
		return result / num
	case "/":
		return result * num
	default:
		panic("unable to reverse " + str)
	}
}
