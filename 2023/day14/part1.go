package day14

import "fmt"

type point struct {
	row int
	col int
}

func runA(image []string) int {
	// for _, line := range image {
	// 	fmt.Println(line)
	// }
	image = rollRocksNorth(image)
	// for _, line := range image {
	// 	fmt.Println(line)
	// }
	return totalPressure(image)
}

func rollRocksNorth(image []string) []string {
	for r, row := range image {
		for c, char := range row {
			if char != 'O' {
				continue
			}

			p := point{r, c}
			newRow := getNextFreeNorthRow(p, image)
			// fmt.Printf("  r: %d, c: %d, char: %c, newRow: %d\n", r, c, char, newRow)
			if newRow != r {
				newPoint := point{newRow, c}
				image = updateImage(image, newPoint, "O")
				image = updateImage(image, p, ".")
			}
		}
	}
	return image
}

func getNextFreeNorthRow(p point, image []string) int {
	r := p.row
	c := p.col
	for r > 0 {
		char := image[r-1][c]
		if char != '.' {
			return r
		}
		r--
	}
	return r
}

func updateImage(image []string, p point, newChar string) []string {
	row := image[p.row]
	row = row[:p.col] + newChar + row[p.col+1:]
	image[p.row] = row
	return image
}

func totalPressure(image []string) int {
	sum := 0
	for r, row := range image {
		for _, char := range row {
			if char != 'O' {
				continue
			}
			sum += len(image) - r
		}
	}
	return sum
}

func printImage(image []string) {
	fmt.Println()
	for _, line := range image {
		fmt.Println(line)
	}
	fmt.Println()
}
