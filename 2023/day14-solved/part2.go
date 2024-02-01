package day14

import (
	"fmt"
)

type sortedBlocks struct {
	sortedForN []point
	sortedForW []point
	sortedForS []point
	sortedForE []point
}

func runB(image []string) int {
	rocks, blocks := findRocks(image)

	maxRow := len(image) - 1
	maxCol := len(image[0]) - 1

	historicRocks := [][]point{}
	repeatedRocksIndex := -1
	firstRepeatedOnCycle := -1
	secondRepeatedOnCycle := -1

	for i := 0; i < 1000000000; i++ {
		rocks = cycleRocks(rocks, blocks, maxRow, maxCol, image)
		previousCycle := isRepeatOfCycle(rocks, historicRocks)
		if previousCycle > -1 {
			if firstRepeatedOnCycle == -1 {
				repeatedRocksIndex = previousCycle
				firstRepeatedOnCycle = i
				continue
			}

			if previousCycle == repeatedRocksIndex {
				secondRepeatedOnCycle = i
				break
			}
		}

		copyRocks := make([]point, len(rocks))
		copy(copyRocks, rocks)
		historicRocks = append(historicRocks, copyRocks)
	}

	repeatDurations := secondRepeatedOnCycle - firstRepeatedOnCycle
	iStart := ((1000000000/repeatDurations)-2)*repeatDurations + firstRepeatedOnCycle + 1
	for i := iStart; i < 1000000000; i++ {
		rocks = cycleRocks(rocks, blocks, maxRow, maxCol, image)
	}
	return totalPressureNorth(rocks, maxRow)
}

func findRocks(image []string) ([]point, sortedBlocks) {
	rocks := []point{}

	blocksSortedForN := initiateBlocks(len(image[0])-1, "row", -1)
	blocksSortedForW := initiateBlocks(len(image)-1, "col", -1)
	blocksSortedForS := initiateBlocks(len(image[0])-1, "row", len(image))
	blocksSortedForE := initiateBlocks(len(image)-1, "col", len(image[0]))
	for r, row := range image {
		for c, char := range row {
			if char == '.' {
				continue
			}

			p := point{r, c}
			if char == 'O' {
				rocks = append(rocks, p)
				continue
			}

			blocksSortedForN = insertPointInOrder(blocksSortedForN, p, "row", 1)
			blocksSortedForW = insertPointInOrder(blocksSortedForW, p, "col", 1)
			blocksSortedForS = insertPointInOrder(blocksSortedForS, p, "row", -1)
			blocksSortedForE = insertPointInOrder(blocksSortedForE, p, "col", -1)

		}
	}
	return rocks, sortedBlocks{blocksSortedForN, blocksSortedForW, blocksSortedForS, blocksSortedForE}
}

func initiateBlocks(maxOther int, setDimmension string, setValue int) []point {
	points := []point{}
	if setDimmension == "row" {
		if maxOther != 0 {
			for i := 0; i <= maxOther; i++ {
				points = append(points, point{setValue, i})
			}
			return points
		}

		for i := maxOther; i >= 0; i -= 1 {
			points = append(points, point{setValue, i})
			return points
		}
	}

	if setDimmension == "col" {
		if maxOther != 0 {
			for i := 0; i <= maxOther; i++ {
				points = append(points, point{i, setValue})
			}
			return points
		}

		for i := maxOther; i >= 0; i -= 1 {
			points = append(points, point{i, setValue})
			return points
		}
	}
	return points
}

func insertPointInOrder(points []point, newP point, sortBy string, direction int) []point {
	if len(points) == 0 {
		return []point{newP}
	}

	insertBefore := -1
	for i, p := range points {
		if sortBy == "row" && direction*p.row > direction*newP.row {
			insertBefore = i
			break
		}

		if sortBy == "col" && direction*p.col > direction*newP.col {
			insertBefore = i
			break
		}
	}

	if insertBefore == -1 {
		return append(points, newP)
	}

	newPoints := append(points[:insertBefore+1], points[insertBefore:]...)
	newPoints[insertBefore] = newP
	return newPoints
}

func cycleRocks(rocks []point, blocks sortedBlocks, maxRow int, maxCol int, image []string) []point {
	copyBlocksN := make([]point, len(blocks.sortedForN))
	copyBlocksW := make([]point, len(blocks.sortedForW))
	copyBlocksS := make([]point, len(blocks.sortedForS))
	copyBlocksE := make([]point, len(blocks.sortedForE))
	copy(copyBlocksN, blocks.sortedForN)
	copy(copyBlocksW, blocks.sortedForW)
	copy(copyBlocksS, blocks.sortedForS)
	copy(copyBlocksE, blocks.sortedForE)
	copyBlocks := sortedBlocks{copyBlocksN, copyBlocksW, copyBlocksS, copyBlocksE}

	rocks = moveRocks(rocks, copyBlocks, "n", maxRow, maxCol)
	rocks = moveRocks(rocks, copyBlocks, "w", maxRow, maxCol)
	rocks = moveRocks(rocks, copyBlocks, "s", maxRow, maxCol)
	rocks = moveRocks(rocks, copyBlocks, "e", maxRow, maxCol)

	return rocks
}

func moveRocks(rocks []point, allBlocks sortedBlocks, direction string, maxRow int, maxCol int) []point {
	switch direction {
	case "n":
		{
			blocks := allBlocks.sortedForN
			i := 0
			for {
				if i >= len(blocks) {
					break
				}

				block := blocks[i]
				nextBlock := point{maxRow + 1, block.col} // default to a point past the end
				for j, possibleNextBlock := range blocks {
					if j > i && possibleNextBlock.col == block.col {
						nextBlock = possibleNextBlock
						break
					}
				}

				for j, rock := range rocks {
					if rock.col != block.col {
						continue
					}

					if rock.row > block.row && rock.row < nextBlock.row {
						blocks[i].row++
						rocks[j].row = blocks[i].row
					}
				}
				i++
			}
		}
	case "w":
		{
			blocks := allBlocks.sortedForW
			i := 0
			for {
				if i >= len(blocks) {
					break
				}

				block := blocks[i]
				nextBlock := point{block.row, maxCol + 1} // default to a point past the end
				for j, possibleNextBlock := range blocks {
					if j > i && possibleNextBlock.row == block.row {
						nextBlock = possibleNextBlock
						break
					}
				}

				for j, rock := range rocks {
					if rock.row != block.row {
						continue
					}

					if rock.col > block.col && rock.col < nextBlock.col {
						blocks[i].col++
						rocks[j].col = blocks[i].col
					}
				}
				i++
			}
		}
	case "s":
		{
			blocks := allBlocks.sortedForS
			i := 0
			for {
				if i >= len(blocks) {
					break
				}

				block := blocks[i]
				nextBlock := point{-1, block.col} // default to a point past the end
				for j, possibleNextBlock := range blocks {
					if j > i && possibleNextBlock.col == block.col {
						nextBlock = possibleNextBlock
						break
					}
				}

				for j, rock := range rocks {
					if rock.col != block.col {
						continue
					}

					if rock.row < block.row && rock.row > nextBlock.row {
						blocks[i].row -= 1
						rocks[j].row = blocks[i].row
					}
				}
				i++
			}
		}
	case "e":
		{
			blocks := allBlocks.sortedForE
			i := 0
			for {
				if i >= len(blocks) {
					break
				}

				block := blocks[i]
				nextBlock := point{block.row, -1} // default to a point past the end
				for j, possibleNextBlock := range blocks {
					if j > i && possibleNextBlock.row == block.row {
						nextBlock = possibleNextBlock
						break
					}
				}

				for j, rock := range rocks {
					if rock.row != block.row {
						continue
					}

					if rock.col < block.col && rock.col > nextBlock.col {
						blocks[i].col -= 1
						rocks[j].col = blocks[i].col
					}
				}
				i++
			}
		}
	}
	return rocks
}

func totalPressureNorth(rocks []point, maxRow int) int {
	sum := 0
	for _, rock := range rocks {
		sum += maxRow + 1 - rock.row
	}
	return sum
}

func printRocks(rocks []point, image []string) {
	fmt.Println()
	maxRow := len(image)
	maxCol := len(image[0])
	for r := 0; r < maxRow; r++ {
		line := ""
		for c := 0; c < maxCol; c++ {
			foundRock := false
			for _, rock := range rocks {
				if rock.row == r && rock.col == c {
					line += "O"
					foundRock = true
					break
				}
			}
			if !foundRock {
				originalChar := string(image[r][c])
				if originalChar == "O" {
					originalChar = "."
				}
				line += originalChar
			}
		}
		fmt.Println(line)
	}
}

func isRepeatOfCycle(rocks []point, historicRocks [][]point) int {
	for i, pastRocks := range historicRocks {
		if isEqual(rocks, pastRocks) {
			return i
		}
	}
	return -1
}

func isEqual(rocks1 []point, rocks2 []point) bool {
	for _, r1 := range rocks1 {
		foundMatch := false
		for _, r2 := range rocks2 {
			if r1.col == r2.col && r1.row == r2.row {
				foundMatch = true
				break
			}
		}
		if !foundMatch {
			return false
		}
	}
	return true
}
