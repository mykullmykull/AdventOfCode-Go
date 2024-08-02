package day21

import "time"

func runB(input []string, steps int) int {
	countOddity := steps % 2
	count := 0
	g := grid{input}
	start := g.getStart()
	edges := []location{start}
	prev := map[int]map[int]int{}
	lastTime := time.Now().UnixNano()
	for i := 0; i < steps; i++ {
		newEdges := []location{}
		for _, l := range edges {
			toAdd := l.moveOne(g, prev)
			newEdges = append(newEdges, toAdd...)
			prev = addToPrev(toAdd, prev)
			if i%2 == countOddity {
				count += len(toAdd)
			}
		}

		prev = cleanPrev(prev)
		edges = newEdges

		newTime := time.Now().UnixNano()
		if i%100 == 0 {
			println("step: ", i, ", ms: ", (newTime-lastTime)/100000)
		}
		lastTime = newTime
	}
	return count
}

func addToPrev(toAdd []location, prev map[int]map[int]int) map[int]map[int]int {
	for _, l := range toAdd {
		if prev[l.r] == nil {
			prev[l.r] = map[int]int{}
		}
		prev[l.r][l.c] = 1
	}
	return prev
}

func cleanPrev(prev map[int]map[int]int) map[int]map[int]int {
	for r, row := range prev {
		for c, _ := range row {
			if prev[r][c] > 2 {
				delete(prev[r], c)
				continue
			}
			prev[r][c]++
		}
		if len(prev[r]) == 0 {
			delete(prev, r)
		}
	}
	return prev
}
