package day12

type node struct {
	guessed string
}

type remainder struct {
	springs         string
	groupsLeft      int
	inProgressCount int
}

func (n *node) yesChild() node {
	return node{
		guessed: n.guessed + "Y",
	}
}

func (n *node) noChild() node {
	return node{
		guessed: n.guessed + "N",
	}
}

func (n *node) parent() node {
	return node{
		guessed: n.guessed[:len(n.guessed)-1],
	}
}

func (n *node) fails(base string, groups []int, unknowns int) bool {
	guessedIndex := 0
	groupIndex := 0
	group := groups[groupIndex]
	counted := 0
	for i, r := range base {
		c := string(r)
		if c == "?" {
			if guessedIndex >= len(n.guessed) {
				return false
			}

			c = string(n.guessed[guessedIndex])
			guessedIndex++
		}

		if c == "Y" || c == "#" {
			counted++
			if i == len(base)-1 && groupIndex == len(groups)-1 {
				return counted != group
			}
		}

		if counted > group {
			return true
		}

		if c == "." || c == "N" {
			if counted == 0 {
				continue
			}

			if counted < group {
				return true
			}

			// we now know that counted == group
			counted = 0
			groupIndex++
			if groupIndex >= len(groups) {
				group = 0
				continue
			}
			group = groups[groupIndex]
		}
	}
	return groupIndex != len(groups)
}

func (n *node) getRemainder(springs string, groups []int) remainder {
	remainingSprings := ""
	guessedIndex := 0
	groupIndex := 0
	group := groups[groupIndex]
	counted := 0
	for i, c := range springs {
		if guessedIndex >= len(n.guessed) {
			remainingSprings = springs[i:]
			break
		}

		if c == '?' {
			c = rune(n.guessed[guessedIndex])
			guessedIndex++
		}

		if c == 'Y' || c == '#' {
			counted++
		}

		if counted > group {
			return remainder{}
		}

		if c == '.' || c == 'N' {
			if counted == 0 {
				continue
			}

			if counted < group {
				return remainder{}
			}

			// we now know that counted == group
			counted = 0
			groupIndex++
			if groupIndex >= len(groups) {
				group = 0
				continue
			}
			group = groups[groupIndex]
		}
	}
	return remainder{
		springs:         remainingSprings,
		groupsLeft:      len(groups) - groupIndex,
		inProgressCount: counted,
	}
}
