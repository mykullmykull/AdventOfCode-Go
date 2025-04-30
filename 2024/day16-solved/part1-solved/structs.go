package day

import (
	"fmt"
	"goAoc2024/utils"
)

const (
	up = iota
	dn
	lt
	rt
)

type racer struct {
	s     spot
	d     int
	score int
}

type spot struct {
	row int
	col int
}

func (p racer) getDistanceToEnd(end spot) int {
	return utils.Abs(p.s.row-end.row) + utils.Abs(p.s.col-end.col)
}

func (s spot) getVal(g []string) byte {
	return g[s.row][s.col]
}

func (r racer) isDeadEnd(g []string) bool {
	return r.s.getVal(g) == '.' && len(r.getValidAdjacents(g)) == 1
}

func (r racer) getValidAdjacents(g []string) []racer {
	adjacents := []racer{}
	upRacer := racer{s: spot{r.s.row - 1, r.s.col}, d: up, score: r.newScore(up)}
	dnRacer := racer{s: spot{r.s.row + 1, r.s.col}, d: dn, score: r.newScore(dn)}
	ltRacer := racer{s: spot{r.s.row, r.s.col - 1}, d: lt, score: r.newScore(lt)}
	rtRacer := racer{s: spot{r.s.row, r.s.col + 1}, d: rt, score: r.newScore(rt)}
	for _, adj := range []racer{upRacer, dnRacer, ltRacer, rtRacer} {
		if adj.isInBounds(g) && adj.isWalkable(g) {
			adjacents = append(adjacents, adj)
		}
	}
	return adjacents
}

func (r racer) newScore(d int) int {
	if r.d == d {
		return r.score + 1
	}
	return r.score + 1001
}

func (r racer) isInBounds(g []string) bool {
	return r.s.row >= 0 &&
		r.s.row < len(g) &&
		r.s.col >= 0 &&
		r.s.col < len(g[0])
}

func (r racer) isWalkable(g []string) bool {
	return r.s.getVal(g) != '#'
}

func (r racer) isBackwards(adj racer) bool {
	switch r.d {
	case up:
		return adj.d == dn
	case dn:
		return adj.d == up
	case lt:
		return adj.d == rt
	case rt:
		return adj.d == lt
	default:
		panic(fmt.Sprintf("invalid direction %d", r.d))
	}
}

func (s spot) toInt() int {
	return s.row*1000 + s.col
}
