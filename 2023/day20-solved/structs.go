package day20

import "fmt"

type strength string

type signal struct {
	from   string
	to     string
	isHigh bool
}

type flipFlopModule struct {
	isOn bool
	outs []string
}

type conjunctionModule struct {
	lastSignalWasHigh map[string]bool
	outs              []string
}

func (f flipFlopModule) processSignal(s signal) (flipFlopModule, []signal) {
	if s.isHigh {
		return f, []signal{}
	}
	f.isOn = !f.isOn
	isOutGoingHigh := false
	if f.isOn {
		isOutGoingHigh = true
	}
	return f, generateSignals(s.to, f.outs, isOutGoingHigh)
}

func (c conjunctionModule) processSignal(s signal) (conjunctionModule, []signal) {
	c.lastSignalWasHigh[s.from] = s.isHigh
	isOutGoingHigh := false
	for _, wasHigh := range c.lastSignalWasHigh {
		if !wasHigh {
			isOutGoingHigh = true
			break
		}
	}
	return c, generateSignals(s.to, c.outs, isOutGoingHigh)
}

func generateSignals(from string, outs []string, isHigh bool) []signal {
	var signals []signal
	for _, out := range outs {
		signals = append(signals, signal{
			from:   from,
			to:     out,
			isHigh: isHigh,
		})
	}
	return signals
}

func (s signal) toString() string {
	strength := "low"
	if s.isHigh {
		strength = "high"
	}
	return fmt.Sprintf("%s -%s-> %s", s.from, strength, s.to)
}
