package day

type fold struct {
	dimension string
	distance  int
}

type dot struct {
	x int
	y int
}

func main(input []string) int {
	dots, folds := parse(input)

	for _, f := range folds {
		for d := range dots {
			newDot := foldDot(d, f)
			delete(dots, d) // Remove the original dot
			dots[newDot] = true
		}
	}

	printDots(dots)
	return len(dots)
}

func foldDot(d dot, f fold) dot {
	switch f.dimension {
	case "x":
		if d.x > f.distance {
			d.x = f.distance - (d.x - f.distance)
		}
	case "y":
		if d.y > f.distance {
			d.y = f.distance - (d.y - f.distance)
		}
	default:
		panic("Invalid fold dimension")
	}

	return d
}

func printDots(dots map[dot]bool) {
	maxX, maxY := 0, 0
	for d := range dots {
		if d.x > maxX {
			maxX = d.x
		}
		if d.y > maxY {
			maxY = d.y
		}
	}

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if dots[dot{x, y}] {
				print("#")
				continue
			}

			print(".")
		}
		println()
	}
}
