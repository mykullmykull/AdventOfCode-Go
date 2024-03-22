package day22

import "fmt"

type span struct {
	min int
	max int
}

func (s span) overlaps(s2 span) bool {
	return s.min <= s2.max && s2.min <= s.max
}

func (s span) toString() string {
	return fmt.Sprintf("%d-%d", s.min, s.max)
}

type brick struct {
	x span
	y span
	z span
}

func (b brick) equals(b2 brick) bool {
	return b.x == b2.x && b.y == b2.y && b.z == b2.z
}

func (b brick) overlaps(b2 brick) bool {
	return b.x.overlaps(b2.x) && b.y.overlaps(b2.y)
}

func (b brick) isSupporting(b2 brick) bool {
	return b.overlaps(b2) && b.z.max == b2.z.min-1
}

func (b brick) canFall(bricks []brick) bool {
	for _, b2 := range bricks {
		if b2.isSupporting(b) || b.z.min == 1 {
			return false
		}
	}
	return true
}

func (b brick) toString() string {
	return fmt.Sprintf("x: %v, y: %v, z: %v", b.x.toString(), b.y.toString(), b.z.toString())
}
