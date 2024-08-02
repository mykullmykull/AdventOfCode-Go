package day25

import (
	"fmt"
	"math/rand"
	"time"
)

func runA(input []string) int {
	var lastNodes map[string]node
	lt := ""
	rt := ""
	for {
		fmt.Println()
		fmt.Printf("wires cut: ")
		for _, n := range lastNodes {
			for _, v := range n.connections {
				fmt.Printf("%d ", v)
			}
			break
		}
		fmt.Println()

		lt = ""
		rt = ""
		for k := range lastNodes {
			if lt == "" {
				lt = k
				continue
			}
			if rt == "" {
				rt = k
				break
			}
		}

		if lt != "" && rt != "" {
			if len(lastNodes) != 2 {
				panic(fmt.Sprintf("expected 2 nodes after collapse but found %d", len(lastNodes)))
			}
			if len(lastNodes[lt].connections) != 1 || len(lastNodes[rt].connections) != 1 {
				panic(fmt.Sprintf("expected 1 connection each for %v", lastNodes))
			}

			if lastNodes[lt].connections[rt] == 3 && lastNodes[rt].connections[lt] == 3 {
				break
			}
		}

		nodes := parseInput(input)
		lastNodes = collapseNodesToTwo(nodes)
	}
	return (len(lastNodes[lt].name) + 1) / 4 * (len(lastNodes[rt].name) + 1) / 4
}

func collapseNodesToTwo(nodes map[string]node) map[string]node {
	for {
		fmt.Printf("  nodes: %d\n", len(nodes))
		if len(nodes) == 2 {
			break
		}

		keys := make([]string, len(nodes))
		i := 0
		for k := range nodes {
			keys[i] = k
			i++
		}

		rand.Seed(time.Now().UnixNano())
		ltIndex := rand.Intn(len(keys) - 1)
		lt := keys[ltIndex]

		keys = make([]string, len(nodes[lt].connections))
		i = 0
		for k := range nodes[lt].connections {
			keys[i] = k
			i++
		}

		rt := keys[0]
		if len(keys) > 1 {
			rand.Seed(time.Now().UnixNano())
			rtIndex := rand.Intn(len(keys) - 1)
			rt = keys[rtIndex]
		}

		ltNode := nodes[lt]
		nodes = ltNode.combineNodes(rt, nodes)
	}

	return nodes
}
