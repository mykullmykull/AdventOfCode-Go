package day

func reduce(testData []string) string {
	s := testData[0]
	for _, data := range testData[1:] {
		s = add([]string{s, data})
	}

	hasChanged := true
	for hasChanged {
		// println("Reducing: ", s)
		reduced := explodeAndSplit(s)
		if reduced == s {
			hasChanged = false
			continue
		}
		s = reduced
	}
	// println("Reduced:  ", s)
	return s
}

func explodeAndSplit(s string) string {
	hasChanged := true
	for hasChanged {
		// println("Exploding:", s)
		exploded := explode(s)
		if exploded == s {
			hasChanged = false
			continue
		}
		s = exploded
	}
	// println("Exploded: ", s)

	s = split(s)
	// println("Splitted: ", s)

	return s
}
