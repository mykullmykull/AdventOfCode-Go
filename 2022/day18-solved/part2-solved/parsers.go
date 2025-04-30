package day

import (
	"goAoc2022/utils"
	"strings"
)

func (s steam) parse(input []string) steam {
	pixels := parsePixels(input)
	s = s.parseGrid(pixels)
	s = s.fillGrid(pixels)
	s.cursors = []cursor{{0, 0, 0}}
	return s
}

func parsePixels(input []string) []cursor {
	pixels := make([]cursor, len(input))
	for n, line := range input {
		split := strings.Split(line, ",")
		pixels[n] = cursor{
			utils.StrToInt(split[0]) + 1,
			utils.StrToInt(split[1]) + 1,
			utils.StrToInt(split[2]) + 1,
		}
	}
	return pixels
}

func (s steam) parseGrid(pixels []cursor) steam {
	depth, height, width := getMaxPlus2(pixels)
	s.grid = make([][][]rune, depth)
	for z := 0; z < depth; z++ {
		s.grid[z] = make([][]rune, height)
		for y := 0; y < height; y++ {
			s.grid[z][y] = make([]rune, width)
			for x := 0; x < width; x++ {
				s.grid[z][y][x] = '.'
			}
		}
	}
	return s
}

func (s steam) fillGrid(pixels []cursor) steam {
	for _, pixel := range pixels {
		s.set(pixel, '#')
	}
	return s
}

func getMaxPlus2(pixels []cursor) (int, int, int) {
	depth, height, width := 0, 0, 0
	for _, pixel := range pixels {
		depth = utils.UpdateMost(depth, pixel.z)
		height = utils.UpdateMost(height, pixel.y)
		width = utils.UpdateMost(width, pixel.x)
	}
	return depth + 3, height + 3, width + 3
}

func (s steam) printGrid() {
	for y := 0; y < len(s.grid[0]); y++ {
		for z := 0; z < len(s.grid); z++ {
			for x := 0; x < len(s.grid[0][0]); x++ {
				print(string(s.grid[z][y][x]))
			}
			print("     ")
		}
		println()
	}
	println()
}

func (s steam) printCursors() {
	for _, c := range s.cursors {
		print(c.toString())
		print("     ")
	}
	println()
}
