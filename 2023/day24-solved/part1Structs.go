package day24

import "fmt"

type point struct {
	x float64
	y float64
	z float64
}

type stn struct {
	start point
	dx    float64
	dy    float64
	dz    float64
}

type combo struct {
	s1 stn
	s2 stn
}

type alwaysFn func() bool
type timeFn func() float64

func (s stn) advanceX(velocityMax float64) stn {
	s.dx, s.start.x = negateOrAdd(s.dx, s.start.x, velocityMax)
	return s
}

func (s stn) resetYZ() stn {
	// continue looking for a new x and dx
	s.start.y = 0
	s.start.z = 0
	s.dy = 0
	s.dz = 0
	return s
}

func (s stn) getNextStone(velocityMax float64) stn {
	s.dz += 1
	s.dz, s.dy = maybeCarryThe1(s.dz, s.dy, velocityMax)
	s.dy, s.dx = maybeCarryThe1(s.dy, s.dx, velocityMax)
	s.dx, s.start.z = maybeCarryThe1(s.dx, s.start.z, velocityMax)
	return s
}

func negateOrAdd(a float64, b float64, max float64) (float64, float64) {
	a = a * -1
	if a >= 0 {
		a, b = maybeCarryThe1(a+1, b, max)
	}
	return a, b
}

func maybeCarryThe1(a float64, b float64, max float64) (float64, float64) {
	if a > max {
		a = 0
		b += 1
	}
	return a, b
}

func (c combo) timeXCrosses() float64 {
	t := (c.s1.start.x - c.s2.start.x) / (c.s2.dx - c.s1.dx)
	// fmt.Printf("    t: %f\n", t)
	return t
}

func (c combo) timeYCrosses() float64 {
	t := (c.s1.start.y - c.s2.start.y) / (c.s2.dy - c.s1.dy)
	return t
}

func (c combo) timeZCrosses() float64 {
	t := (c.s1.start.z - c.s2.start.z) / (c.s2.dz - c.s1.dz)
	return t
}

func (c combo) xCrosses() bool {
	return c.xAlwaysMatches() || c.timeXCrosses() > 0
}

func (c combo) crosses(always1 alwaysFn, always2 alwaysFn, time1 timeFn, time2 timeFn) bool {
	if always1() || always2() {
		return true
	}
	t1 := time1()
	t2 := time2()
	return t1 > 0 && floatsEqual(t1, t2)
}

func (c combo) xyCrosses() bool {
	return c.crosses(c.xAlwaysMatches, c.yAlwaysMatches, c.timeXCrosses, c.timeYCrosses)
}

func (c combo) xzCrosses() bool {
	return c.crosses(c.xAlwaysMatches, c.zAlwaysMatches, c.timeXCrosses, c.timeZCrosses)
}

func (c combo) xyzCrosses() bool {
	return c.xyCrosses() && c.xzCrosses()
}

func floatsEqual(a float64, b float64) bool {
	return fmt.Sprintf("%.4f", a) == fmt.Sprintf("%.4f", b)
}

func (c combo) xAlwaysMatches() bool {
	return c.s1.start.x == c.s2.start.x && c.s1.dx == c.s2.dx
}

func (c combo) yAlwaysMatches() bool {
	return c.s1.start.y == c.s2.start.y && c.s1.dy == c.s2.dy
}

func (c combo) zAlwaysMatches() bool {
	return c.s1.start.z == c.s2.start.z && c.s1.dz == c.s2.dz
}

// --------------------------- PART 1 -------------------------- //
func (s stn) yxSlope() float64 {
	return float64(s.dy) / float64(s.dx)
}

func (s stn) zxSlope() float64 {
	return float64(s.dz) / float64(s.dx)
}

func (s stn) zySlope() float64 {
	return float64(s.dz) / float64(s.dy)
}

func (s stn) yInterceptWithX() float64 {
	return s.start.y - s.yxSlope()*s.start.x
}

func (s stn) zInterceptWithX() float64 {
	return s.start.z - s.zxSlope()*s.start.x
}

func (s stn) zInterceptWithY() float64 {
	return s.start.z - s.zySlope()*s.start.y
}

func (c combo) yxWillCrossPaths(min float64, max float64) bool {
	// s1y := c.s1.slope() * x + c.s1.yIntercept()
	// s2y := c.s2.slope() * x + c.s2.yIntercept()
	// s1y = s2y, solve for x
	x := (c.s2.yInterceptWithX() - c.s1.yInterceptWithX()) / (c.s1.yxSlope() - c.s2.yxSlope())
	t1 := (x - c.s1.start.x) / c.s1.dx
	t2 := (x - c.s2.start.x) / c.s2.dx
	y := c.s1.start.y + (t1 * c.s1.dy)
	isValid := true &&
		t1 > 0 &&
		t2 > 0 &&
		x >= min &&
		x <= max &&
		y >= min &&
		y <= max
	return isValid
}
