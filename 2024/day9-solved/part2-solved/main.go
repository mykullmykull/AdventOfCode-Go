package day

type file struct {
	id            int
	startingIndex int
	size          int
}

type space struct {
	startingIndex int
	size          int
}

func main(input string) int {
	files, spaces := parse(input)
	for f := len(files) - 1; f >= 0; f-- {
		for s := 0; s < len(spaces); s++ {
			if spaces[s].startingIndex >= files[f].startingIndex {
				break
			}
			if spaces[s].size >= files[f].size {
				files[f].startingIndex = spaces[s].startingIndex
				spaces[s].startingIndex += files[f].size
				spaces[s].size -= files[f].size
				if spaces[s].size == 0 {
					spaces = append(spaces[:s], spaces[s+1:]...)
				}
				break
			}
		}
	}

	checksum := 0
	for _, f := range files {
		for x := 0; x < f.size; x++ {
			checksum += f.id * (f.startingIndex + x)
		}
	}
	return checksum
}
