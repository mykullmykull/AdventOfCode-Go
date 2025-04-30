package day

type file struct {
	id   int
	size int
}

func main(input string) int {
	files, spaces := parse(input)
	isMovingFile := false
	diskPosition := 0
	checksum := 0
	for len(files) > 0 {
		if diskPosition == 0 {
			isMovingFile = true
			diskPosition = files[0].size
			files = files[1:]
			continue
		}

		f := 0
		if isMovingFile {
			s := spaces[0]
			if s == 0 {
				spaces = spaces[1:]
				isMovingFile = false
				continue
			}

			f = len(files) - 1
			spaces[0]--
		}

		toAdd := diskPosition * files[f].id
		checksum += toAdd
		diskPosition++
		files[f].size--
		if files[f].size == 0 {
			files = removeFile(files, f)
			isMovingFile = true
		}
	}
	return checksum
}

func removeFile(files []file, f int) []file {
	if f == 0 {
		return files[1:]
	}
	return files[:f]
}
