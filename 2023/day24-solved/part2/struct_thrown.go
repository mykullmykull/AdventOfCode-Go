package day24Part2

import (
	"fmt"
	"math"
)

type thrown struct {
	dx     float64
	dy     float64
	dz     float64
	dDiff  float64
	rangeX rng
	rangeY rng
	rangeZ rng
}

func (t thrown) narrowThrownRanges(s stone) thrown {
	t = t.narrowWithSlope(s, 'x')
	if !t.rangeX.isValid() {
		return t
	}

	t = t.narrowWithSlope(s, 'y')
	if !t.rangeY.isValid() {
		return t
	}

	t = t.narrowWithSlope(s, 'z')
	t = t.narrowWithTime(s)
	return t
}

func (t thrown) narrowWithSlope(s stone, d rune) thrown {
	newR := rng{
		min: math.MaxFloat64 * -1,
		max: math.MaxFloat64,
	}
	tr, td, ss, sd := t.getDimensionValues(s, d)
	if td == sd {
		newR.min = ss
		newR.max = ss
	}
	if td > sd {
		newR.max = ss
	}
	if td < sd {
		newR.min = ss
	}

	tr = tr.update(newR)
	t = t.assignNewRange(tr, d)
	return t
}

func (t thrown) getDimensionValues(s stone, dimension rune) (rng, float64, float64, float64) {
	switch dimension {
	case 'x':
		return t.rangeX, t.dx, s.x, s.dx
	case 'y':
		return t.rangeY, t.dy, s.y, s.dy
	case 'z':
		return t.rangeZ, t.dz, s.z, s.dz
	default:
		panic(fmt.Sprintf("Invalid dimension %v", dimension))
	}
}

func (t thrown) assignNewRange(r rng, d rune) thrown {
	switch d {
	case 'x':
		t.rangeX = r
	case 'y':
		t.rangeY = r
	case 'z':
		t.rangeZ = r
	default:
		panic(fmt.Sprintf("Invalid dimension %v", d))
	}
	return t
}

// all 3 dimensions must have the same time range
func (t thrown) narrowWithTime(s stone) thrown {
	invalidRange := rng{max: -1.0}
	invalidT := t
	invalidT.rangeX = invalidRange

	timeRange := rng{
		min: 0,
		max: math.MaxFloat64,
	}

	timeRange = t.updateTimeRange(s, timeRange, 'x')
	if !timeRange.isValid() {
		return invalidT
	}

	timeRange = t.updateTimeRange(s, timeRange, 'y')
	if !timeRange.isValid() {
		return invalidT
	}

	timeRange = t.updateTimeRange(s, timeRange, 'z')
	if !timeRange.isValid() {
		return invalidT
	}

	t = t.narrowWithTimeAndDimension(s, timeRange, 'x')
	if !t.rangeX.isValid() {
		return invalidT
	}

	t = t.narrowWithTimeAndDimension(s, timeRange, 'y')
	if !t.rangeY.isValid() {
		return invalidT
	}

	t = t.narrowWithTimeAndDimension(s, timeRange, 'z')
	if !t.rangeZ.isValid() {
		return invalidT
	}

	return t
}

func (t thrown) updateTimeRange(s stone, timeRange rng, d rune) rng {
	tr, td, ss, sd := t.getDimensionValues(s, d)
	// ss + sd * time = ts + td * time
	// sd * time - td * time = ts - ss
	// time * (sd - td) = ts - ss
	// time = (ts - ss) / (sd - td)
	if sd == td {
		return timeRange
	}

	time1 := (tr.min - ss) / (sd - td)
	time2 := (tr.max - ss) / (sd - td)
	r := rng{}.new(time1, time2)
	return timeRange.update(r)
}

func (t thrown) narrowWithTimeAndDimension(s stone, timeRange rng, d rune) thrown {
	tr, td, ss, sd := t.getDimensionValues(s, d)
	// ss + sd * time = ts + td * time
	// ts = ss + sd * time - td * time
	v1 := ss + sd*timeRange.min - td*timeRange.min
	v2 := ss + sd*timeRange.max - td*timeRange.max
	r := rng{}.new(v1, v2)
	tr = tr.update(r)
	t = t.assignNewRange(tr, d)
	return t
}

func (t thrown) getNextDs() thrown {
	if t.dx == 0 && t.dy == 0 && t.dz == 0 {
		t.dDiff++
		t.dx = t.dDiff
		return t
	}

	switch t.dDiff {
	case math.Abs(t.dx):
		t.dx, t.dy, t.dz = rollNext(t.dx, getNextD(t.dy), t.dz, t.dDiff)
	case math.Abs(t.dy):
		t.dy, t.dz, t.dx = rollNext(t.dy, getNextD(t.dx), t.dz, t.dDiff)
	case math.Abs(t.dz):
		t.dz, t.dx, t.dy = rollNext(t.dz, getNextD(t.dx), t.dy, t.dDiff)
		if t.dz == 0 {
			t.dDiff++
			t.dx = t.dDiff
			return t
		}
	default:
		panic(fmt.Sprintf("none of the ds (%f, %f, %f) match t.dDiff: %v\n", t.dx, t.dy, t.dz, t.dDiff))
	}

	return t
}

func rollNext(focus float64, moved float64, rollover float64, max float64) (float64, float64, float64) {
	if moved < max {
		return focus, moved, rollover
	}
	moved = 0

	rollover = getNextD(rollover)
	if rollover < max {
		return focus, moved, rollover
	}
	rollover = 0

	focus = getNextD(focus)
	if focus <= max {
		return focus, moved, rollover
	}

	focus = 0
	moved = max
	return focus, moved, rollover
}

func getNextD(d float64) float64 {
	d = d * -1
	if d >= 0 {
		d++
	}
	return d
}

func (t thrown) resetRanges() thrown {
	t.rangeX = rng{
		min: math.MaxFloat64 * -1,
		max: math.MaxFloat64,
	}
	t.rangeY = rng{
		min: math.MaxFloat64 * -1,
		max: math.MaxFloat64,
	}
	t.rangeZ = rng{
		min: math.MaxFloat64 * -1,
		max: math.MaxFloat64,
	}
	return t
}

func (t thrown) toString() string {
	str := fmt.Sprintf("dx: %d, dy: %d, dz: %d", int(t.dx), int(t.dy), int(t.dz))
	str += ", rangeX: " + t.rangeX.toString()
	str += ", rangeY: " + t.rangeY.toString()
	str += ", rangeZ: " + t.rangeZ.toString()
	return str
}

func (t thrown) slopesToString() string {
	dxSign := "+"
	if t.dx < 0 {
		dxSign = ""
	}

	dySign := "+"
	if t.dy < 0 {
		dySign = ""
	}

	dzSign := "+"
	if t.dz < 0 {
		dzSign = ""
	}

	return fmt.Sprintf("dx: %s%d, dy: %s%d, dz: %s%d, dDiff: %d", dxSign, int(t.dx), dySign, int(t.dy), dzSign, int(t.dz), int(t.dDiff))
}

func (t thrown) rangesToString() string {
	str := "x: " + t.rangeX.toString()
	str += ", y: " + t.rangeY.toString()
	str += ", z: " + t.rangeZ.toString()
	return str
}
