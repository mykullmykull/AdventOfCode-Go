package day24Part2

import (
	"fmt"
	"math"
)

type rng struct {
	min float64
	max float64
}

func (r rng) isValid() bool {
	return r.min <= r.max
}

func (r rng) update(newRange rng) rng {
	if newRange.min > r.min {
		r.min = newRange.min
	}
	if newRange.max < r.max {
		r.max = newRange.max
	}
	return r
}

func (r rng) new(val1 float64, val2 float64) rng {
	if val1 < val2 {
		r.min = val1
		r.max = val2
	} else {
		r.min = val2
		r.max = val1
	}
	return r
}

func (r rng) isFinished() bool {
	return r.min == r.max
}

func (r rng) toString() string {
	limit := 10.0 * 100
	minStr := "-∞"
	maxStr := "∞"

	if math.Abs(r.min) > limit {
		minStr = fmt.Sprintf("%v", r.min)
	}

	if math.Abs(r.max) > limit {
		maxStr = fmt.Sprintf("%v", r.max)
	}

	return fmt.Sprintf("[%s, %s]", minStr, maxStr)
}
