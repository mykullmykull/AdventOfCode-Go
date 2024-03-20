package day20

import (
	"fmt"
	"strings"
)

func parseInput(input []string) ([]string, map[string]flipFlopModule, map[string]conjunctionModule) {
	broadcastOuts := []string{}
	fModules := map[string]flipFlopModule{}
	cModules := map[string]conjunctionModule{}
	for _, line := range input {
		if strings.Contains(line, "broadcaster") {
			broadcastOuts = parseBroadcaster(line)
			continue
		}

		if strings.Contains(line, "%") {
			name, module := parseFlipFlopModule(line)
			fModules[name] = module
			continue
		}
		if strings.Contains(line, "&") {
			name, module := parseConjunctionModule(line)
			cModules[name] = module
			continue
		}
		panic(fmt.Sprintf("unknown module type: %s", line))
	}
	cModules = populateCModuleIns(broadcastOuts, fModules, cModules)

	return broadcastOuts, fModules, cModules
}

func parseBroadcaster(line string) []string {
	return strings.Split(strings.Split(line, " -> ")[1], ", ")
}

func parseFlipFlopModule(line string) (string, flipFlopModule) {
	name, outs := parseModule(line)
	return name, flipFlopModule{outs: outs}
}

func parseConjunctionModule(line string) (string, conjunctionModule) {
	name, outs := parseModule(line)
	return name, conjunctionModule{lastSignalWasHigh: map[string]bool{}, outs: outs}
}

func parseModule(line string) (string, []string) {
	split := strings.Split(line, " -> ")
	name := split[0][1:]
	outs := strings.Split(split[1], ", ")
	return name, outs
}

func populateCModuleIns(
	broadcastOuts []string,
	fModules map[string]flipFlopModule,
	cModules map[string]conjunctionModule,
) map[string]conjunctionModule {
	for cName, _ := range cModules {
		for _, out := range broadcastOuts {
			if out == cName {
				cModules[cName].lastSignalWasHigh["broadcaster"] = false
			}
		}
		for fName, fModule := range fModules {
			for _, out := range fModule.outs {
				if out == cName {
					cModules[cName].lastSignalWasHigh[fName] = false
				}
			}
		}
		for cName2, cModule := range cModules {
			for _, out := range cModule.outs {
				if out == cName {
					cModules[cName].lastSignalWasHigh[cName2] = false
				}
			}
		}
	}
	return cModules
}
