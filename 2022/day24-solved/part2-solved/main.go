package day

func main(input []string) int {
	start := point{r: 0, c: 1}
	v := valley{
		grid: input,
		elfs: map[point]bool{start: true},
	}
	for {
		if v.haveReachedTheFinish() {
			break
		}
		v.grid = v.moveBlizzards()
		v.elfs = v.getNewElfs()
		v.time++
	}
	toFinish := v.time
	println("time to finish", v.time)
	finish := point{r: len(v.grid) - 1, c: len(v.grid[0]) - 2}
	v.elfs = map[point]bool{finish: true}
	for {
		if v.haveReachedTheStart() {
			break
		}
		v.grid = v.moveBlizzards()
		v.elfs = v.getNewElfs()
		v.time++
	}
	toStart := v.time
	println("time to start", v.time-toFinish)
	v.elfs = map[point]bool{start: true}
	for {
		if v.haveReachedTheFinish() {
			break
		}
		v.grid = v.moveBlizzards()
		v.elfs = v.getNewElfs()
		v.time++
	}
	println("time to finish", v.time-toStart)
	println("total time", v.time)
	return v.time
}
