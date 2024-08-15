package day

import "goAoc2022/utils"

type pointer struct {
	visibleTrees map[int]bool
	maxHeight    int
	r            int
	c            int
}

func (p pointer) resetMax() pointer {
	p.maxHeight = -1
	return p
}

func (p pointer) getTreeId() int {
	return p.r*1000 + p.c
}

func (p pointer) addToVisibleTrees(input []string) pointer {
	treeHeight := utils.StrToInt(string(input[p.r][p.c]))
	if treeHeight <= p.maxHeight {
		return p
	}
	p.maxHeight = treeHeight
	p.visibleTrees[p.getTreeId()] = true
	return p
}
