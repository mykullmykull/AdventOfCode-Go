package day

func main(input []string) int {
	t := trench{
		scanners:       parse(input),
		solvedScanners: map[int]bool{},
		beacons:        map[[3]int]bool{},
	}

	// Mark scanner 0 as solved
	t.solvedScanners[0] = true

	// Add beacons from scanner 0
	t.addBeacons(t.scanners[0].beacons)

	for len(t.solvedScanners) < len(t.scanners) {
		println("Solving scanner ", len(t.solvedScanners))
		t.solveScanners()
	}
	printScanners(t.scanners)
	return len(t.beacons)
}
