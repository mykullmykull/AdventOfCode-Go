package day25

import "fmt"

type node struct {
	name        string
	connections map[string]int
}

func (n node) combineNodes(rt string, nodes map[string]node) map[string]node {
	combinedName := fmt.Sprintf("%s-%s", n.name, rt)
	newNode := node{
		name:        combinedName,
		connections: n.combineConnections(n.connections, nodes[rt]),
	}
	nodes[combinedName] = newNode

	// replace node connections to lt and rt
	for name := range nodes[combinedName].connections {
		for k := range nodes[name].connections {
			if k == n.name || k == rt {
				nodes[name].connections[combinedName] += nodes[name].connections[k]
				delete(nodes[name].connections, k)
			}
		}
	}

	// delete nodes for lt and rt
	delete(nodes, n.name)
	delete(nodes, rt)
	return nodes
}

func (n node) combineConnections(l map[string]int, rtNode node) map[string]int {
	combined := map[string]int{}
	for k, v := range l {
		if k == n.name || k == rtNode.name {
			continue
		}
		combined[k] = combined[k] + v
	}
	for k, v := range rtNode.connections {
		if k == n.name || k == rtNode.name {
			continue
		}
		combined[k] = combined[k] + v
	}
	return combined
}
