package day

import (
	"goAoc2022/utils"
	"strings"
)

func main(input []string) int {
	cubes := parse(input)
	freeSides := 0
	for _, c := range cubes {
		freeSides += c.countFreeSides(cubes)
	}
	return freeSides
}

func parse(input []string) []cube {
	cubes := []cube{}
	for _, line := range input {
		split := strings.Split(line, ",")
		cubes = append(cubes, cube{
			x: utils.StrToInt(split[0]),
			y: utils.StrToInt(split[1]),
			z: utils.StrToInt(split[2]),
		})
	}
	return cubes
}
