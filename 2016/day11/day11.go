package day11FirstTry

import (
	"fmt"
	"strings"
)

type Snapshot struct {
	elevator int
	floors   [4]string
}

func runA(input []string) int {
	possibleSnapshots := []Snapshot{parse(input)}
	printSnapshot(possibleSnapshots[0])
	steps := 0
	previousSnapshots := []Snapshot{}
	for {
		for _, s := range possibleSnapshots {
			if isSolution(s) {
				break
			}
		}

		steps++
		possibleSnapshots = getPossibleSnapshots(possibleSnapshots, previousSnapshots)
		previousSnapshots = append(previousSnapshots, possibleSnapshots...)
	}

	return steps
}

func parse(input []string) Snapshot {
	s := Snapshot{elevator: 0}
	for i, line := range input[0:3] {
		floorStr := ""
		componentsStr := strings.Split(line, " contains ")[1]
		componentsStr = strings.Replace(componentsStr, ", and ", ", ", -1)
		componentsStr = strings.Replace(componentsStr, " and ", ", ", -1)
		componentsStr = strings.Replace(componentsStr, "-compatible ", " ", -1)
		componentsStr = strings.Replace(componentsStr, " a ", " ", -1)
		componentsStr = strings.Replace(componentsStr, "a ", " ", -1)
		componentsStr = strings.Trim(strings.Replace(componentsStr, ".", "", -1), " ")
		componentsSplit := strings.Split(componentsStr, ", ")

		for _, componentStr := range componentsSplit {
			split := strings.Split(componentStr, " ")
			floorStr += split[0][0:2]
			if split[1] == "generator" {
				floorStr += "G "
			} else {
				floorStr += "M "
			}
		}
		s.floors[i] = strings.Trim(floorStr, " ")
	}
	return s
}

func printSnapshot(s Snapshot) {
	for i, floor := range s.floors {
		elevatorStr := " "
		if s.elevator == i {
			elevatorStr = "E"
		}
		fmt.Printf("F%d: %s %s\n", i+1, elevatorStr, floor)
	}
	fmt.Println()
}

func isSolution(s Snapshot) bool {
	return s.floors[0] == "" && s.floors[1] == "" && s.floors[2] == ""
}

func getPossibleSnapshots(possibleSnapshots []Snapshot, previousSnapshots []Snapshot) []Snapshot {
	var newSnapshots []Snapshot
	for _, s := range possibleSnapshots {
		newSnapshots = append(newSnapshots, getPossibleSnapshotsFromSnapshot(s, previousSnapshots)...)
	}
	return newSnapshots
}

func getPossibleSnapshotsFromSnapshot(s Snapshot, previousSnapshots []Snapshot) []Snapshot {
	var newSnapshots []Snapshot
	possibleLoads := getPossibleLoads(s)
	for _, load := range possibleLoads {
		for _, direction := range []int{-1, 1} {
			if s.elevator+direction < 0 || s.elevator+direction > 3 {
				continue
			}

			newSnapshot := getNewSnapshot(s, load, direction)
			if isValid(newSnapshot, previousSnapshots) {
				newSnapshots = append(newSnapshots, newSnapshot)
			}
		}
	}

	return newSnapshots
}

func getPossibleLoads(s Snapshot) []string {
	components := strings.Split(
		strings.Split(
			s.floors[s.elevator],
			" E ",
		)[1],
		" ",
	)

	possibleLoads := components
	for i, component := range components {
		for j, component2 := range components {
			if i >= j {
				continue
			}

			possibleLoads = append(possibleLoads, component+" "+component2)
		}
	}
	return possibleLoads
}
