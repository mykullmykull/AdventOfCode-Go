package day18

import "fmt"

type horizontal struct {
	row int
	min int
	max int
}

func (h horizontal) width() int {
	return h.max - h.min + 1
}

func (h horizontal) updateWith(other horizontal) (bool, []horizontal) {
	debug := true
	log(fmt.Sprintf("  updating active %s", h.toStringWithRow()), debug)
	log(fmt.Sprintf("  by switching new %s", other.toStringWithRow()), debug)
	if !h.overlaps(other) {
		log("    they don't overlap", debug)
		return false, []horizontal{h}
	}

	if other.contains(h.min) && other.contains(h.max) {
		log("    new totally encompasses h", debug)
		return true, []horizontal{}
	}

	if h.min == other.max {
		log("    new stretches h left", debug)
		return true, []horizontal{{row: h.row, min: other.min, max: h.max}}
	}

	if h.min == other.min {
		log("    new shrinks h from the left", debug)
		return true, []horizontal{{row: h.row, min: other.max, max: h.max}}
	}

	if h.max == other.min {
		log("    new stretches h right", debug)
		return true, []horizontal{{row: h.row, min: h.min, max: other.max}}
	}

	if h.max == other.max {
		log("    new shrinks h from the right", debug)
		return true, []horizontal{{row: h.row, min: h.min, max: other.min}}
	}

	log("    new splits h in two", debug)
	left := horizontal{row: h.row, min: h.min, max: other.min - 1}
	right := horizontal{row: h.row, min: other.max + 1, max: h.max}
	return true, []horizontal{left, right}
}

func (h horizontal) overlaps(other horizontal) bool {
	return other.min <= h.max && other.max >= h.min
}

func (h horizontal) contains(col int) bool {
	return col >= h.min && col <= h.max
}

func (h horizontal) toStringWithRow() string {
	return fmt.Sprintf("row: %d, %d to %d", h.row, h.min, h.max)
}

func (h horizontal) toStringNoRow() string {
	return fmt.Sprintf("%d to %d", h.min, h.max)
}
