package day22

func printBricks(bricks []brick) {
	for _, b := range bricks {
		println(b.toString())
	}
}
