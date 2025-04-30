package day

import "strings"

type valley struct {
	time int
	grid []string
	elfs map[point]bool
}

type point struct {
	r int
	c int
}

const (
	rt   = '>' // right-blowing blizzard
	lt   = '<' // left-blowing blizzard
	dn   = 'v' // down-blowing blizzard
	up   = '^' // up-blowing blizzard
	rl   = '-' // right and left
	rd   = '1' // right and down
	ru   = '2' // right and up
	rld  = '3' // right and left and down
	rlu  = '4' // right and left and up
	rdu  = '5' // right and down and up
	rldu = '6' // right and left and down and up
	ld   = '7' // left and down
	lu   = '8' // left and up
	ldu  = '9' // left and down and up
	du   = '|' // down and up
)

type combo struct {
	dir     rune
	current rune
}

func (v valley) getNewGrid() []string {
	height := len(v.grid)
	width := len(v.grid[0])
	newGrid := make([]string, height)
	newGrid[0] = "#." + strings.Repeat("#", width-2)
	for r := 1; r < height-1; r++ {
		newGrid[r] = "#" + strings.Repeat(".", width-2) + "#"
	}
	newGrid[height-1] = strings.Repeat("#", width-2) + ".#"
	return newGrid
}
