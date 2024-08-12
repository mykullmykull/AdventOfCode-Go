package day

import (
	"goAoc2022/utils"
	"strings"
)

func part1(input []string) int {
	directories := parse(input)
	sumOfSmallDirectories := 0
	for k, v := range directories {
		println(k, ":", v)
		if v <= 100000 {
			// fmt.Printf("directory %s has small value %d\n", k, v)
			sumOfSmallDirectories += v
			// fmt.Printf("new sum: %d\n", sumOfSmallDirectories)
		}
	}
	return sumOfSmallDirectories
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
