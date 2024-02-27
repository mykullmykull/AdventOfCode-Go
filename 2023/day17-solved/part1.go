package day17

func runA(grid []string) int {
	initialState := pathState{
		location:        point{row: 0, col: 0},
		facing:          dn,
		straightRepeats: 0,
	}
	states := make(map[pathState]int)
	states[initialState] = 0
	queue := []queueItem{{initialState, 0}}

	for {
		finalHeatLoss := getFinalHeatLoss(queue, grid)
		if finalHeatLoss != -1 {
			return finalHeatLoss
		}

		states, queue = calculateNextStates(states, queue, grid)
	}
}

func getFinalHeatLoss(queue []queueItem, grid []string) int {
	for _, q := range queue {
		if q.state.location.atFinishLine(grid) {
			return q.heatLoss
		}
	}
	return -1
}

func calculateNextStates(states map[pathState]int, queue []queueItem, grid []string) (map[pathState]int, []queueItem) {
	qItem := queue[0]
	newQueue := queue[1:]

	adjacentPoints := qItem.state.getAdjacents(grid)

	// // we no longer need the previous state if all of its adjacent points are updated
	for _, p := range adjacentPoints {
		if p.isValidMove(qItem.state) {
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

func insertQueueItem(queue []queueItem, item queueItem) []queueItem {
	for i, q := range queue {
		if item.heatLoss < q.heatLoss {
			return append(queue[:i], append([]queueItem{item}, queue[i:]...)...)
		}
	}
	return append(queue, item)
}

func log(msg string, debug bool) {
	if debug {
		println(msg)
	}
}
