package day14

import "fmt"

func runB_brute_force(image []string) int {
	for i := 0; i < 1000000000; i++ {
		if i%1000000 == 0 {
			fmt.Printf("after %d cycles:\n", i)
		}
		image = rollRocksInACycle(image)
	}
	return totalPressure(image)
}

func rollRocksInACycle(image []string) []string {
	image = tiltPlatformNW("n", image)
	image = tiltPlatformNW("w", image)
	image = tiltPlatformSE("s", image)
	image = tiltPlatformSE("e", image)
	return image
}

func tiltPlatformNW(direction string, image []string) []string {
	for r, row := range image {
		for c, char := range row {
			if char != 'O' {
				continue
			}

			origin := point{r, c}
			image = moveRock(origin, direction, image)
		}
	}
	return image
}

func tiltPlatformSE(direction string, image []string) []string {
	for r := len(image) - 1; r >= 0; r-- {
		for c := len(image[0]) - 1; c >= 0; c-- {
			char := image[r][c]
			if char != 'O' {
				continue
			}

			origin := point{r, c}
			image = moveRock(origin, direction, image)
		}
	}
	return image
}

func moveRock(origin point, direction string, image []string) []string {
	destination := findNextFreePoint(origin, direction, image)
	if destination != origin {
		image = updateImage(image, destination, "O")
		image = updateImage(image, origin, ".")
	}
	return image
}

func findNextFreePoint(p point, direction string, image []string) point {
	rowChange := 0
	colChange := 0
	for {
		// fmt.Printf("\n  p: %v, direction: %s, len(image): %d, len(image[0]): %d\n", p, direction, len(image), len(image[0]))
		if (direction == "n" && p.row == 0) ||
			(direction == "w" && p.col == 0) ||
			(direction == "s" && p.row == len(image)-1) ||
			(direction == "e" && p.col == len(image[0])-1) {
			// fmt.Printf("  returning because p is at the edge\n")
			return p
		}

		switch direction {
		case "n":
			rowChange = -1
		case "w":
			colChange = -1
		case "s":
			rowChange = 1
		case "e":
			colChange = 1
		}

		char := image[p.row+rowChange][p.col+colChange]
		// fmt.Printf("  char: %c\n", char)
		if char != '.' {
			// fmt.Printf("  returning because char is not a '.'\n")
			return p
		}
		p.row += rowChange
		p.col += colChange
	}
}
