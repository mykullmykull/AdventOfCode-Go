package day

import (
	"goAoc2021/utils"
)

type board struct {
	rows [][]string
}

func main(input []string) int {
	draws, boards := parse(input)
	score := 0
	for _, draw := range draws {
		updated := []board{}
		for len(boards) > 0 {
			b := boards[0]
			boards = boards[1:]
			b.mark(draw)
			won := b.isWinner()
			if won && len(boards)+len(updated) == 0 {
				score = b.calculateScore(draw)
				return score
			}
			if won {
				continue
			}
			updated = append(updated, b)
		}
		boards = updated
	}

	panic("Not all boards are winners")
}

func (b *board) mark(num string) {
	for r, row := range b.rows {
		for c, col := range row {
			if col == num {
				b.rows[r][c] = "X"
			}
		}
	}
}

func (b *board) isWinner() bool {
	for _, row := range b.rows {
		rowXs := 0
		for _, col := range row {
			if col == "X" {
				rowXs++
			}
		}
		if rowXs == 5 {
			return true
		}
	}

	for c, _ := range b.rows[0] {
		colXs := 0
		for _, row := range b.rows {
			if row[c] == "X" {
				colXs++
			}
		}
		if colXs == 5 {
			return true
		}
	}
	return false
}

func (b *board) calculateScore(lastDrawn string) int {
	score := 0
	for _, row := range b.rows {
		for _, col := range row {
			if col == "X" {
				continue
			}
			score += utils.Atoi(col)
		}
	}
	return score * utils.Atoi(lastDrawn)
}
