package day22

import "fmt"

func runB(input []string) int {
	bricks := parseInput(input)
	bricks = finishFalling(bricks)
	printBricks(bricks)
	falls := countFalls(bricks)
	return falls
}

func countFalls(bricks []brick) int {
	count := 0
	bricksCopy := make([]brick, len(bricks))
	for i, b := range bricks {
		copy(bricksCopy, bricks)
		fmt.Printf("Checking brick %s %d/%d\n", b.toString(), i, len(bricksCopy))
		remaining := removeBrick(bricksCopy, i)

		for {
			stillFalling := false
			for j, b2 := range remaining {
				if b2.canFall(remaining) {
					count++
					remaining = removeBrick(remaining, j)
					stillFalling = true
					break
				}
			}
			if !stillFalling {
				break
			}
		}
	}
	return count
}

func removeBrick(bricks []brick, i int) []brick {
	return append(bricks[:i], bricks[i+1:]...)
}
