package day

import (
	"fmt"
	"goAoc2024/utils"
)

func printGrid(robots []robot, width int, height int) {
	println("printing grid with width: ", width, " and height: ", height)
	grid := make([]string, height)
	for row := 0; row < height; row++ {
		grid[row] = ""
		for col := 0; col < width; col++ {
			grid[row] += "0"
		}
	}
	for _, r := range robots {
		count := utils.Atoi(grid[r.pos.y][r.pos.x : r.pos.x+1])
		newCount := fmt.Sprint(count + 1)
		grid[r.pos.y] = grid[r.pos.y][:r.pos.x] + newCount + grid[r.pos.y][r.pos.x+1:]
	}
	for _, row := range grid {
		println(row)
	}
}
