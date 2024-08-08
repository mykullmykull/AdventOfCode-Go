package day1

import "goAoc2022/utils"

func part1(input []string) int {
	mostCalories := 0
	caloriesByElf := 0
	for _, food := range input {
		if food == "" {
			mostCalories = utils.UpdateMost(mostCalories, caloriesByElf)
			caloriesByElf = 0
			continue
		}
		caloriesByElf += utils.StrToInt(food)
	}
	return mostCalories
}
