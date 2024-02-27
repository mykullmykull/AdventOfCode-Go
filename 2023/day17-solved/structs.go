package day17

import (
	"fmt"
	"goAoc2023/utils"
)

type point struct {
	row int
	col int
}

func (p point) isValid(grid []string) bool {
	return p.row >= 0 && p.row < len(grid) && p.col >= 0 && p.col < len(grid[0])
}

func (p point) getAdjacents(grid []string) []adjacent {
	up := adjacent{point{p.row - 1, p.col}, up}
	rt := adjacent{point{p.row, p.col + 1}, rt}
	dn := adjacent{point{p.row + 1, p.col}, dn}
	lt := adjacent{point{p.row, p.col - 1}, lt}
	candidates := []adjacent{up, rt, dn, lt}
	adjacents := []adjacent{}
	for _, c := range candidates {
		if c.location.isValid(grid) {
			adjacents = append(adjacents, c)
		}
	}
	return adjacents
}

func (p point) getHeatLoss(grid []string) int {
	return utils.Atoi(grid[p.row][p.col : p.col+1])
}

func (p point) atFinishLine(grid []string) bool {
	atFinish := p.row == len(grid)-1 && p.col == len(grid[0])-1
	return atFinish
}

type adjacent struct {
	location  point
	direction direction
}

func (a adjacent) isValidMove(prevState pathState) bool {
	return a.direction != prevState.facing || prevState.straightRepeats < 3
}

func (a adjacent) isValidUltraMove(prevState pathState, grid []string) bool {
	return a.straightForAtLeast4(prevState) &&
		a.straightForNoMoreThan10(prevState) &&
		a.finishAfterAtLeast4Straight(prevState, grid)
}

func (a adjacent) straightForAtLeast4(prevState pathState) bool {
	st4 := a.direction == prevState.facing || prevState.straightRepeats >= 4
	return st4
}

func (a adjacent) straightForNoMoreThan10(prevState pathState) bool {
	st10 := a.direction != prevState.facing || prevState.straightRepeats < 10
	return st10
}

func (a adjacent) finishAfterAtLeast4Straight(prevState pathState, grid []string) bool {
	fin4 := !a.location.atFinishLine(grid) || (a.direction == prevState.facing && prevState.straightRepeats >= 3)
	return fin4
}

type direction struct {
	dRow int
	dCol int
}

var up = direction{-1, 0}
var rt = direction{0, 1}
var dn = direction{1, 0}
var lt = direction{0, -1}

func (d direction) opposite() direction {
	return direction{d.dRow * -1, d.dCol * -1}
}

func (d direction) toString() string {
	switch d {
	case up:
		return "up"
	case rt:
		return "rt"
	case dn:
		return "dn"
	case lt:
		return "lt"
	}
	return "unknown"
}

type pathState struct {
	location        point
	facing          direction
	straightRepeats int
}

func (s pathState) getAdjacents(grid []string) []adjacent {
	candidates := s.location.getAdjacents(grid)
	adjacents := []adjacent{}
	for _, c := range candidates {
		if c.direction == s.facing.opposite() {
			continue
		}
		adjacents = append(adjacents, c)
	}
	return adjacents
}

func (s pathState) toString() string {
	return fmt.Sprintf("location: %v, facing: %v, repeats: %d", s.location, s.facing.toString(), s.straightRepeats)
}

type queueItem struct {
	state    pathState
	heatLoss int
}

func (q queueItem) toString() string {
	return fmt.Sprintf("%v, heatLoss: %d", q.state.toString(), q.heatLoss)
}
