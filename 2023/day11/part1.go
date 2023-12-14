package day11

import "math"

type point struct {
	row int
	col int
}

func runA(image []string) int {
	image = expandEmptyRows(image)
	image = expandEmptyCols(image)
	galaxies := findAllGalaxies(image)
	return sumAllDistances(galaxies)
}

func expandEmptyRows(image []string) []string {
	newImage := []string{}
	for _, row := range image {
		isEmpty := true
		for _, char := range row {
			if char != '.' {
				isEmpty = false
				break
			}
		}
		newImage = append(newImage, row)
		if isEmpty {
			newImage = append(newImage, row)
		}
	}
	return newImage
}

func expandEmptyCols(image []string) []string {
	newImage := make([]string, len(image))
	for i, _ := range image[0] {
		isEmpty := true
		for j, row := range image {
			char := row[i : i+1]
			newImage[j] += char
			if row[i] != '.' {
				isEmpty = false
			}
		}
		if isEmpty {
			for j, row := range image {
				char := row[i : i+1]
				newImage[j] += char
			}
		}
	}
	return newImage
}

func findAllGalaxies(image []string) []point {
	galaxies := []point{}
	for r := 0; r < len(image); r++ {
		for c := 0; c < len(image[0]); c++ {
			char := image[r][c]
			if char == '#' {
				galaxies = append(galaxies, point{r, c})
			}
		}
	}
	return galaxies
}

func sumAllDistances(galaxies []point) int {
	sum := 0
	for i, g1 := range galaxies {
		for j, g2 := range galaxies {
			if i <= j {
				continue
			}
			sum += getDistance(g1, g2)
		}
	}
	return sum
}

func getDistance(g1 point, g2 point) int {
	return int(
		math.Abs(float64(g1.row)-float64(g2.row)) +
			math.Abs(float64(g1.col)-float64(g2.col)))
}

/*
func findNextGalaxy(p point, image []string) point {
	for r := p.row; r < len(image); r++ {
		for c := 0; c < len(image[0]); c++ {
			// fmt.Printf("r: %d, c: %d\n", r, c)
			if r == p.row && c <= p.col {
				continue
			}

			char := image[r][c]
			if char == '#' {
				return point{r, c}
			}
		}
	}
	return point{-1, -1}
}

func findDistancesToLaterGalaxies(p point, image []string) int {
	sum := findDistancesLaterOnSameRow(p, image)
	if p.col == 0 {
		// fmt.Printf("  sum: %d\n", sum)
		// printImage(image)
	}

	pointsToCheck := []point{}
	// check dn first if not at last row
	if p.row != len(image)-1 {
		pointsToCheck = append(pointsToCheck, point{row: p.row + 1, col: p.col})
	}

	distance := 1
	for {
		if len(pointsToCheck) == 0 {
			break
		}

		sum += distancesToFoundGalaxies(distance, pointsToCheck, image)
		image = markCheckedPoints(image, pointsToCheck)

		// expandSearch needs to update image to keep track of points and not create duplicates
		// but not persist the changes to this image
		imageCopy := make([]string, len(image))
		copy(imageCopy, image)
		pointsToCheck = expandSearch(distance, pointsToCheck, imageCopy)

		distance++
	}
	return sum
}

func findDistancesLaterOnSameRow(p point, image []string) int {
	sum := 0
	distance := 1
	for c := p.col + 1; c < len(image[0]); c++ {
		if image[p.row][c] == '#' {
			// fmt.Printf("  same row, {%d %d}, distance: %d\n", p.row, c, distance)
			sum += distance
		}
		distance++
	}
	return sum
}

func distancesToFoundGalaxies(distance int, pointsToCheck []point, image []string) int {
	sum := 0
	for _, p := range pointsToCheck {
		if image[p.row][p.col] == '#' {
			// fmt.Printf("  lower: %v, distance: %d\n", p, distance)
			sum += distance
		}
	}
	return sum
}

func markCheckedPoints(image []string, points []point) []string {
	for _, p := range points {
		row := image[p.row]
		row = row[0:p.col] + "+" + row[p.col+1:]
		image[p.row] = row
	}
	return image
}

func expandSearch(distance int, pointsToCheck []point, image []string) []point {
	newPoints := []point{}
	for _, p := range pointsToCheck {
		dn := point{row: p.row + 1, col: p.col}
		rt := point{row: p.row, col: p.col + 1}
		lt := point{row: p.row, col: p.col - 1}
		possiblePoints := []point{dn, rt, lt}
		validPoints := validatePoints(possiblePoints, image)

		newPoints = append(newPoints, validPoints...)
		image = markCheckedPoints(image, validPoints)
	}
	return newPoints
}

func validatePoints(points []point, image []string) []point {
	validPoints := []point{}
	for _, p := range points {
		if p.row < 0 ||
			p.col < 0 ||
			p.row >= len(image) ||
			p.col >= len(image[0]) ||
			image[p.row][p.col] == '+' {
			continue
		}
		validPoints = append(validPoints, p)
	}
	return validPoints
}

func advanceToNextPoint(p point, maxCol int) point {
	if p.col+1 > maxCol {
		p.row++
	} else {
		p.col++
	}
	return p
}

func printImage(image []string) {
	for _, row := range image {
		fmt.Println(row)
	}
	fmt.Println()
}
*/
