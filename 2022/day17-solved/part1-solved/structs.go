package day

import "fmt"

type point struct {
	row int
	col int
}

type rockPoints struct {
	points []point
}

type chamber struct {
	mapStr []string

	winds     string
	windIndex int
	wind      string

	rocks     [][]string
	rockIndex int
	rock      rockPoints
}

func (r rockPoints) isFalling() bool {
	return len(r.points) > 0
}

func (c chamber) getNextWind() chamber {
	c.windIndex = (c.windIndex + 1) % len(c.winds)
	c.wind = string(c.winds[c.windIndex])
	return c
}

func (c chamber) getNextRock() chamber {
	c.rockIndex = (c.rockIndex + 1) % len(c.rocks)
	rock := c.rocks[c.rockIndex]
	for x := 0; x < 3; x++ {
		c = c.getNextWind()
		rock = moveRock(rock, c.wind)
	}
	c.rock = rockPoints{}.getPoints(rock)

	c.mapStr = append(rock, c.mapStr...)
	return c
}

func moveRock(rock []string, wind string) []string {
	newRock := make([]string, len(rock))
	copy(newRock, rock)
	switch wind {
	case "<":
		for row := 0; row < len(rock); row++ {
			if rock[row][1] != '.' {
				return rock
			}
			for col := 0; col < len(rock[row]); col++ {
				if rock[row][col] == '@' {
					newRock[row] = newRock[row][:col-1] + "@." + newRock[row][col+1:]
				}
			}
		}
	case ">":
		for row := 0; row < len(rock); row++ {
			if rock[row][7] != '.' {
				return rock
			}
			for col := len(rock[row]) - 1; col > 0; col-- {
				if rock[row][col] == '@' {
					newRock[row] = newRock[row][:col] + ".@" + newRock[row][col+2:]
				}
			}
		}
	default:
		panic(fmt.Sprintf("invalid wind %s", wind))
	}
	return newRock
}

func (r rockPoints) getPoints(rock []string) rockPoints {
	points := []point{}
	for row, line := range rock {
		for col, char := range line {
			if char == '@' {
				points = append(points, point{row, col})
			}
		}
	}
	return rockPoints{points}
}

func (c chamber) moveRock() chamber {
	newRock := rockPoints{}
	switch c.wind {
	case "<":
		newRock = c.rock.moveLeft()
		c = c.updateRockIfValid(newRock)
		c.wind = "v"
		return c.moveRock()
	case ">":
		newRock = c.rock.moveRight()
		c = c.updateRockIfValid(newRock)
		c.wind = "v"
		return c.moveRock()
	case "v":
		newRock = c.rock.moveDown()
		isValid := newRock.isValid(c.mapStr)
		if !isValid {
			c = c.settleRock()
			c.rock = rockPoints{}
			return c
		}
		c.rock = newRock
		c = c.updateMap()
		c = c.getNextWind()
		return c.moveRock()
	default:
		panic(fmt.Sprintf("invalid wind %s", c.wind))
	}
}

func (r rockPoints) moveLeft() rockPoints {
	newPoints := []point{}
	for _, p := range r.points {
		newPoints = append(newPoints, point{p.row, p.col - 1})
	}
	return rockPoints{newPoints}
}

func (r rockPoints) moveRight() rockPoints {
	newPoints := []point{}
	for _, p := range r.points {
		newPoints = append(newPoints, point{p.row, p.col + 1})
	}
	return rockPoints{newPoints}
}

func (r rockPoints) moveDown() rockPoints {
	newPoints := []point{}
	for _, p := range r.points {
		newPoints = append(newPoints, point{p.row + 1, p.col})
	}
	return rockPoints{newPoints}
}

func (r rockPoints) isValid(mapStr []string) bool {
	for _, p := range r.points {
		if p.row < 0 ||
			p.row >= len(mapStr) ||
			p.col < 0 ||
			p.col >= len(mapStr[p.row]) ||
			mapStr[p.row][p.col] == '#' ||
			mapStr[p.row][p.col] == '|' ||
			mapStr[p.row][p.col] == '+' ||
			mapStr[p.row][p.col] == '-' {
			return false
		}
	}
	return true
}

func (c chamber) updateRockIfValid(newRock rockPoints) chamber {
	if newRock.isValid(c.mapStr) {
		c.rock = newRock
		c = c.updateMap()
		return c
	}
	return c
}

func (c chamber) updateMap() chamber {
	mapStr := make([]string, len(c.mapStr))
	for y, row := range c.mapStr {
		newRow := ""
		for x, char := range row {
			if char == '.' || char == '@' {
				newRow += "."
			} else {
				newRow += c.mapStr[y][x : x+1]
			}
		}
		mapStr[y] = newRow
	}

	for _, p := range c.rock.points {
		mapStr[p.row] = mapStr[p.row][:p.col] + "@" + mapStr[p.row][p.col+1:]
	}

	c.mapStr = mapStr
	return c
}

func (c chamber) settleRock() chamber {
	mapStr := make([]string, len(c.mapStr))
	copy(mapStr, c.mapStr)
	for _, p := range c.rock.points {
		mapStr[p.row] = mapStr[p.row][:p.col] + "#" + mapStr[p.row][p.col+1:]
	}

	// remove empty rows from the top
	for x := 0; x < len(mapStr); x++ {
		if mapStr[0] == "|.......|" {
			mapStr = mapStr[1:]
		}
	}

	c.mapStr = mapStr
	return c
}

func (c chamber) printMap() {
	for x, row := range c.mapStr {
		println(row)
		if x > 10 {
			break
		}
	}
	println()
}
