package day20

import "fmt"

func printModules(fModules map[string]flipFlopModule, cModules map[string]conjunctionModule, count int) {
	fmt.Println()
	for c1 := 'a'; c1 <= 'z'; c1++ {
		for c2 := 'a'; c2 <= 'z'; c2++ {
			if module, ok := fModules[string(c1)+string(c2)]; ok {
				fmt.Printf("\n%s%s: %v ", string(c1), string(c2), module)
			}
		}
	}
	fmt.Println()
	for c1 := 'a'; c1 <= 'z'; c1++ {
		for c2 := 'a'; c2 <= 'z'; c2++ {
			if module, ok := cModules[string(c1)+string(c2)]; ok {
				fmt.Printf("\n%s%s ins: ", string(c1), string(c2))
				for k, _ := range module.lastSignalWasHigh {
					fmt.Printf("%s ", k)
				}
				fmt.Printf(", outs: %v", module.outs)
			}
		}
	}
}

func printOns(fModules map[string]flipFlopModule, cModules map[string]conjunctionModule, count int) {
	fmt.Println()
	for c1 := 'a'; c1 <= 'z'; c1++ {
		for c2 := 'a'; c2 <= 'z'; c2++ {
			if module, ok := fModules[string(c1)+string(c2)]; ok {
				fmt.Printf("%s", printBool(module.isOn))
			}
		}
	}
	fmt.Print("  |  ")
	for c1 := 'a'; c1 <= 'z'; c1++ {
		for c2 := 'a'; c2 <= 'z'; c2++ {
			if module, ok := cModules[string(c1)+string(c2)]; ok {
				fmt.Printf("%s", printAllHigh(module.lastSignalWasHigh))
			}
		}
	}
	fmt.Printf("  |  count: %d", count)
}

func printBool(b bool) string {
	if b {
		return "#"
	}
	return "."
}

func printAllHigh(m map[string]bool) string {
	// if any low then send high
	for _, v := range m {
		if !v {
			return printBool(true)
		}
	}
	// if all high then send low
	return printBool(false)
}

func copyFModules(m map[string]flipFlopModule) map[string]flipFlopModule {
	newMap := map[string]flipFlopModule{}
	for k, v := range m {
		newMap[k] = v
	}
	return newMap
}

func copyCModules(m map[string]conjunctionModule) map[string]conjunctionModule {
	newMap := map[string]conjunctionModule{}
	for k, v := range m {
		newMap[k] = v
	}
	return newMap
}

func printChanges(
	oldFModules map[string]flipFlopModule,
	newFModules map[string]flipFlopModule,
	oldCModules map[string]conjunctionModule,
	newCModules map[string]conjunctionModule,
	buttonPresses int,
) {
	fmt.Println()
	fmt.Printf("p: %d, ", buttonPresses)
	fmt.Print("f: ")
	for c1 := 'a'; c1 <= 'z'; c1++ {
		for c2 := 'a'; c2 <= 'z'; c2++ {
			k := string(c1) + string(c2)
			v := newFModules[k]
			if v.isOn != oldFModules[k].isOn {
				fmt.Printf("%s: %v -> %v, ", k, printBool(oldFModules[k].isOn), printBool(v.isOn))
			}
		}
	}
	fmt.Print("c: ")
	for c1 := 'a'; c1 <= 'z'; c1++ {
		for c2 := 'a'; c2 <= 'z'; c2++ {
			k := string(c1) + string(c2)
			v := newCModules[k]
			if fmt.Sprintf("%v", v) != fmt.Sprintf("%v", oldCModules[k]) {
				fmt.Printf("%s: %v -> %v, ", k, oldCModules[k], v)
			}
		}
	}
}
