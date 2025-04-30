package day

func main(grid []string, instructions string) int {
	b := parse(grid, instructions)
	println(b.toString())
	for b.instructions != "" {
		inst := ""
		inst, b.instructions = popInstruction(b.instructions)
		b.p = b.p.turn(inst)
		b = b.move(inst)

		println("after", inst, "side:", side(b.getSide(b.p)), "facing:", facing(b.p.getFacing()))
		b.printSide()
		println()
	}
	score := 0
	score += 1000 * (b.p.row + 1)
	score += 4 * (b.p.col + 1)
	score += b.getFacingScore()
	return score
}

func popInstruction(instructions string) (string, string) {
	firstChar := instructions[0]
	if firstChar == 'R' {
		return "R", instructions[1:]
	}
	if firstChar == 'L' {
		return "L", instructions[1:]
	}
	str := ""
	for x := 0; x < len(instructions); x++ {
		if instructions[x] == 'R' || instructions[x] == 'L' {
			break
		}
		str += instructions[x : x+1]
	}
	return str, instructions[len(str):]
}
