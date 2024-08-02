package day25

import (
	"strings"
)

func parseInput(input []string) map[string]node {
	nodes := map[string]node{}

	for _, line := range input {
		split := strings.Split(line, ":")
		leftName := split[0]
		split2 := strings.Split(split[1], " ")
		for _, rightName := range split2 {
			if rightName != "" {
				nodes[rightName] = getOrUpdateNode(nodes, rightName, leftName)
				nodes[leftName] = getOrUpdateNode(nodes, leftName, rightName)
			}
		}
	}
	return nodes
}

func getOrUpdateNode(nodes map[string]node, name string, toConnect string) node {
	n, ok := nodes[name]
	if !ok {
		n = node{name: name, connections: map[string]int{}}
	}
	n.connections[toConnect] = n.connections[toConnect] + 1
	return n
}
