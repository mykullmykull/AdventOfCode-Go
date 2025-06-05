package utils

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Atoi(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Sprintf("Could not parse %s to int", str))
	}
	return i
}

func Rtoi(r rune) int {
	return Atoi(string(r))
}

func Btoi(b byte) int {
	return Rtoi(rune(b))
}

func Itoa(i int) string {
	return strconv.Itoa(i)
}

func Abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

func SplitInts(str string, separator string) []int {
	var ints []int
	for _, s := range strings.Split(str, separator) {
		ints = append(ints, Atoi(s))
	}
	return ints
}

func CountTens(n int) int {
	if n == 0 {
		return 1
	}
	count := 0
	for n > 0 {
		n /= 10
		count++
	}
	return count
}

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

func AllBinaries(x int, y int) []string {
	var binaries []string
	length := Abs(x) + Abs(y)
	max := int(math.Pow(2, float64(length)))
	for z := 0; z < max; z++ {
		bin := fmt.Sprintf("%b", z)
		for len(bin) < length {
			bin = "0" + bin
		}
		if strings.Count(bin, "1") != Abs(x) || strings.Count(bin, "0") != Abs(y) {
			continue
		}
		binaries = append(binaries, bin)
	}
	return binaries
}

func BinToInt(bin string) int {
	i, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		panic(fmt.Sprintf("unable to parse int from bin %s, err: %s", bin, err.Error()))
	}
	return int(i)
}

func IntToBin(i int) string {
	return fmt.Sprintf("%b", i)
}

func Factorial(n int) int {
	pair := 1 + n
	extraForOdds := 0
	if n%2 == 1 {
		extraForOdds = n/2 + 1
	}
	return pair*(n/2) + extraForOdds
}
