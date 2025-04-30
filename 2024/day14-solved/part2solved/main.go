package day

func main(input []string, width int, height int) int {
	robots := parse(input)
	seconds := 0
	for {
		if (seconds-72)%103 == 0 {
			println()
			println("seconds: ", seconds)
			printGrid(robots, width, height)
		}
		if seconds > 1000000 {
			break
		}
		for i, r := range robots {
			newX := (((r.pos.x + r.vel.x) % width) + width) % width
			newY := (((r.pos.y + r.vel.y) % height) + height) % height
			newRobot := robot{point{newX, newY}, r.vel}
			robots[i] = newRobot
		}
		seconds++
	}
	return seconds
}
