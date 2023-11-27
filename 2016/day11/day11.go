package day11

import (
	"fmt"
	"strings"
)

type Component struct {
	element     string
	isGenerator bool
	floor       int
}

type Snapshot struct {
	elevator   int
	components []Component
}

func runA(input []string) int {
	components := parse(input)
	elevator := 1
	possibleSnapshots := []Snapshot{{
		elevator:   elevator,
		components: components,
	}}
	steps := 0
	fmt.Printf("first step\n")
	printSnapshots(possibleSnapshots)
	for {
		possibleSnapshots = getPossibleSnapshots(possibleSnapshots)
		if isFinished(possibleSnapshots) {
			break
		}
		if len(possibleSnapshots) == 0 {
			fmt.Println("No valid snapshots")
			break
		}
		steps++
		fmt.Printf("steps: %d\n", steps)
		printSnapshots(possibleSnapshots)
	}

	return steps
}

func parse(input []string) []Component {
	components := []Component{}
	for i, line := range input[0:3] {
		componentsStr := strings.Split(line, " contains ")[1]
		componentsStr = strings.Replace(componentsStr, " and ", ", ", -1)
		componentsStr = strings.Replace(componentsStr, "-compatible ", " ", -1)
		componentsStr = strings.Replace(componentsStr, " a ", " ", -1)
		componentsStr = strings.Replace(componentsStr, "a ", " ", -1)
		componentsStr = strings.Trim(strings.Replace(componentsStr, ".", "", -1), " ")
		componentsSplit := strings.Split(componentsStr, ", ")
		for _, componentStr := range componentsSplit {
			split := strings.Split(componentStr, " ")
			components = append(components, Component{
				element:     split[0],
				isGenerator: split[1] == "generator",
				floor:       i + 1,
			})
		}
	}
	return components
}

func getPossibleSnapshots(possibleSnapshots []Snapshot) []Snapshot {
	newPossibleSnapshots := []Snapshot{}
	for _, snapshot := range possibleSnapshots {
		floor := snapshot.elevator
		possibleLoads := getPossibleLoads(snapshot.components, floor)
		fmt.Printf("possibleLoads: %v\n", possibleLoads)
		for _, load := range possibleLoads {
			newFloors := []int{floor - 1, floor + 1}
			for _, newFloor := range newFloors {
				if newFloor < 1 || newFloor > 4 {
					continue
				}
				components := updateComponents(snapshot.components, load, newFloor)
				newSnapshot := Snapshot{
					elevator:   newFloor,
					components: components,
				}
				if isValid(newSnapshot) {
					newPossibleSnapshots = append(newPossibleSnapshots, newSnapshot)
				}
			}
		}
	}
	return newPossibleSnapshots
}

func getPossibleLoads(components []Component, floor int) [][]Component {
	possibleLoads := [][]Component{}
	for _, component := range components {
		if component.floor == floor {
			possibleLoads = append(possibleLoads, []Component{component})
		}
	}
	for i, component1 := range components {
		for j, component2 := range components {
			if i == j {
				continue
			}
			if component1.floor == floor && component2.floor == floor {
				possibleLoads = append(possibleLoads, []Component{component1, component2})
			}
		}
	}
	return possibleLoads
}

func updateComponents(components []Component, load []Component, newFloor int) []Component {
	newComponents := []Component{}
	for _, component := range components {
		isInLoad := componentIsInLoad(component, load)
		if isInLoad {
			newComponents = append(newComponents, Component{
				element:     component.element,
				isGenerator: component.isGenerator,
				floor:       newFloor,
			})
		} else {
			newComponents = append(newComponents, component)
		}
	}
	return newComponents
}

func componentIsInLoad(component Component, load []Component) bool {
	for _, item := range load {
		if componentsAreEqual(component, item) {
			return true
		}
	}
	return false
}

func componentsAreEqual(c1 Component, c2 Component) bool {
	return c1.isGenerator == c2.isGenerator &&
		c1.element == c2.element &&
		c1.floor == c2.floor
}

func isValid(snapshot Snapshot) bool {
	for _, component := range snapshot.components {
		if isFried(component, snapshot.components) {
			return false
		}
	}
	return true
}

func isFried(component Component, components []Component) bool {
	return !component.isGenerator &&
		isExposed(component, components) &&
		!isPowered(component, components)
}

func isExposed(component Component, components []Component) bool {
	for _, component2 := range components {
		if component2.floor == component.floor && component2.isGenerator && component2.element != component.element {
			return true
		}
	}
	return false
}

func isPowered(component Component, components []Component) bool {
	for _, component2 := range components {
		if component2.floor == component.floor && component2.element == component.element && component2.isGenerator {
			return true
		}
	}
	return false
}

func isFinished(possibleSnapshots []Snapshot) bool {
	for _, snapshot := range possibleSnapshots {
		isFinished := true
		for _, component := range snapshot.components {
			if component.floor != 4 {
				isFinished = false
			}
		}
		if isFinished {
			return true
		}
	}
	return false
}

//-------------------- Print --------------------
func printSnapshots(snapshots []Snapshot) {
	for _, snapshot := range snapshots {
		printSnapshot(snapshot)
	}
}

func printSnapshot(snapshot Snapshot) {
	for floor := 4; floor > 0; floor-- {
		elevator := " "
		if snapshot.elevator == floor {
			elevator = "E"
		}
		components := []string{}
		for _, component := range snapshot.components {
			if component.floor == floor {
				components = append(components, printComponent(component))
			}
		}
		fmt.Printf("F%d: %s %v\n", floor, elevator, components)
	}
	fmt.Println()
}

func printComponent(component Component) string {
	t := "M"
	if component.isGenerator {
		t = "G"
	}
	return component.element[0:2] + t
}
