package day

import "fmt"

type steam struct {
	grid    [][][]rune
	cursors []cursor
	count   int
}

type cursor struct {
	x int
	y int
	z int
}

func (s steam) get(c cursor) rune {
	return s.grid[c.z][c.y][c.x]
}

func (s steam) set(c cursor, r rune) steam {
	s.grid[c.z][c.y][c.x] = r
	return s
}

func (s steam) countExposedFaces() steam {
	for len(s.cursors) > 0 {
		c := s.cursors[0]
		s = s.set(c, '0')
		s.cursors = s.cursors[1:]
		for _, new := range c.expand() {
			if s.isOutOfBounds(new) {
				continue
			}
			if s.get(new) == '0' {
				continue
			}
			if s.get(new) == '.' {
				s.cursors = append(s.cursors, new)
				s.set(new, '0')
				continue
			}
			if s.get(new) == '#' {
				s.count++
				continue
			}
			panic("unexpected rune at cursor" + new.toString() + string(s.get(new)))
		}
	}
	return s
}

func (c cursor) expand() []cursor {
	posX := cursor{c.x + 1, c.y, c.z}
	negX := cursor{c.x - 1, c.y, c.z}
	posY := cursor{c.x, c.y + 1, c.z}
	negY := cursor{c.x, c.y - 1, c.z}
	posZ := cursor{c.x, c.y, c.z + 1}
	negZ := cursor{c.x, c.y, c.z - 1}
	return []cursor{posX, negX, posY, negY, posZ, negZ}
}

func (c cursor) toString() string {
	return fmt.Sprintf("x%d y%d z%d", c.x, c.y, c.z)
}

func (s steam) isOutOfBounds(c cursor) bool {
	return c.x < 0 || c.y < 0 || c.z < 0 || c.x >= len(s.grid[0][0]) || c.y >= len(s.grid[0]) || c.z >= len(s.grid)
}
