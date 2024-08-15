package day

import (
	"goAoc2022/utils"
	"strings"
)

func part2(input []string) int {
	directories := parse(input)
	spaceUsed := directories["/"]
	spaceNeeded := 30000000 - (70000000 - spaceUsed)
	println("\nspaceUsed", spaceUsed, "spaceNeeded", spaceNeeded)
	smallestBigEnoughDirectory := 70000000
	for _, v := range directories {
		println("before: smallest", smallestBigEnoughDirectory, "v", v)
		if v >= spaceNeeded && v < smallestBigEnoughDirectory {
			println("  after: smallest", smallestBigEnoughDirectory)
			smallestBigEnoughDirectory = v
		}
	}
	return smallestBigEnoughDirectory
}

func parse(input []string) map[string]int {
	directories := map[string]int{}
	p := path{}
	for _, line := range input {
		if len(line) >= 5 && line[:5] == "$ cd " {
			p = p.cd(line[5:])
			continue
		}
		if len(line) >= 4 && line[:4] == "$ ls" || line[:4] == "dir " {
			continue
		}
		split := strings.Split(line, " ")
		size := utils.StrToInt(split[0])
		dirs := p.dirs
		for _, d := range dirs {
			directories[d] += size
		}
	}
	return directories
}
