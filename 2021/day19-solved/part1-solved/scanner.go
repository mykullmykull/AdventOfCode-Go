package day

type scanner struct {
	id      int
	beacons map[[3]int]bool // List of beacons in the scanner's range
	m       manipulator     // Manipulator for transforming beacon coordinates to s0's perspecitve
}

func (s *scanner) getTranslation(beacon1, beacon2 [3]int) {
	beacon2 = s.m.rotate(beacon2)
	beacon2 = s.m.flip(beacon2)
	s.m.translation = s.m.getTranslation(beacon1, beacon2)
}

func (s scanner) countMatches(beacons map[[3]int]bool) int {
	matches := 0

	// Iterate through each beacon in s2
	for b := range beacons {
		// does beacon match anything in s1?
		if _, exists := s.beacons[b]; exists {
			matches++
		}
	}

	return matches
}

func (unsolved *scanner) translateToS0(solved scanner) {
	unsolved.m.translation = solved.m.rotate(unsolved.m.translation)
	newRotation := solved.m.rotate(allRotations[unsolved.m.rotationX])
	for x, r := range allRotations {
		if r == newRotation {
			unsolved.m.rotationX = x
		}
	}

	newFlipation := solved.m.flip(allFlips[unsolved.m.flipationX])
	for x, f := range allFlips {
		if f == newFlipation {
			unsolved.m.flipationX = x
		}
	}
}
