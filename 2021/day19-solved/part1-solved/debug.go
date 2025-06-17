package day

import "fmt"

func (s scanner) printBeacons() {
	println("ID:", s.id)
	for beacon := range s.beacons {
		// Print each beacon's coordinates
		println("Beacon:", beacon[0], beacon[1], beacon[2])
	}
	fmt.Printf("Manipulator: r:%v, f: %v, t: %v", allRotations[s.m.rotationX], allFlips[s.m.flipationX], s.m.translation)
	println()
}

func printBeacons(beacons map[[3]int]bool) {
	id := 0
	for beacon := range beacons {
		// Print each beacon's coordinates
		println(id, "Beacon:", beacon[0], beacon[1], beacon[2])
		id++
	}
	println()
}

func (s scanner) printManipulator() {
	println("Manipulator for scanner ID:", s.id)
	fmt.Printf("Rotation: %v, Flipation: %v, Translation: %v\n", allRotations[s.m.rotationX], allFlips[s.m.flipationX], s.m.translation)
}

func printScanners(scanners []scanner) {
	println("Total scanners:", len(scanners))
	println("Scanners:")
	for _, s := range scanners {
		fmt.Printf("  ID: %d, Translation: %v\n", s.id, s.m.translation)
	}
	println()
}
