package day

import "math"

func main(input []string, row int) int {
	sensors := parse(input)
	minCol, maxCol := getMinMax(sensors)
	noBeaconCount := 0
	for col := minCol; col <= maxCol; col++ {
		p := point{col: col, row: row}
		for _, s := range sensors {
			if p.distanceFrom(s.closestBeacon) > 0 && p.distanceFrom(s.p) <= s.distance {
				noBeaconCount++
				break
			}
		}
	}
	return noBeaconCount
}

func getMinMax(sensors []sensor) (int, int) {
	min := math.MaxInt
	max := math.MinInt
	for _, s := range sensors {
		sMin := s.p.col - s.distance
		sMax := s.p.col + s.distance
		if sMin < min {
			min = sMin
		}
		if sMax > max {
			max = sMax
		}
	}
	return min, max
}
