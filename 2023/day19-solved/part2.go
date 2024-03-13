package day19

import "fmt"

func runB(input []string) int {
	piles, _ := parseInput(input)
	count := 0
	acceptablePartSpans := getAllAcceptiblePartSpans(piles)
	for _, s := range acceptablePartSpans {
		count += s.totalPossibilities()
	}
	return count
}

func getAllAcceptiblePartSpans(piles map[string]pile) []partSpan {
	debug := false
	processing := []partSpan{{
		x:            span{min: 1, max: 4000},
		m:            span{min: 1, max: 4000},
		a:            span{min: 1, max: 4000},
		s:            span{min: 1, max: 4000},
		atPileName:   "",
		nextPileName: "A",
	}}
	finished := []partSpan{}
	for {
		log(("\n\nprocessing:"), debug)
		if debug {
			printPartSpans(processing)
		}
		log(("\nfinished:"), debug)
		if debug {
			printPartSpans(finished)
		}
		if len(processing) == 0 {
			break
		}
		nextSpan := processing[0]
		processing = processing[1:]

		if nextSpan.nextPileName == "in" {
			finished = append(finished, nextSpan)
			continue
		}

		for k, nextPile := range piles {
			newPartSpans := nextPile.targetedPartSpans(nextSpan, k)
			processing = append(processing, newPartSpans...)
		}
	}
	return finished
}

func printPartSpans(spans []partSpan) {
	for _, s := range spans {
		fmt.Println(s.toString())
	}
}
