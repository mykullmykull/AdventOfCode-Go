package day21

func runB(input []string, stepsLeft int) int {
	g := parse(input, stepsLeft)

	if stepsLeft < g.subFarmSize {
		startingSpots := []spot{g.start}
		e := expander{rows: input, up: true, dn: true, lt: true, rt: true}
		return g.spotsFromEdgeSubFarm(stepsLeft+1, startingSpots, e)
	}

	totalSpots := 0
	fullSubFarmsAtCurrentDistance := 1
	edgeSubFarmsAtCurrentDistance := -1
	currentFullSpots := 0

	for stepsLeft > g.subFarmSize {
		currentFullSpots = g.swapFullSpots(currentFullSpots)

		// add spots from fully reachable subFarms
		totalSpots += currentFullSpots * fullSubFarmsAtCurrentDistance

		// update counters
		fullSubFarmsAtCurrentDistance -= fullSubFarmsAtCurrentDistance % 4
		fullSubFarmsAtCurrentDistance += 4
		edgeSubFarmsAtCurrentDistance++
		if stepsLeft > g.subFarmSize*2 {
			// we still have full subFarms to count
			stepsLeft -= g.subFarmSize
			continue
		}

		// only corners and edges left
		g.edgesOddity = g.getEdgesOddityFromCurrentFullSpots(currentFullSpots)
		middle := g.subFarmSize / 2
		stepsLeft -= middle
		end := g.subFarmSize - 1

		// spots
		topMiddle := spot{r: 0, c: middle}
		rightMiddle := spot{r: middle, c: end}
		bottomMiddle := spot{r: end, c: middle}
		leftMiddle := spot{r: middle, c: 0}

		// add spots from corner subFarms
		topCorner := edgeSF{
			startingSpots: []spot{bottomMiddle},
			e:             expander{rows: input, up: true, dn: false, lt: true, rt: false},
		}
		rightCorner := edgeSF{
			startingSpots: []spot{leftMiddle},
			e:             expander{rows: input, up: true, dn: false, lt: false, rt: true},
		}
		bottomCorner := edgeSF{
			startingSpots: []spot{topMiddle},
			e:             expander{rows: input, up: false, dn: true, lt: false, rt: true},
		}
		leftCorner := edgeSF{
			startingSpots: []spot{rightMiddle},
			e:             expander{rows: input, up: false, dn: true, lt: true, rt: false},
		}
		for _, edge := range []edgeSF{topCorner, rightCorner, bottomCorner, leftCorner} {
			toAdd := g.spotsFromEdgeSubFarm(stepsLeft, edge.startingSpots, edge.e)
			totalSpots += toAdd
		}

		if edgeSubFarmsAtCurrentDistance == 0 {
			break
		}

		// add spots from edge subFarms
		topRightEdge := edgeSF{
			startingSpots: []spot{leftMiddle, bottomMiddle},
			e:             expander{rows: input, up: true, dn: false, lt: false, rt: false},
		}
		bottomRightEdge := edgeSF{
			startingSpots: []spot{leftMiddle, topMiddle},
			e:             expander{rows: input, up: false, dn: false, lt: false, rt: true},
		}
		bottomLeftEdge := edgeSF{
			startingSpots: []spot{rightMiddle, topMiddle},
			e:             expander{rows: input, up: false, dn: true, lt: false, rt: false},
		}
		topLeftEdge := edgeSF{
			startingSpots: []spot{rightMiddle, bottomMiddle},
			e:             expander{rows: input, up: false, dn: false, lt: true, rt: false},
		}
		for _, edge := range []edgeSF{topRightEdge, bottomRightEdge, bottomLeftEdge, topLeftEdge} {
			toAdd := g.spotsFromEdgeSubFarm(stepsLeft, edge.startingSpots, edge.e)
			totalSpots += toAdd * edgeSubFarmsAtCurrentDistance
		}
		break
	}

	return totalSpots
}
