package day

import "fmt"

type trench struct {
	scanners       []scanner
	solvedScanners map[int]bool
	beacons        map[[3]int]bool // Unique beacons across all scanners
}

func (t *trench) solveScanners() {
	for unsolvedX := range t.scanners {

		// Skip already solved scanners
		if _, ok := t.solvedScanners[unsolvedX]; ok {
			continue
		}

		t.solveScanner(unsolvedX)
	}
}

func (t *trench) solveScanner(unsolvedX int) {
	for solvedX := range t.solvedScanners {
		fmt.Printf("  solved: %d, trying scanner %d with solved scanner %d\n", len(t.solvedScanners), unsolvedX, solvedX)
		t.tryToMatch(unsolvedX, solvedX)
		if t.solvedScanners[unsolvedX] {
			return
		}
	}
}

func (t *trench) tryToMatch(unsolvedX, solvedX int) {
	for r := range allRotations {
		t.scanners[unsolvedX].m.rotationX = r
		// println("    trying rotation", r)
		t.tryRotation(unsolvedX, solvedX)

		if t.solvedScanners[unsolvedX] {
			return
		}
	}
}

func (t *trench) tryRotation(unsolvedX, solvedX int) {
	for f := range allFlips {
		t.scanners[unsolvedX].m.flipationX = f
		// println("      trying flipation", f)
		t.tryFlipation(unsolvedX, solvedX)
		if t.solvedScanners[unsolvedX] {
			return
		}
	}
}

func (t *trench) tryFlipation(unsolvedX, solvedX int) {
	solved := t.scanners[solvedX]
	unsolved := t.scanners[unsolvedX]
	for b1 := range solved.beacons {
		for b2 := range unsolved.beacons {
			unsolved.getTranslation(b1, b2)
			t.tryTranslation(unsolved, solved)
			if t.solvedScanners[unsolvedX] {
				return
			}
		}
	}
}

func (t *trench) tryTranslation(unsolved, solved scanner) {
	// Apply the manipulator to beacons
	manipulatedBeacons := unsolved.m.manipulateBeacons(unsolved.beacons)

	// Count how many beacons now match
	matches := solved.countMatches(manipulatedBeacons)

	// Do we have a match?
	if matches >= 12 {
		println(".   Found", matches, "matches between scanner", unsolved.id, "and solved scanner", solved.id)
		t.scanners[unsolved.id] = unsolved
		t.scanners[unsolved.id].beacons = manipulatedBeacons
		t.addBeacons(manipulatedBeacons)
		t.solvedScanners[unsolved.id] = true
	}
}

func (t *trench) addBeacons(beacons map[[3]int]bool) {
	for beacon := range beacons {
		t.beacons[beacon] = true
	}
}
