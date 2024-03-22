package day22

func runA(input []string) int {
	bricks := parseInput(input)
	bricks = finishFalling(bricks)
	printBricks(bricks)
	return countDisintegratable(bricks)
}

func finishFalling(bricks []brick) []brick {
	for {
		stillFalling := false
		for i, b := range bricks {
			if b.canFall(bricks) {
				stillFalling = true
				b.z.min--
				b.z.max--
				bricks[i] = b
			}
		}
		if !stillFalling {
			break
		}
	}
	return bricks
}

func countDisintegratable(bricks []brick) int {
	count := 0
	for _, b := range bricks {
		withoutB := []brick{}
		for _, b2 := range bricks {
			if b.equals(b2) {
				continue
			}
			withoutB = append(withoutB, b2)
		}
		isDisintegratable := true
		for _, b2 := range withoutB {
			if b2.canFall(withoutB) {
				isDisintegratable = false
				break
			}
		}
		if isDisintegratable {
			count++
		}
	}
	return count
}
