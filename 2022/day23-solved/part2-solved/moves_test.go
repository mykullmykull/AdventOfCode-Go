package day

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_moves(t *testing.T) {
	fmt.Println("Testing moves...") // keep this line to maintain imports
	assert.True(t, true)            // keep this line to maintain imports
}

// func Test_maybeProposeMove(t *testing.T) {
// 	assert.False(t, true)
// }

// func Test_proposeMoves(t *testing.T) {
// 	assert.False(t, true)
// }

// func Test_proposeMove(t *testing.T) {
// 	assert.False(t, true)
// }

func Test_hasBlocker(t *testing.T) {
	blockers := []string{
		"#",
		"^",
		"v",
		"<",
		">",
	}
	nonBlockers := []string{
		".",
		"O",
		"X",
	}
	for _, a := range nonBlockers {
		for _, b := range nonBlockers {
			for _, c := range nonBlockers {
				assert.False(t, hasBlocker(a, b, c))
			}
		}
	}
	values := []string{".", ".", "."}
	for x := 0; x < 3; x++ {
		for _, b := range blockers {
			values[x] = b
			assert.True(t, hasBlocker(values[0], values[1], values[2]))
			values[x] = nonBlockers[0]
			assert.False(t, hasBlocker(values[0], values[1], values[2]))
		}
	}
}

func Test_applyMoves(t *testing.T) {
	grid := []string{
		".v.#.#",
		".OO>X<",
		"..^.O<",
	}
	expected := []string{
		"...#.#",
		".###.#",
		"....#.",
	}
	assert.Equal(t, expected, applyMoves(grid))
}

func Test_makeMove(t *testing.T) {
	grid := []string{
		".v....",
		".OO>O.",
		"..^.O<",
	}
	expected := []string{
		"......",
		".##.#.",
		"....#.",
	}
	grid = makeMove(grid, 1, 1)
	grid = makeMove(grid, 1, 2)
	grid = makeMove(grid, 1, 4)
	grid = makeMove(grid, 2, 4)
	assert.Equal(t, expected, grid)
}

func Test_undoMoves(t *testing.T) {
	grid := []string{
		".v....",
		".X.>X<",
		".^....",
	}
	expected := []string{
		".#....",
		"...#.#",
		".#....",
	}
	grid = undoMoves(grid, 1, 1)
	grid = undoMoves(grid, 1, 4)
	assert.Equal(t, expected, grid)
}

func Test_setEmpty(t *testing.T) {
	grid := []string{
		".X..O.",
		".^>v<.",
		"......",
	}
	expected := []string{
		"......",
		"......",
		"......",
	}
	grid = setEmpty(grid, 0, 1, xRunes())
	grid = setEmpty(grid, 0, 4, oRunes())
	grid = setEmpty(grid, 1, 1, arrowRunes())
	grid = setEmpty(grid, 1, 2, arrowRunes())
	grid = setEmpty(grid, 1, 3, arrowRunes())
	grid = setEmpty(grid, 1, 4, arrowRunes())
	assert.Equal(t, expected, grid)
}

func Test_setElf(t *testing.T) {
	grid := []string{
		".X..O.",
		".^>v<.",
		"......",
	}
	expected := []string{
		".#..#.",
		".####.",
		"......",
	}
	grid = setElf(grid, 0, 1, xRunes())
	grid = setElf(grid, 0, 4, oRunes())
	grid = setElf(grid, 1, 1, arrowRunes())
	grid = setElf(grid, 1, 2, arrowRunes())
	grid = setElf(grid, 1, 3, arrowRunes())
	grid = setElf(grid, 1, 4, arrowRunes())
	assert.Equal(t, expected, grid)
}
