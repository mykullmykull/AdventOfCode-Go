package day7

func runB(input []string) int {
	countValid := 0
	for _, row := range input {
		if isValidB(row) {
			countValid++
		}
	}
	return countValid
}

func isValidB(row string) bool {
	inside, outside := parseRow(row)
	for _, outStr := range outside {
		for _, inStr := range inside {
			if outsideABAAndInsideBAB(outStr, inStr) {
				return true
			}
		}
	}
	return false
}

func outsideABAAndInsideBAB(outStr string, inStr string) bool {
	for i := 0; i < len(outStr)-2; i++ {
		a := outStr[i]
		b := outStr[i+1]
		if a != b && a == outStr[i+2] {
			for j := 0; j < len(inStr)-2; j++ {
				if b == inStr[j] && a == inStr[j+1] && b == inStr[j+2] {
					return true
				}
			}
		}
	}
	return false
}

func runA(input []string) int {
	countValid := 0
	for _, row := range input {
		if isValid(row) {
			countValid++
		}
	}
	return countValid
}

func isValid(row string) bool {
	inside, outside := parseRow(row)
	if hasABBA(outside) && !hasABBA(inside) {
		return true
	}
	return false
}

func parseRow(row string) ([]string, []string) {
	inside := []string{}
	outside := []string{}
	str := ""
	for i := 0; i < len(row); i++ {
		char := string(row[i])
		if char == "[" {
			outside = append(outside, str)
			str = ""
			continue
		}
		if char == "]" {
			inside = append(inside, str)
			str = ""
			continue
		}
		str += char
	}
	outside = append(outside, str)
	return inside, outside
}

func hasABBA(strs []string) bool {
	for _, str := range strs {
		if isABBA(str) {
			return true
		}
	}
	return false
}

func isABBA(str string) bool {
	for i := 0; i < len(str)-3; i++ {
		if str[i] != str[i+1] && str[i] == str[i+3] && str[i+1] == str[i+2] {
			return true
		}
	}
	return false
}
