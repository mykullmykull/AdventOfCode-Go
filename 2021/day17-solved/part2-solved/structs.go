package day

import "fmt"

type probe struct {
	minX, maxX int
	minY, maxY int
	velX, velY int
}

func newProbe(input string) probe {
	var minX, maxX, minY, maxY int
	fmt.Sscanf(input, "target area: x=%d..%d, y=%d..%d", &minX, &maxX, &minY, &maxY)
	return probe{minX: minX, maxX: maxX, minY: minY, maxY: maxY}
}

func (p *probe) getMinVx() {
	// For minimum velocity: we need to find the smallest vx such that vx(vx+1)/2 >= minX
	// This is the sum of arithmetic series 1+2+3+...+vx which gives us the furthest x position
	// when the velocity eventually becomes 0
	p.velX = 0
	for p.velX*(p.velX+1)/2 < p.minX {
		p.velX++
	}
}

func (p *probe) getMaxVy() {
	// When the probe returns to y=0 after going up, its velocity will be -velY-1
	// The max y velocity occurs when the probe hits the bottom of the target area
	// in the first step after crossing y=0
	// In this case, vy should be such that |minY| - 1 = vy
	p.velY = -p.minY - 1 // -minY gives the absolute value of minY (which is negative)
}

func (p *probe) getMaxY() int {
	// The maximum height is achieved when the vertical velocity reaches zero
	// This happens after exactly velY steps from launch
	// The maximum height equals the sum of the series velY + (velY-1) + (velY-2) + ... + 1
	// Which is velY*(velY+1)/2
	return p.velY * (p.velY + 1) / 2
}

func (p *probe) getMinVy() {
	// The minimum y velocity is simply the minimum y coordinate of the target area
	// Since we want to reach the bottom of the target area, we need to be able to reach minY
	// from y=0, which means we need a velocity that can reach -minY
	p.velY = p.minY
}

func (p *probe) countSuccessfulLaunches() int {
	count := 0
	minVelX := p.velX
	maxVelX := p.maxX // Maximum X velocity (any higher would overshoot in first step)
	minVelY := p.velY
	maxVelY := -p.minY - 1 // From part 1, this is the maximum possible y velocity

	for p.velX = minVelX; p.velX <= maxVelX; p.velX++ {
		for p.velY = minVelY; p.velY <= maxVelY; p.velY++ {
			if p.hitsTarget() {
				count++
			}
		}
	}
	return count
}

func (p *probe) hitsTarget() bool {
	x, y := 0, 0
	vx, vy := p.velX, p.velY

	// Maximum steps to simulate - prevents infinite loops
	// The probe will either hit or miss the target within a reasonable number of steps
	maxSteps := 1000

	for step := 0; step < maxSteps; step++ {
		x += vx
		y += vy
		if vx > 0 {
			vx--
		}
		vy--

		// Check if we hit the target
		if x >= p.minX && x <= p.maxX && y >= p.minY && y <= p.maxY {
			return true
		}

		// Early termination conditions
		// 1. Overshot horizontally
		// 2. Overshot vertically and still going down
		// 3. Can't reach target horizontally (vx=0 and still left of target)
		if x > p.maxX || (y < p.minY && vy <= 0) || (vx == 0 && x < p.minX) {
			return false
		}
	}

	return false
}
