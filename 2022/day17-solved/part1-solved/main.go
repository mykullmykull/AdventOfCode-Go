package day

func main(winds string) int {
	c := chamber{
		mapStr:    []string{"+-------+"},
		winds:     winds,
		windIndex: -1,
		rocks:     rocks,
		rockIndex: -1,
	}

	count := 1
	for count < 2022 {
		c = c.getNextRock()
		c = c.getNextWind()
		for c.rock.isFalling() {
			c = c.moveRock()
		}
		// println()
		// println("droppedCount:", count, "height:", len(c.mapStr)-1)
		// c.printMap()
		count++
	}

	return len(c.mapStr) - 1 + 3
}
