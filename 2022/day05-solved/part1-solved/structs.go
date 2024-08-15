package day

import (
	"goAoc2022/utils"
	"strings"
)

type cargo struct {
	stacks []stack
}

type stack struct {
	topToBottom []string
}

func parseStacks(input []string) cargo {
	c := cargo{}
	for _, line := range input {
		split := strings.Split(line, "] [")
		if len(c.stacks) == 0 {
			for i := 0; i < len(split); i++ {
				c.stacks = append(c.stacks, stack{})
			}
		}
		for i, str := range split {
			name := strings.Replace(strings.Replace(str, "]", "", -1), "[", "", -1)
			if name != " " {
				c.stacks[i].topToBottom = append(c.stacks[i].topToBottom, name)
			}
		}
	}
	return c
}

func (c cargo) getTops() string {
	tops := ""
	for _, stack := range c.stacks {
		tops += string(stack.topToBottom[0])
	}
	return tops
}

func (c cargo) executeMoves(input []string) cargo {
	for _, command := range input {
		split := strings.Split(command, " ")
		numberOfCrates := utils.StrToInt(split[1])
		fromId := utils.StrToInt(split[3]) - 1
		toId := utils.StrToInt(split[5]) - 1
		for i := 0; i < numberOfCrates; i++ {
			crate := c.stacks[fromId].topToBottom[0]
			c.stacks[fromId].topToBottom = c.stacks[fromId].topToBottom[1:]
			c.stacks[toId].topToBottom = append([]string{crate}, c.stacks[toId].topToBottom...)
		}
	}
	return c
}
