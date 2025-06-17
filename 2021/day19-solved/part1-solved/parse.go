package day

import "fmt"

func parse(input []string) []scanner {
	scanners := []scanner{}
	currentScanner := scanner{
		id:      0,
		beacons: make(map[[3]int]bool),
		m:       manipulator{},
	}

	for _, line := range input {
		if line == "" {
			scanners = append(scanners, currentScanner)
			currentScanner = scanner{
				id:      currentScanner.id + 1,
				beacons: make(map[[3]int]bool),
				m:       manipulator{},
			}
			continue
		}
		if line[:3] == "---" {
			continue
		}
		var x, y, z int
		fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		currentScanner.beacons[[3]int{x, y, z}] = true
	}
	return scanners
}
