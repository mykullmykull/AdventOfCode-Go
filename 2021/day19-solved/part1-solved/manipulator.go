package day

type manipulator struct {
	rotationX   int    // which index should each coord be rotated to
	flipationX  int    // which direction should each coord be flipped to
	translation [3]int // what to add to each coord after rotation to match s1's perspective
}

func (m manipulator) rotate(beacon [3]int) [3]int {
	rotated := [3]int{}
	rotation := allRotations[m.rotationX]

	// Apply the rotation pattern to each coordinate
	rotated[0] = beacon[rotation[0]]
	rotated[1] = beacon[rotation[1]]
	rotated[2] = beacon[rotation[2]]

	return rotated
}

func (m manipulator) flip(beacon [3]int) [3]int {
	flipped := [3]int{}
	flip := allFlips[m.flipationX]

	// Apply the flip pattern to each coordinate
	flipped[0] = beacon[0] * flip[0]
	flipped[1] = beacon[1] * flip[1]
	flipped[2] = beacon[2] * flip[2]

	return flipped
}

func (m manipulator) translate(beacon [3]int) [3]int {
	translated := [3]int{}

	// Add the translation vector to the beacon coordinates
	translated[0] = beacon[0] + m.translation[0]
	translated[1] = beacon[1] + m.translation[1]
	translated[2] = beacon[2] + m.translation[2]

	return translated
}

func (m manipulator) getTranslation(beacon1, beacon2 [3]int) [3]int {
	translation := [3]int{}

	// Calculate the translation vector from beacon2 to beacon1
	translation[0] = beacon1[0] - beacon2[0]
	translation[1] = beacon1[1] - beacon2[1]
	translation[2] = beacon1[2] - beacon2[2]

	return translation
}

func (m manipulator) manipulateBeacons(beacons map[[3]int]bool) map[[3]int]bool {
	newBeacons := make(map[[3]int]bool)

	for b := range beacons {
		// Apply the manipulator to each beacon
		rotated := m.rotate(b)
		flipped := m.flip(rotated)
		translated := m.translate(flipped)
		newBeacons[translated] = true
	}
	return newBeacons
}
