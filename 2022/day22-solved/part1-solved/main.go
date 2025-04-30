package day

func main(grid []string, instructions string) int {
	b := parse(grid, instructions)
	for b.instructions != "" {
		inst := ""
		inst, b = b.popInstruction()
		b = b.turn(inst)
		b = b.move(inst)
	}
	score := 0
	score += 1000 * (b.row + 1)
	score += 4 * (b.col + 1)
	score += b.getFacingScore()
	return score
}
