package day20

import "fmt"

func runA(input []string) int {
	broadcastOuts, fModules, cModules := parseInput(input)
	lowCount := 0
	highCount := 0
	for i := 0; i < 1000; i++ {
		lowCount++ // one for the button press to trigger broadcast
		signals := broadcastInitialSignals(broadcastOuts)
		for {
			if len(signals) == 0 {
				break
			}
			nextSignal := signals[0]
			signals = signals[1:]
			if nextSignal.isHigh {
				highCount++
			} else {
				lowCount++
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
	return lowCount * highCount
}

func broadcastInitialSignals(broadcastOuts []string) []signal {
	signals := make([]signal, len(broadcastOuts))
	for i, out := range broadcastOuts {
		signals[i] = signal{
			from: "broadcaster",
			to:   out,
		}
	}
	return signals
}

func log(s string, debug bool) {
	if debug {
		fmt.Println(s)
	}
}
