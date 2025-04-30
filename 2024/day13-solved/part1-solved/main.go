package day

import (
	"goAoc2024/utils"
	"math"
)

func main(input []string) int {
	machines := parse(input)
	total := 0
	for i, m := range machines {
		println()
		println("machine", i)
		cheapestWin := math.MaxInt
		won := false
		for a := 0; a < 100; a++ {
			for b := 0; b < 100; b++ {
				x := a*m.a.xChange + b*m.b.xChange
				y := a*m.a.yChange + b*m.b.yChange
				if x == m.prize.x && y == m.prize.y {
					won = true
					println("    win", a, b)
					println("    tokens", a*m.a.tokens+b*m.b.tokens)
					cheapestWin = utils.UpdateLeast(cheapestWin, a*m.a.tokens+b*m.b.tokens)
				}
			}
		}
		println("  cheapest win", cheapestWin)
		if won {
			total += cheapestWin
			println("total", total)
		}
	}
	return total
}
