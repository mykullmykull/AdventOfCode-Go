package day

func main(input []string) int {
	start := point{r: 0, c: 1}
	v := valley{
		grid: input,
		elfs: map[point]bool{start: true},
	}
	for {
		if v.isFinished() {
			return v.time
		}
		v.grid = v.moveBlizzards()
		v.elfs = v.getNewElfs()
		v.time++
	}
}
