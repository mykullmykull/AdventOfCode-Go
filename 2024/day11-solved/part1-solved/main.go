package day

import (
	"fmt"
	"goAoc2024/utils"
)

func main(stones []string) int {
	for blink := 0; blink < 25; blink++ {
		newStones := []string{}
		for _, stone := range stones {
			if stone == "0" {
				newStones = append(newStones, "1")
				continue
			}

			if len(stone)%2 == 0 {
				half := len(stone) / 2
				newStones = append(newStones, stone[:half])
				secondHalf := removeLeadingZeros(stone[half:])
				newStones = append(newStones, secondHalf)
				continue
			}

			newStone := utils.Atoi(stone) * 2024
			newStones = append(newStones, fmt.Sprintf("%d", newStone))
		}
		stones = newStones
	}
	return len(stones)
}

func removeLeadingZeros(s string) string {
	for len(s) > 0 && s[0] == '0' {
		s = s[1:]
	}
	if len(s) == 0 {
		return "0"
	}
	return s
}
