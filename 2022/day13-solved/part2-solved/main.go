package day

func main(input []string) int {
	sets := parse(input)
	dSet1 := parseSet("[[2]]")
	dSet2 := parseSet("[[6]]")
	sets = append(sets, dSet1)
	sets = append(sets, dSet2)

	orderedSets := []set{}
	for _, s := range sets {
		orderedSets = s.insertInOrder(orderedSets)
	}
	dIndex1 := dSet1.getIndex(orderedSets)
	dIndex2 := dSet2.getIndex(orderedSets)
	return dIndex1 * dIndex2
}
