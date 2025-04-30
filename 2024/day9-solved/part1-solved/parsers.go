package day

import "goAoc2024/utils"

func parse(input string) ([]file, []int) {
	isFile := true
	files := []file{}
	spaces := []int{}
	id := 0
	for x := 0; x < len(input); x++ {
		size := utils.Atoi(input[x : x+1])
		if isFile {
			f := file{id, size}
			files = append(files, f)
			id++
		}
		if !isFile {
			spaces = append(spaces, size)
		}
		isFile = !isFile
	}
	return files, spaces
}
