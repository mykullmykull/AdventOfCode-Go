package day

func main(input []string) int {
	machines := parse(input, 10000000000000)
	total := 0
	for _, m := range machines {
		b := (m.prize.x*m.a.yChange - m.prize.y*m.a.xChange) / (m.b.xChange*m.a.yChange - m.b.yChange*m.a.xChange)
		a := (m.prize.x - b*m.b.xChange) / m.a.xChange
		x := a*m.a.xChange + b*m.b.xChange
		y := a*m.a.yChange + b*m.b.yChange
		if x != m.prize.x || y != m.prize.y {
			continue
		}

		cost := a*m.a.tokens + b*m.b.tokens
		total += cost
	}
	return total
}
