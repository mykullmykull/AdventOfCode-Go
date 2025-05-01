package day

import (
	"fmt"
	"strings"
)

func parse(input []string) ([]string, []board) {
	draws := []string{}
	boards := []board{}

	b := board{}
	for _, line := range input {
		if len(draws) == 0 {
			draws = strings.Split(line, ",")
			continue
		}
		if line == "" {
			if len(b.rows) > 0 {
				boards = append(boards, b)
				b = board{}
			}
			continue
		}
		cells := strings.Split(line, " ")
		row := []string{}
		for _, cell := range cells {
			trimmed := strings.Trim(cell, " ")
			if trimmed != "" {
				row = append(row, trimmed)
			}
		}
		b.rows = append(b.rows, row)
	}
	return draws, boards
}

func printBoards(boards []board) {
	for _, b := range boards {
		// fmt.Printf("Board: %v\n", b.rows)
		for _, row := range b.rows {
			fmt.Printf("Row: %v\n", row)
		}
		println()
	}
}
