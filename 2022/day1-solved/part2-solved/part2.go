package day1

import "goAoc2022/utils"

func part2(input []string) int {
	calories1st := 0
	calories2nd := 0
	calories3rd := 0
	caloriesByElf := 0
	for _, food := range input {
		if food == "" {
			if caloriesByElf < calories3rd {
				caloriesByElf = 0
				continue
			}
			if caloriesByElf < calories2nd {
				calories3rd = caloriesByElf
				caloriesByElf = 0
				continue
			}
			calories3rd = calories2nd
			calories2nd = caloriesByElf
			if caloriesByElf > calories1st {
				calories2nd = calories1st
				calories1st = caloriesByElf
			}
			caloriesByElf = 0
			continue
		}
		caloriesByElf += utils.StrToInt(food)
	}
	return calories1st + calories2nd + calories3rd
}
