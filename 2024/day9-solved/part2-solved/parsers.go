package day

import "goAoc2024/utils"

func parse(input string) ([]file, []space) {
	isFile := true
	files := []file{}
	spaces := []space{}
	id := 0
	index := 0
	for x := 0; x < len(input); x++ {
		size := utils.Atoi(input[x : x+1])
		if isFile {
			f := file{id, index, size}
			files = append(files, f)
			id++
		}
		if !isFile && size > 0 {
			s := space{index, size}
			spaces = append(spaces, s)
		}
		index += size
		isFile = !isFile
	}
	return files, spaces
}
