package day

import (
	"fmt"
	"strings"
)

type caves struct {
	countPaths  int
	connections map[string][]string
	visited     map[string]int
	usedRepeat  bool
	path        []string
	indexes     []int
	isFinished  bool
}

func (c *caves) init(input []string) {
	c.visited = make(map[string]int)
	c.usedRepeat = false
	c.path = []string{"start"}
	c.indexes = []int{}
	c.parseConnections(input)
}

func (c *caves) reachedEnd() bool {
	return c.path[len(c.path)-1] == "end"
}

func (c *caves) rewind() int {
	lastCave := c.path[len(c.path)-1]
	c.path = c.path[:len(c.path)-1]
	if lastCave == "start" {
		c.isFinished = true
		return -1
	}

	if strings.ToLower(lastCave) == lastCave {
		visits := c.visited[lastCave]
		if visits == 2 {
			c.usedRepeat = false
		}
	}
	c.visited[lastCave]--
	if c.visited[lastCave] == 0 {
		delete(c.visited, lastCave)
	}

	lastIndex := c.indexes[len(c.indexes)-1]
	c.indexes = c.indexes[:len(c.indexes)-1]
	return lastIndex
}

func (c *caves) addNextCave(lastIndex int) {
	if c.isFinished {
		return
	}

	nextIndex := lastIndex + 1
	lastCave := c.path[len(c.path)-1]
	nextCaves := c.connections[lastCave]
	if nextIndex >= len(nextCaves) {
		lastIndex = c.rewind()
		c.addNextCave(lastIndex)
		return
	}

	nextCave := nextCaves[nextIndex]
	c.appendIfValid(nextCave, nextIndex)
}

func (c *caves) appendIfValid(cave string, index int) {
	if cave == "start" || c.isSecondRepeat(cave) {
		c.addNextCave(index)
		return
	}
	c.path = append(c.path, cave)
	c.indexes = append(c.indexes, index)
	c.visited[cave]++
}

func (c *caves) isSecondRepeat(cave string) bool {
	alreadyUsed := c.usedRepeat
	if !c.isRepeat(cave) {
		return false
	}
	return alreadyUsed
}

func (c *caves) isRepeat(cave string) bool {
	if strings.ToLower(cave) != cave {
		return false
	}
	if c.visited[cave] > 0 {
		c.usedRepeat = true
		return true
	}
	return false
}

func (c *caves) Print() {
	fmt.Printf("Path: %v\n", c.path)
	fmt.Printf("  Visited: %v\n", c.visited)
	fmt.Printf("  Used Repeat: %v\n", c.usedRepeat)
	fmt.Printf("  Count Paths: %v\n", c.countPaths)
	fmt.Printf("  Indexes: %v\n", c.indexes)
	println()
}
