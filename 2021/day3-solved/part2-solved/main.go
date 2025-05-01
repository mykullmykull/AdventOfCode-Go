package day

import (
	"fmt"
	"strconv"
)

func main(input []string) int64 {
	o2Lines := make([]string, len(input))
	co2Lines := make([]string, len(input))
	copy(o2Lines, input)
	copy(co2Lines, input)

	o2Line := filterLines(o2Lines, true)
	co2Line := filterLines(co2Lines, false)

	o2, oOk := strconv.ParseInt(o2Line, 2, 64)
	co2, cOk := strconv.ParseInt(co2Line, 2, 64)
	fmt.Printf("o2: %s %d\n", o2Line, o2)
	fmt.Printf("co2: %s %d\n", co2Line, co2)

	if oOk != nil || cOk != nil {
		panic("Failed to convert binary to decimal")
	}

	return co2 * o2
}

func filterLines(lines []string, getMostCommon bool) string {
	for charIndex := range lines[0] {
		if len(lines) == 1 {
			return lines[0]
		}

		mostCommonBit := getMostCommonBit(lines, charIndex)
		expectedBit := maybeReverse(mostCommonBit, !getMostCommon)

		toKeep := []string{}
		for len(lines) > 0 {
			line := lines[0]
			lines = lines[1:]
			if line[charIndex] == expectedBit {
				toKeep = append(toKeep, line)
				continue
			}
		}
		lines = toKeep
	}
	if len(lines) != 1 {
		panic(fmt.Sprintf("More than one line left: %v", lines))
	}
	return lines[0]
}

func getMostCommonBit(input []string, index int) byte {
	count0s, count1s := 0, 0
	for _, line := range input {
		if line[index] == '0' {
			count0s++
		} else {
			count1s++
		}
	}

	bit := '1'
	if count0s > count1s {
		bit = '0'
	}
	return byte(bit)
}

func maybeReverse(bit byte, reverse bool) byte {
	if !reverse {
		return bit
	}
	if bit == '0' {
		return '1'
	}
	return '0'
}
