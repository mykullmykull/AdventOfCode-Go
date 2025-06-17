package day

func reduce(testData []string) string {
	s := testData[0]
	for _, data := range testData[1:] {
		s = add([]string{s, data})
	}

	hasChanged := true
	for hasChanged {
		reduced := explodeAndSplit(s)
		if reduced == s {
			hasChanged = false
			continue
		}
		s = reduced
	}
	return s
}

func explodeAndSplit(s string) string {
	hasChanged := true
	for hasChanged {
		exploded := explode(s)
		if exploded == s {
			hasChanged = false
			continue
		}
		s = exploded
	}

	s = split(s)
	return s
}
