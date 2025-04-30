package day

import "fmt"

func main(input []string, dRobotCount int) int {
	complexitySum := 0
	nPad = nPad.mapPaths()
	dPad = dPad.mapPaths()
	for k, v := range dPad.paths {
		fmt.Printf("pair %v, paths: %v\n", k, v)
	}
	return complexitySum
}
