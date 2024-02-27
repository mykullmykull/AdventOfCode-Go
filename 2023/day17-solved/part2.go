package day17

func runB(grid []string) int {
	initialState := pathState{
		location:        point{row: 0, col: 0},
		facing:          dn,
		straightRepeats: 0,
	}
	initialState2 := pathState{
		location:        point{row: 0, col: 0},
		facing:          rt,
		straightRepeats: 0,
	}
	states := make(map[pathState]int)
	states[initialState] = 0
	states[initialState2] = 0
	queue := []queueItem{
		{initialState, 0},
		{initialState2, 0},
	}

	for {
		finalHeatLoss := getFinalHeatLoss(queue, grid)
		if finalHeatLoss != -1 {
			return finalHeatLoss
		}

		states, queue = calculateNextStatesForUltra(states, queue, grid)
	}
}

func calculateNextStatesForUltra(states map[pathState]int, queue []queueItem, grid []string) (map[pathState]int, []queueItem) {
	qItem := queue[0]
	newQueue := queue[1:]

	adjacentPoints := qItem.state.getAdjacents(grid)

	// // we no longer need the previous state if all of its adjacent points are updated
	for _, p := range adjacentPoints {
		if p.isValidUltraMove(qItem.state, grid) {
			straightRepeats := 1
			if qItem.state.facing == p.direction {
				straightRepeats = qItem.state.straightRepeats + 1
			}

			newState := pathState{
				location:        p.location,
				facing:          p.direction,
				straightRepeats: straightRepeats,
			}

			// we can skip this state if we've already been here
			prevHeatLoss := states[newState]
			if prevHeatLoss > 0 {
				continue
			}

			newHeatLoss := qItem.heatLoss + newState.location.getHeatLoss(grid)
			states[newState] = newHeatLoss
			newQueue = insertQueueItem(newQueue, queueItem{newState, newHeatLoss})
			continue
		}
	}

	return states, newQueue
}
