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
