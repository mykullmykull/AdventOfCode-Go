package day

import (
	"goAoc2022/utils"
	"strings"
)

func parse(input []string) []sensor {
	sensors := make([]sensor, len(input))
	for x, line := range input {
		sensorVsBeacon := strings.Split(line, ": closest beacon is at x=")

		sensorXAndY := strings.Split(sensorVsBeacon[0], "Sensor at x=")[1]
		sensorXVsY := strings.Split(sensorXAndY, ", y=")
		sensorX := utils.StrToInt(sensorXVsY[0])
		sensorY := utils.StrToInt(sensorXVsY[1])

		beaconXVsY := strings.Split(sensorVsBeacon[1], ", y=")
		beaconX := utils.StrToInt(beaconXVsY[0])
		beaconY := utils.StrToInt(beaconXVsY[1])

		p := point{col: sensorX, row: sensorY}
		b := point{col: beaconX, row: beaconY}
		s := sensor{p: p, closestBeacon: b}
		s.distance = p.distanceFrom(b)

		sensors[x] = s
	}
	return sensors
}
