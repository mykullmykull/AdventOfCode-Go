package day14

func runB(image []string) int {
	// start := time.Now().UnixMilli()
	for i := 0; i < 1; i++ {
		image = moveRocksNorth(image)
		// image = moveRocksWest(image)
		// image = moveRocksSouth(image)
		// image = moveRocksEast(image)

		printImage(image)
		// if i%1000000 == 0 {
		// 	newStart := time.Now().UnixMilli()
		// 	lapsed := newStart - start
		// 	start = newStart
		// 	fmt.Printf("cycles: %d, ms lapsed: %d\n", i, lapsed)
		// }
	}
	return totalPressure(image)
}

func moveRocksNorth(image []string) []string {
	newImage := make([]string, len(image))
	for c := 0; c < len(image[0]); c++ {
		newCol := ""
		lastBlockRow := -1
		for r := 0; r < len(image); r++ {
			char := image[r][c]
			if char == 'O' {
				lastBlockRow++
				newCol += "O"
			}

			if char == '#' {
				lastBlockRow = r
				for len(newCol) < lastBlockRow-1 {
					newCol += "."
				}
				newCol += "#"
			}
		}
		newImage[c] = newCol
	}

	rotatedImage := make([]string, len(image))
	for c, col := range newImage {
		for _, char := range col {
			rotatedImage[c] += string(char)
		}
	}
	return rotatedImage
}

func moveRocksWest(image []string) []string {
	for r := 0; r < len(image); r++ {
		lastBlockCol := -1
		for c := 0; c < len(image[0]); c++ {
			char := image[r][c]
			if char == 'O' {
				lastBlockCol++
				if c != lastBlockCol {
					image = updateImage(image, point{r, lastBlockCol}, "O")
					image = updateImage(image, point{r, c}, ".")
				}
			}

			if char == '#' {
				lastBlockCol = c
			}
		}
	}
	return image
}

func moveRocksSouth(image []string) []string {
	for c := 0; c < len(image[0]); c++ {
		lastBlockRow := len(image)
		for r := len(image) - 1; r >= 0; r -= 1 {
			char := image[r][c]
			if char == 'O' {
				lastBlockRow -= 1
				if r != lastBlockRow {
					image = updateImage(image, point{lastBlockRow, c}, "O")
					image = updateImage(image, point{r, c}, ".")
				}
			}

			if char == '#' {
				lastBlockRow = r
			}
		}
	}
	return image
}

func moveRocksEast(image []string) []string {
	for r := 0; r < len(image); r++ {
		lastBlockCol := len(image[0])
		for c := len(image[0]) - 1; c >= 0; c -= 1 {
			char := image[r][c]
			if char == 'O' {
				lastBlockCol -= 1
				if c != lastBlockCol {
					image = updateImage(image, point{r, lastBlockCol}, "O")
					image = updateImage(image, point{r, c}, ".")
				}
			}

			if char == '#' {
				lastBlockCol = c
			}
		}
	}
	return image
}
