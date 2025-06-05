package day

import "goAoc2021/utils"

type tangle struct {
	octopi   []string
	flashing int
	count    int
}

type point struct {
	row int
	col int
}

func main(input []string) int {
	t := tangle{
		octopi:   input,
		flashing: 0,
		count:    0,
	}

	for step := 0; step < 100; step++ {
		for r, row := range input {
			for c := range row {
				p := point{row: r, col: c}
				t.buildEnergy(p)
			}
		}

		for t.flashing > 0 {
			for r, row := range input {
				for c, cell := range row {
					p := point{row: r, col: c}
					if cell == '*' {
						t.octopi[r] = utils.ReplaceAtIndex(t.octopi[r], c, "x")
						t.flashing--
						t.count++
						t.buildNeighborsEnergy(p)
					}
				}
			}
		}

		t.resetFlashed()
	}

	return t.count
}

func (t *tangle) buildEnergy(p point) {
	val := t.octopi[p.row][p.col]
	if val == '*' || val == 'x' {
		return
	}
	if val == '9' {
		t.octopi[p.row] = utils.ReplaceAtIndex(t.octopi[p.row], p.col, "*")
		t.flashing++
	} else {
		valInt := utils.Btoi(val)
		valInt++
		t.octopi[p.row] = utils.ReplaceAtIndex(t.octopi[p.row], p.col, utils.Itoa(valInt))
	}
}

func (t *tangle) buildNeighborsEnergy(center point) {
	neighbors := []point{
		{row: center.row - 1, col: center.col - 1},
		{row: center.row - 1, col: center.col},
		{row: center.row - 1, col: center.col + 1},
		{row: center.row, col: center.col - 1},
		{row: center.row, col: center.col + 1},
		{row: center.row + 1, col: center.col - 1},
		{row: center.row + 1, col: center.col},
		{row: center.row + 1, col: center.col + 1},
	}

	for _, n := range neighbors {
		if n.row >= 0 && n.row < len(t.octopi) && n.col >= 0 && n.col < len(t.octopi[n.row]) {
			t.buildEnergy(n)
		}
	}
}

func (t *tangle) resetFlashed() {
	for r, row := range t.octopi {
		for c, cell := range row {
			if cell == 'x' {
				t.octopi[r] = utils.ReplaceAtIndex(t.octopi[r], c, "0")
			}
		}
	}
}
