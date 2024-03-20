package day20

func runB(input []string) int {
	// debug := true
	broadcastOuts, fModules, cModules := parseInput(input)

	// for rx to receive low all of kj's inputs [dr ln vn zx] must send high
	cycleMap := map[string]int{}
	cycleMap["dr"] = 0
	cycleMap["ln"] = 0
	cycleMap["vn"] = 0
	cycleMap["zx"] = 0

	buttonPresses := 0
	for {
		if isFinished(cycleMap) {
			break
		}
		buttonPresses++
		signals := broadcastInitialSignals(broadcastOuts)
		for {
			if len(signals) == 0 {
				break
			}
			nextSignal := signals[0]
			signals = signals[1:]

			if nextSignal.isHigh && (nextSignal.from == "dr" ||
				nextSignal.from == "ln" ||
				nextSignal.from == "vn" ||
				nextSignal.from == "zx") &&
				cycleMap[nextSignal.from] == 0 {
				cycleMap[nextSignal.from] = buttonPresses
			}

			newSignals := []signal{}
			if module, ok := fModules[nextSignal.to]; ok {
				fModules[nextSignal.to], newSignals = module.processSignal(nextSignal)
			} else if module, ok := cModules[nextSignal.to]; ok {
				cModules[nextSignal.to], newSignals = module.processSignal(nextSignal)
			}

			signals = append(signals, newSignals...)
		}
	}

	return getLCM(cycleMap)
}

func isFinished(m map[string]int) bool {
	for _, v := range m {
		if v == 0 {
			return false
		}
	}
	return true
}

func getLCM(m map[string]int) int {
	lcm := 1
	for _, v := range m {
		lcm = lcm * v / getGCD(lcm, v)
	}
	return lcm
}

func getGCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
