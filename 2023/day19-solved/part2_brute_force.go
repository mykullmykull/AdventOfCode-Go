package day19

import "fmt"

func runB_brute_force(input []string) int {
	piles, _ := parseInput(input)
	count := 0
	for x := 1; x <= 4000; x++ {
		for m := 1; m <= 4000; m++ {
			for a := 1; a <= 4000; a++ {
				for s := 1; s <= 4000; s++ {
					if a == 1 {
						fmt.Printf("x: %d, m: %d, a: %d, s: %d\n", x, m, a, s)
					}
					p := part{
						x:        x,
						m:        m,
						a:        a,
						s:        s,
						pileName: "in",
					}
					p = p.getFinalDestination(piles)
					if p.pileName == "A" {
						count++
					}
				}
			}
		}
	}
	return count
}
