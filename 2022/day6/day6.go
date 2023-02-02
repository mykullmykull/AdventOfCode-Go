package day6

func FirstIndexAfterFourteenUnique(str string) int {
	firstIndex := -1
	for i := 0; i < len(str); i++ {
		if i < 13 {
			continue
		} else if isAllUnique(str[i-13 : i+1]) {
			firstIndex = i + 1
			break
		}
	}
	return firstIndex
}

func FirstIndexAfterFourUnique(str string) int {
	firstIndex := -1
	for i := 0; i < len(str); i++ {
		if i < 3 {
			continue
		} else if isAllUnique(str[i-3 : i+1]) {
			firstIndex = i + 1
			break
		}
	}
	return firstIndex
}

func isAllUnique(str string) bool {
	for i := 0; i < len(str); i++ {
		for j := 0; j < len(str); j++ {
			if i == j {
				continue
			} else if str[i] == str[j] {
				return false
			}
		}
	}
	return true
}
