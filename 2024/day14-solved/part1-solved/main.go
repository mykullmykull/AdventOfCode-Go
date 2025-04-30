package day

func main(input []string, width int, height int) int {
	robots := parse(input)
	quadrantCounts := [4]int{}
	for i, r := range robots {
		newX := (((r.pos.x + 100*r.vel.x) % width) + width) % width
		newY := (((r.pos.y + 100*r.vel.y) % height) + height) % height
		newRobot := robot{point{newX, newY}, r.vel}
		robots[i] = newRobot

		quadrant := getQuadrant(robots[i].pos, width, height)
		if quadrant < 0 {
			continue
		}
		quadrantCounts[quadrant]++
	}
	printGrid(robots, width, height)
	product := 1
	for _, count := range quadrantCounts {
		product *= count
	}
	return product
}

func getQuadrant(p point, width, height int) int {
	if p.x < width/2 && p.y < height/2 {
		return 0
	}
	if p.x > width/2 && p.y < height/2 {
		return 1
	}
	if p.x < width/2 && p.y > height/2 {
		return 2
	}
	if p.x > width/2 && p.y > height/2 {
		return 3
	}
	return -1
}
