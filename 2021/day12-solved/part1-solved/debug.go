package day

func printFinals(finals map[string]bool) {
	for path := range finals {
		println(path)
	}
	println()
}
