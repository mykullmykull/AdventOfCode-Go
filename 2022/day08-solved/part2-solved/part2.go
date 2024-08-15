package day

import "goAoc2022/utils"

func part2(input []string) int {
	p := pointer{
		visibleTrees: map[int]bool{},
	}
	maxScenicScore := 0
	for p.r = 0; p.r < len(input); p.r++ {
		for p.c = 0; p.c < len(input[0]); p.c++ {
			p.maxHeight = utils.StrToInt(string(input[p.r][p.c]))
			score := p.getScenicScore(input)
			if score > maxScenicScore {
				maxScenicScore = score
			}
		}
	}

	return maxScenicScore
}
