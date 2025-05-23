package utils

import (
	"fmt"
	"math"
	"strconv"
)

func StrToInt(str string) int {
	out, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Sprintf("cannot convert %s to an int", str))
	}
	return out
}

func UpdateMost(prev int, new int) int {
	if new > prev {
		return new
	}
	return prev
}

func UpdateLeast(prev int, new int) int {
	if new < prev {
		return new
	}
	return prev
}

func AbsForInt(n int) int {
	return int(math.Abs(float64(n)))
}

func AppendToStrArray(a []string, toAdd string) []string {
	copyOfArray := make([]string, len(a))
	copy(copyOfArray, a)
	copyOfArray = append(copyOfArray, toAdd)
	return copyOfArray
}

func JoinArray(a []string, joiner string) string {
	str := ""
	for _, s := range a {
		str += s
		str += joiner
	}
	return str
}

func JoinArrayOfRunes(a []rune, joiner rune) string {
	str := ""
	for _, r := range a {
		str += string(r)
		str += string(joiner)
	}
	return str
}

func JoinArrayOfInts(a []int, joiner string) string {
	str := ""
	for _, n := range a {
		str += strconv.Itoa(n)
		str += joiner
	}
	return str
}

func ContainsStr(a []string, s string) bool {
	for _, str := range a {
		if str == s {
			return true
		}
	}
	return false
}

func RuneToInt(r rune) int {
	return int(r) - 48
}

func ReplaceAtIndex(str string, index int, replacement string) string {
	return str[:index] + replacement + str[index+1:]
}

func DivideAndRoundUp(numerator int, denominator int) int {
	mod := numerator % denominator
	if mod == 0 {
		return numerator / denominator
	}
	remainder := denominator - (mod)
	return (numerator + remainder) / denominator
}

func StrToFloat(str string) float64 {
	out, err := strconv.ParseFloat(str, 64)
	if err != nil {
		panic(fmt.Sprintf("cannot convert %s to a float", str))
	}
	return out
}
