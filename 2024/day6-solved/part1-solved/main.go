package day

func main(grid []string) int {
	l := lab{grid: grid}
	l = l.findGuard()
	for {
		oldGuard := guard{
			row: l.guard.row,
			col: l.guard.col,
			dir: l.guard.dir,
		}
		l = l.moveGuard()
		l = l.updateGrid(oldGuard)
		if l.isFinished() {
			break
		}
	}
	return l.countVisitedCells()
}
